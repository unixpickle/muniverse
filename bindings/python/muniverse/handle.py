"""
Manage and communicate with the muniverse-bind process.
"""

import os
import subprocess
import threading

from . import proto
from .errors import ProtoError, MuniverseError, CallError

# Enable global variables, which are used to share one
# global process between otherwise separate objects.

# pylint: disable=W0603

_CURRENT_SESSION_LOCK = threading.Lock()
_CURRENT_SESSION = None
_CURRENT_SESSION_REFS = 0

class Handle:
    """
    A thread-safe handle on a muniverse-bind instance.
    """
    def __init__(self):
        """
        Create a new handle.

        This may throw a FileNotFoundError or a related
        error if the muniverse-bind command cannot be
        executed.
        """
        global _CURRENT_SESSION
        global _CURRENT_SESSION_LOCK
        global _CURRENT_SESSION_REFS
        _CURRENT_SESSION_LOCK.acquire()
        try:
            if _CURRENT_SESSION is None:
                _CURRENT_SESSION = _Session()
            self.session = _CURRENT_SESSION
            _CURRENT_SESSION_REFS += 1
        finally:
            _CURRENT_SESSION_LOCK.release()
        self.call_lock = threading.Lock()
        self.closed = False

    def close(self):
        """
        Closes the handle.

        This waits for pending calls, and any calls made
        afterwards will fail.

        Calling close more than once has no effect.
        """
        global _CURRENT_SESSION
        global _CURRENT_SESSION_LOCK
        global _CURRENT_SESSION_REFS
        self.call_lock.acquire()
        _CURRENT_SESSION_LOCK.acquire()
        try:
            if self.closed:
                return
            self.closed = True
            _CURRENT_SESSION_REFS -= 1
            if _CURRENT_SESSION_REFS == 0:
                _CURRENT_SESSION.close()
                _CURRENT_SESSION = None
        finally:
            self.call_lock.release()
            _CURRENT_SESSION_LOCK.release()

    def call(self, name, args):
        """
        Make an API call and wait for the result.

        This may throw a MuniverseError or an error
        relating to a closed pipe in the case of a
        communication failure.

        This does not throw an exception for errors
        in the return value itself.
        """
        self.call_lock.acquire()
        try:
            if self.closed:
                raise MuniverseError('call on closed handle')
            return self.session.call(name, args)
        finally:
            self.call_lock.release()

    def checked_call(self, name, args):
        """
        Like call(), but raises a CallError if the call
        results in an error on the back-end.
        """
        res = self.call(name, args)
        if 'Error' in res:
            raise CallError(res['Error'])
        return res

class _Session:
    """
    Manages an instance of muniverse-bind.
    """
    def __init__(self):
        try:
            self.proc = subprocess.Popen(['muniverse-bind'],
                                         stdin=subprocess.PIPE,
                                         stdout=subprocess.PIPE)
        except OSError:
            gopath = os.environ['GOPATH']
            exc_path = os.path.join(gopath, 'bin', 'muniverse-bind')
            self.proc = subprocess.Popen([exc_path],
                                         stdin=subprocess.PIPE,
                                         stdout=subprocess.PIPE)
        self.read_thread = threading.Thread(target=self._read_loop)
        self.read_thread.start()
        self.waiting_lock = threading.Lock()
        self.waiting = {}
        self.current_id_lock = threading.Lock()
        self.current_id = 0
        self.send_lock = threading.Lock()

    def close(self):
        """
        Terminate the background process and wait for
        threads to exit.

        This should only be called one time, and it should
        not be called while any other methods are pending.
        """
        self.proc.kill()
        self.read_thread.join()

    def call(self, name, args):
        """
        Make a named API call and wait for the response.
        """
        self.current_id_lock.acquire()
        call_id = str(self.current_id)
        self.current_id += 1
        self.current_id_lock.release()
        call_obj = {'ID': call_id, name: args}
        waiting_obj = {'event': threading.Event()}
        self.waiting_lock.acquire()
        try:
            if self.waiting is None:
                raise MuniverseError('session has died')
            self.waiting[call_id] = waiting_obj
        finally:
            self.waiting_lock.release()
        self.send_lock.acquire()
        try:
            proto.write_object(self.proc.stdin, call_obj)
        finally:
            self.send_lock.release()
        waiting_obj['event'].wait()
        if 'error' in waiting_obj:
            raise waiting_obj['error']
        else:
            return waiting_obj['payload']

    def _read_loop(self):
        try:
            while True:
                payload = proto.read_object(self.proc.stdout)
                if not 'ID' in payload:
                    raise ProtoError('missing ID in response')
                call_id = payload['ID']
                self.waiting_lock.acquire()
                try:
                    if not call_id in self.waiting:
                        raise ProtoError('unregistered call ID: ' + call_id)
                    waiting = self.waiting[call_id]
                    waiting['payload'] = payload
                    waiting['event'].set()
                    del self.waiting[call_id]
                finally:
                    self.waiting_lock.release()
        except MuniverseError as exc:
            self._kill_waiting(exc)

    def _kill_waiting(self, exc):
        self.waiting_lock.acquire()
        try:
            for uid in self.waiting:
                waiting = self.waiting[uid]
                waiting['error'] = exc
                waiting['event'].set()
        finally:
            self.waiting = None
            self.waiting_lock.release()

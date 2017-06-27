"""
API for controlling muniverse environments.
"""

import numpy as np

from .handle import Handle

class Env:
    """
    An environment instance.
    """
    def __init__(self, spec, container=None, chrome_host=None, game_host=None):
        """
        Create a new environment from the specification.

        If container is set, launch the environment in a
        custom container.

        If chrome_host and game_host are set, launch the
        environment in a running chrome instance.

        This may raise an exception if a muniverse-bind
        process cannot be started.
        It also raises a MuniverseError if the environment
        cannot be created on the backend.
        """
        if (chrome_host is None) != (game_host is None):
            raise ValueError('must set both chrome_host and game_host')
        elif (not container is None) and (not chrome_host is None):
            raise ValueError('cannot mix chrome_host and container options')
        call_obj = {'Spec': spec}
        call_name = 'NewEnv'
        if not container is None:
            call_obj['Container'] = container
            call_name = 'NewEnvContainer'
        elif not chrome_host is None:
            call_obj['Host'] = chrome_host
            call_obj['GameHost'] = game_host
            call_name = 'NewEnvChrome'
        self.handle = None
        handle = Handle()
        try:
            self.uid = handle.checked_call(call_name, call_obj)['UID']
            self.handle = handle
        finally:
            if self.handle is None:
                handle.close()

    def reset(self):
        """
        Reset the environment to a starting state.
        """
        self.handle.checked_call('Reset', {'UID': self.uid})

    def observe(self):
        """
        Capture a screenshot of the environment.

        The returned value is a 3D numpy array.
        The first index is y, the second is x, the third
        is depth (where depth goes in RGB order).
        """
        res = self.handle.checked_call('Observe', {'UID': self.uid})
        obs = res['Observation']
        dim = [obs['Height'], obs['Width'], 3]
        data = obs['RGB']
        return np.frombuffer(data, dtype=np.uint8).reshape(dim)

    def step(self, seconds, *actions):
        """
        Send actions to the environment and advance time.

        The return value is a tuple (reward, done).
        If done is True, then no more steps can be taken
        until a reset().
        """
        events = []
        for action in actions:
            events.append(action.to_bsonable())
        args = {
            'UID': self.uid,
            'Seconds': seconds,
            'Events': events
        }
        res = self.handle.checked_call('Step', args)
        info = res['StepResult']
        return info['Reward'], info['Done'] != 0

    def close(self):
        """
        Stop and clean up the environment.

        You should not close an environment multiple times.
        """
        try:
            self.handle.checked_call('CloseEnv', {'UID': self.uid})
        finally:
            self.handle.close()

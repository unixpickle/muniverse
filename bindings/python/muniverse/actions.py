"""
Types for environment actions.
"""

from .handle import Handle

# Allow more than 5 arguments to a constructor.
# pylint: disable=R0913

def key_for_code(code):
    """
    Create a KeyAction with the given key code.

    The result is None if the code is not found.

    The result will have no 'event', so it will be
    necessary to run with_event() on it.
    """
    handle = Handle()
    try:
        res = handle.checked_call('KeyForCode', {'Code': code})
        if 'KeyEvent' in res:
            action = KeyAction('')
            action.bson_obj = res['KeyEvent']
            return action
        return None
    finally:
        handle.close()

class MouseAction:
    """
    The action type for mouse events.
    """
    def __init__(self, event, x=0, y=0, button="left", click_count=0):
        """
        Create a new mouse event.

        The event is "mousePressed", "mouseReleased",
        or "mouseMoved".
        """
        self.bson_obj = {
            'Type': event,
            'X': x,
            'Y': y,
            'Button': button,
            'ClickCount': click_count
        }

    def to_bsonable(self):
        """
        Create an event object for the back-end.
        """
        return {'MouseEvent': self.bson_obj}

    def with_event(self, event):
        """
        Create a copy of this action with a different
        event name.
        """
        res = MouseAction('', 0, 0)
        res.bson_obj = dict(self.bson_obj, Type=event)
        return res

class KeyAction:
    """
    The action type for keyboard events.
    """
    def __init__(self, event, modifiers=0, timestamp=0, text="", key="",
                 unmodified_text="", key_identifier="", code="",
                 auto_repeat=False, is_keypad=False, is_system_key=False,
                 windows_virtual_key_code=0, native_virtual_key_code=0):
        """
        Create a new keyboard event.

        The event may be "keyDown", "keyUp", or
        "rawKeyDown".
        """
        self.bson_obj = {
            'Type': event,
            'Modifiers': modifiers,
            'Timestamp': timestamp,
            'Text': text,
            'Key': key,
            'UnmodifiedText': unmodified_text,
            'KeyIdentifier': key_identifier,
            'Code': code,
            'AutoRepeat': auto_repeat,
            'IsKeypad': is_keypad,
            'IsSystemKey': is_system_key,
            'WindowsVirtualKeyCode': windows_virtual_key_code,
            'NativeVirtualKeyCode': native_virtual_key_code
        }

    def to_bsonable(self):
        """
        Create an event object for the back-end.
        """
        return {'KeyEvent': self.bson_obj}

    def with_event(self, event):
        """
        Create a copy of this action with a different
        event name.
        """
        res = KeyAction('')
        res.bson_obj = dict(self.bson_obj, Type=event)
        return res

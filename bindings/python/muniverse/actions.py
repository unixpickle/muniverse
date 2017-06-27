"""
Types for environment actions.
"""

# Allow more than 5 arguments to a constructor and
# allow a ton of instance fields.
# pylint: disable=R0913
# pylint: disable=R0902

def lookup_key_action(code):
    """
    Create a KeyAction with the given key code.

    The result is None if the code is not found.

    The result will have no 'event', so it will be
    necessary to run with_event() on it.
    """
    # TODO: this.
    return None

class MouseAction:
    """
    The action type for mouse events.
    """
    def __init__(self, event, x, y, button="left", click_count=0):
        """
        Create a new mouse event.

        The event is "mousePressed", "mouseReleased",
        or "mouseMoved".
        """
        self.coord_x = x
        self.coord_y = y
        self.event = event
        self.button = button
        self.click_count = click_count

    def to_bsonable(self):
        """
        Create an event object for the back-end.
        """
        return {
            'Type': self.event,
            'X': self.coord_x,
            'Y': self.coord_y,
            'Button': self.button,
            'ClickCount': self.click_count
        }

    def with_event(self, event):
        """
        Create a copy of this action with a different
        event name.
        """
        return MouseAction(event, self.coord_x, self.coord_y, self.button,
                           self.click_count)

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
        self.event = event
        self.modifiers = modifiers
        self.timestamp = timestamp
        self.text = text
        self.key = key
        self.unmodified_text = unmodified_text
        self.key_identifier = key_identifier
        self.code = code
        self.auto_repeat = auto_repeat
        self.is_keypad = is_keypad
        self.is_system_key = is_system_key
        self.windows_virtual_key_code = windows_virtual_key_code
        self.native_virtual_key_code = native_virtual_key_code

    def to_bsonable(self):
        """
        Create an event object for the back-end.
        """
        return {
            'Event': self.event,
            'Modifiers': self.modifiers,
            'Timestamp': self.timestamp,
            'Text': self.text,
            'Key': self.key,
            'UnmodifiedText': self.unmodified_text,
            'KeyIdentifier': self.key_identifier,
            'Code': self.code,
            'AutoRepeat': self.auto_repeat,
            'IsKeypad': self.is_keypad,
            'IsSystemKey': self.is_system_key,
            'WindowsVirtualKeyCode': self.windows_virtual_key_code,
            'NativeVirtualKeyCode': self.native_virtual_key_code
        }

    def with_event(self, event):
        """
        Create a copy of this action with a different
        event name.
        """
        return KeyAction(event, self.modifiers, self.timestamp, self.text,
                         self.key, self.unmodified_text, self.key_identifier,
                         self.code, self.auto_repeat, self.is_keypad,
                         self.is_system_key, self.windows_virtual_key_code,
                         self.native_virtual_key_code)

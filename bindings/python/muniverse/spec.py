"""
API for dealing with environment specifications.
"""

from .handle import Handle

def spec_for_name(name):
    """
    Find an environment specification for the given name.

    If the specification is not found, None is returned.
    """
    handle = Handle()
    try:
        return handle.checked_call('SpecForName', {'Name': name})['Spec']
    finally:
        handle.close()

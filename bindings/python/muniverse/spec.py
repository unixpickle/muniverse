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
        res = handle.checked_call('SpecForName', {'Name': name})
        if 'Spec' in res:
            return res['Spec']
        return None
    finally:
        handle.close()

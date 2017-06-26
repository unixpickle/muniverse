"""
Error types for muniverse.
"""

class MuniverseError(Exception):
    """
    Exception base class used for muniverse APIs.
    """
    pass

class ProtoError(MuniverseError):
    """
    Low-level protocol error.
    """
    pass

class CallError(MuniverseError):
    """
    Error returned by muniverse API.
    """
    pass

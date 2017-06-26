"""
APIs to read and write raw objects in the protocol used by
the muniverse-bind command.
"""

import struct
import bson

from .errors import ProtoError

def read_object(pipe):
    """
    Read and decode a BSON object.
    """
    len_data = pipe.read(4)
    if len(len_data) != 4:
        raise ProtoError('EOF reading length field')
    length = struct.unpack('<I', len_data)[0]
    payload = pipe.read(length)
    if len(payload) != length:
        raise ProtoError('EOF reading BSON object')
    try:
        return bson.loads(payload)
    except ValueError as exc:
        raise ProtoError('BSON decode error: ' + str(exc))

def write_object(pipe, obj):
    """
    Encode and write a BSON object.
    """
    try:
        data = bson.dumps(obj)
    except ValueError as exc:
        raise ProtoError('BSON encode error: ' + str(exc))
    pipe.write(struct.pack('<I', len(data)))
    pipe.write(data)
    pipe.flush()

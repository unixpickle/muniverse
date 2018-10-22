"""
Bindings for the muniverse suite of Reinforcement Learning
environments.
"""

from .env import Env
from .spec import spec_for_name
from .actions import MouseAction, KeyAction, key_for_code

__all__ = ['Env', 'KeyAction', 'MouseAction', 'key_for_code', 'spec_for_name']

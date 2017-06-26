"""
Simple program to demonstrate how to use muniverse.
"""

import sys
sys.path.insert(0, '..')

import muniverse

def main():
    spec = muniverse.spec_for_name('KatanaFruits-v0')
    print(str(spec))

if __name__ == '__main__':
    main()

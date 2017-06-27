"""
Simple program to measure the framerate of muniverse on an
environment.
"""

import sys
sys.path.insert(0, '..')

import muniverse
import time

def main():
    print('Looking up environment...')
    spec = muniverse.spec_for_name('TowerMania-v0')
    print('Creating environment...')
    env = muniverse.Env(spec)
    try:
        print('Resetting environment...')
        env.reset()

        count = 60
        print('Timing ' + str(count) + ' frames...')
        start = time.time()
        for _ in range(0, count):
            env.step(0.1)
            env.observe()
        took = time.time() - start
        print('FPS: ' + str(count/took))

    finally:
        env.close()

if __name__ == '__main__':
    main()

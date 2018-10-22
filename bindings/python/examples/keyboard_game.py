"""
Simple program to demonstrate how to use muniverse on a
game that takes keyboard events.
"""

import sys

sys.path.insert(0, '..')
import muniverse  # noqa: E402


def main():
    print('Looking up environment...')
    spec = muniverse.spec_for_name('MinimalDots-v0')
    print('Creating environment...')
    env = muniverse.Env(spec)
    try:
        print('Resetting environment...')
        env.reset()

        print('Playing game...')
        action = muniverse.key_for_code('ArrowRight')
        while True:
            reward, done = env.step(0.1, action.with_event('keyDown'),
                                    action.with_event('keyUp'))
            print('reward: ' + str(reward))
            if done:
                break

    finally:
        env.close()


if __name__ == '__main__':
    main()

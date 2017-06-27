"""
Simple program to demonstrate how to use muniverse on a
game that takes mouse events.
"""

import sys
sys.path.insert(0, '..')

import muniverse
import numpy as np

def main():
    print('Looking up environment...')
    spec = muniverse.spec_for_name('TowerMania-v1')
    print('Creating environment...')
    env = muniverse.Env(spec)
    try:
        print('Resetting environment...')
        env.reset()

        print('Getting observation...')
        obs = env.observe()
        print(ascii_art(obs))

        print('Playing game...')
        step_idx = 0
        action = muniverse.MouseAction('mousePressed', x=100, y=100, click_count=1)
        actions = [action, action.with_event('mouseReleased')]
        while True:
            reward, done = env.step(0.1, actions[step_idx % 2])
            step_idx += 1
            print('reward: ' + str(reward))
            if done:
                break

    finally:
        env.close()

def ascii_art(img):
    brightness = np.sum(img, axis=2) / 3
    downsampled = brightness[::14, ::7]
    binary = downsampled > 128
    height, width = binary.shape
    res = ''
    for y in range(0, height):
        if res != '':
            res += '\n'
        for x in range(0, width):
            if binary[y, x]:
                res += 'X'
            else:
                res += ' '
    return res

if __name__ == '__main__':
    main()

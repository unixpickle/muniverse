"""
Simple program to demonstrate how to use muniverse.
"""

import sys
sys.path.insert(0, '..')

import muniverse

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
    finally:
        env.close()

def ascii_art(img):
    brightness = np.sum(buf, axis=2) / 3
    downsampled = brightness[::5, ::5]
    binary = downsampled > 128
    width, height = binary.shape
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

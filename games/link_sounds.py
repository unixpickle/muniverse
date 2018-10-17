"""
Link all ogg files to mp3s for Chrome compat.
"""

import os

for root, dirs, files in os.walk('downloaded_games'):
    for file in files:
        if file.endswith('.ogg'):
            mp3_path = os.path.join(root, file[:-4] + '.mp3')
            if not os.path.exists(mp3_path):
                os.symlink(file, mp3_path)

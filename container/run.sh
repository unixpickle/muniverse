#!/bin/bash

fsserver -path=/downloaded_games &

/chrome-linux/chrome \
  --headless \
  --mute-audio \
  --no-sandbox \
  --remote-debugging-port=9222 \
  --remote-debugging-address=0.0.0.0 \
  --disable-gpu \
  "$@" &

wait

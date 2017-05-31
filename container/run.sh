#!/bin/bash

fsserver -path=/downloaded_games &

/chrome-linux/chrome \
  --headless \
  --mute-audio \
  --no-sandbox \
  --remote-debugging-port=9222 \
  --remote-debugging-address=0.0.0.0 \
  --disable-gpu \
  --disable-webgl \
  "$@" &

# Use a socket close to trigger termination.
netwait -port 1337 -timeout 1m
kill -9 $(jobs -p)

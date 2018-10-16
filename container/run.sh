#!/bin/bash

fsserver -path=/downloaded_games &

chromium-browser \
  --headless \
  --mute-audio \
  --no-sandbox \
  --remote-debugging-port=9222 \
  --remote-debugging-address=0.0.0.0 \
  --disable-gpu \
  "$@" &

# Use a socket close to trigger termination.
netwait -port 1337 -timeout 1m
kill -9 $(jobs -p)

FROM ubuntu:16.04

# Install chromium-browser for its dependencies; it is an
# old (i.e. stable) version that we won't use.
RUN apt-get update \
  && apt-get install -y \
     unzip \
     chromium-browser \
     libgtk-3-dev \
     libosmesa-dev \
  && apt-get remove -y chromium-browser \
  && rm -rf /var/lib/apt/lists/*

# Chromium revision 474638.
# See https://github.com/scheib/chromium-latest-linux.
ENV chromeURL https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F474638%2Fchrome-linux.zip?alt=media
ADD $chromeURL /chrome.zip
RUN unzip -q /chrome.zip \
  && rm /chrome.zip \
  && ln -s /usr/lib/x86_64-linux-gnu/libOSMesa.so /chrome-linux/libosmesa.so

EXPOSE 9222

ENTRYPOINT /chrome-linux/chrome \
  --headless \
  --mute-audio \
  --no-sandbox \
  --remote-debugging-port=9222 \
  --remote-debugging-address=0.0.0.0

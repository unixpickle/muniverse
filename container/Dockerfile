FROM ubuntu:16.04

# Install webserver and chromium dependencies.
# We install chromium-browser for its dependencies;
# we don't use the package itself because it is old.
RUN apt-get update \
  && apt-get install -y \
     unzip \
     golang \
     git \
     netcat-openbsd \
     chromium-browser \
     libgtk-3-dev \
     libosmesa-dev \
  && mkdir /go \
  && export GOPATH=/go \
  && go get github.com/unixpickle/fsserver \
  && mv /go/bin/fsserver /usr/local/bin \
  && apt-get remove -y chromium-browser golang git \
  && rm -rf /var/lib/apt/lists/*

# Chromium revision 464640.
# See:
# https://www.chromium.org/getting-involved/download-chromium.
# https://github.com/scheib/chromium-latest-linux
# https://omahaproxy.appspot.com
RUN curl https://www.googleapis.com/download/storage/v1/b/chromium-browser-snapshots/o/Linux_x64%2F464640%2Fchrome-linux.zip?alt=media >/chrome.zip \
  && unzip -q /chrome.zip \
  && rm /chrome.zip \
  && ln -s /usr/lib/x86_64-linux-gnu/libOSMesa.so /chrome-linux/libosmesa.so

COPY run.sh /run.sh
COPY downloaded_games /downloaded_games
RUN chmod +x /run.sh

EXPOSE 9222 1337

ENTRYPOINT ["/run.sh"]

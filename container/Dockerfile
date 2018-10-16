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
  && mkdir /go \
  && export GOPATH=/go \
  && go get github.com/unixpickle/fsserver \
  && mv /go/bin/fsserver /usr/local/bin \
  && apt-get remove -y golang git \
  && rm -rf /var/lib/apt/lists/*

COPY run.sh /run.sh
COPY downloaded_games /downloaded_games
COPY netwait /go/src/netwait
RUN chmod +x /run.sh \
  && go build -o /usr/local/bin/netwait /go/src/netwait/main.go

EXPOSE 9222 1337

ENTRYPOINT ["/run.sh"]

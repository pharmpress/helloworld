FROM scratch

ENV HELLO_PORT :8099

ADD https://github.com/pharmpress/helloworld/releases/download/v0.0.1/helloworld-linux-amd64 /helloworld

ENTRYPOINT ["/helloworld"]

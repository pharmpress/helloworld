FROM busybox

ENV HELLO_PORT :8099

ADD https://github.com/pharmpress/helloworld/releases/download/v0.0.1/helloworld-linux-amd64 /helloworld

RUN chmod +x /helloworld

ENTRYPOINT ["/helloworld"]

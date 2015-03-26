FROM scratch

ENV HELLO_PORT :8099

ADD bin/helloworld-linux64-static /helloworld

ENTRYPOINT ["/helloworld"]

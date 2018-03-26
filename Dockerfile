# The dockerfile for hubimage
FROM alpine:3.5

# copy file
COPY bin/hubimage /hubimage

# work dir
WORKDIR /

# start application server
ENTRYPOINT ["/hubimage"]

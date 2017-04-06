FROM alpine
RUN apk --update add ca-certificates
ADD wundergram /
WORKDIR /
EXPOSE 8080
ENTRYPOINT ./wundergram

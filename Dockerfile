FROM alpine
RUN apk --update add ca-certificates
RUN apk add tzdata
ADD wundergram /
WORKDIR /
EXPOSE 8080
ENTRYPOINT ./wundergram

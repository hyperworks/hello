FROM gliderlabs/alpine
RUN apk --update add go
WORKDIR /
ADD main.go /
EXPOSE 8080
ENTRYPOINT ["go", "run", "/main.go"]

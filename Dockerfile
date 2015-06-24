FROM chakrit/dind
RUN apk add --update go

RUN mkdir -p /build
ADD main.go /build/main.go
ADD Dockerfile.run /build/Dockerfile

ENV IMAGE_NAME hello
WORKDIR /build
RUN cd /build && CGO_ENABLED=0 go build -a -ldflags '-s' /build/main.go

CMD docker build -t $IMAGE_NAME .

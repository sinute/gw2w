FROM golang:1.17.9-alpine3.15 as go-builder

WORKDIR /gw2w

COPY . ./

RUN go mod verify
RUN go build -o bin/gw2w cmd/main.go

FROM alpine:3.15

WORKDIR /gw2w

LABEL maintainer="Sinute <sinute@outlook.com>"

COPY --from=go-builder /gw2w/bin/gw2w ./

EXPOSE 7878
ENTRYPOINT [ "/gw2w/gw2w" ]

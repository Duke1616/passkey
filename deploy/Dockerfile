FROM golang:1.21.1-alpine as build
MAINTAINER luankz
WORKDIR /usr/src/backend
COPY . .
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go mod tidy && go build -a -o dist/passkey -ldflags="-w -s" main.go
FROM alpine:latest as deploy
WORKDIR /usr/src/backend
COPY --from=build /usr/src/backend/dist/passkey .
EXPOSE 8000
CMD ["start"]
ENTRYPOINT ["./passkey"]
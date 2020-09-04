FROM golang:1.15-alpine AS build

WORKDIR /go/src
ADD webServer.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webServer .

######################
FROM scratch AS prod
MAINTAINER syun@live.cn
LABEL info="golang-1.15, webServer, usage:http://{ip}:8080?env={envName}"

COPY --from=build /go/src/webServer /

CMD ["/webServer"]

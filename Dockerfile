FROM golang:1.11-stretch as builder

WORKDIR /

COPY . .

RUN go get -v 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM scratch

COPY --from=builder app /
COPY --from=builder config.json /

ENTRYPOINT [ "/app" ]

EXPOSE 9090
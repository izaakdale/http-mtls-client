FROM golang:1.20-alpine as builder
WORKDIR /
COPY go.mod .
RUN go mod download
COPY main.go .
RUN go build -o http-mtls-client .

FROM scratch
WORKDIR /bin
COPY --from=builder /http-mtls-client /bin
CMD [ "/bin/http-mtls-client" ]
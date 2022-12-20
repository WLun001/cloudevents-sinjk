FROM golang as builder
WORKDIR /sink
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/sink .

FROM alpine
RUN addgroup -S sink && adduser -S -D sink -G sink
USER sink
WORKDIR /home/sink
COPY --from=builder /bin/sink ./
EXPOSE 8080
ENTRYPOINT ["./sink"]

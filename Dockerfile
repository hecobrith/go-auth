FROM golang as builder

WORKDIR /home

COPY go.* ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
RUN apk update && apk --no-cache --update add ca-certificates

COPY --from=builder /home/go-auth /go-auth
RUN chmod a+x /go-auth

ENV PORT 8000

CMD ["./go-auth"]
FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache \
        build-base \
        git \
        git-lfs \
        ca-certificates

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
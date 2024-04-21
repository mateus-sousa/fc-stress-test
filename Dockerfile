FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o stresstest .

ENTRYPOINT ["./stresstest", "stress"]

CMD ["--url=", "--requests=0", "--concurrency=0"]
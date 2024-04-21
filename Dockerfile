FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o stresstest .

CMD [ "./stresstest stress --url=http://www.google.com --requests=10 --concurrency=2"]
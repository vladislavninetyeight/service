FROM golang:1.22.4 AS gobuild

WORKDIR /app

RUN apt-get -y install tzdata
ENV TZ=Europe/Moscow

COPY ../../.. .

EXPOSE 8080

RUN go mod download
RUN go build -o main ./cmd/main

CMD ["./main"]
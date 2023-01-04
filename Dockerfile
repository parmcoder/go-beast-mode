# Download base image ubuntu 20.04
FROM golang:1.18.2

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Update Ubuntu Software repository
RUN apt update

RUN apt-get update -y

RUN apt-get install -y gccgo

WORKDIR /go/app

COPY . .

RUN go mod download

RUN go build -compiler=gccgo -gccgoflags="-O3" -o server .

CMD ["./server"]
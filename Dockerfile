FROM gcc:latest as base

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Update Ubuntu Software repository
RUN apt-get update -y

RUN apt-get install -y gccgo

WORKDIR /go/app

COPY . .

RUN go mod tidy
RUN go mod download

RUN go build -compiler=gccgo -gccgoflags="-O3 -static-libgo" -o server .

FROM golang:1.17.13 as dev

WORKDIR /go/app

COPY --from=base /go/app/server /go/app/optimized

CMD ./optimized
FROM gcc:11.3

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Update Ubuntu Software repository
RUN apt update

RUN apt-get update -y

RUN apt-get install -y gccgo

RUN apt-get install gcc

WORKDIR /go/app

COPY . .

RUN go mod download

RUN go build -compiler=gccgo -gccgoflags="-v -O3" -o server .

CMD ["./server"]
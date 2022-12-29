# Download base image ubuntu 20.04
FROM ubuntu:20.04

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Update Ubuntu Software repository
RUN apt update

# Install nginx, php-fpm and supervisord from ubuntu repository
RUN apt install -y nginx php-fpm supervisor && \
    rm -rf /var/lib/apt/lists/* && \
    apt clean

RUN apt-get update -y

RUN apt-get install -y gccgo

# Copy start.sh script and define default command for the container
WORKDIR /go/app
COPY . .

RUN gccgo -O3 -o cool main.o

CMD ["./cool"]
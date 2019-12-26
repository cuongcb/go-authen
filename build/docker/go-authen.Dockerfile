FROM golang:1.13
WORKDIR /app

RUN apt-get update
RUN apt-get install -y apt-utils mariadb-client

CMD make prepare && make run
# CMD make run
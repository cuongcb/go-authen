FROM golang:1.13
WORKDIR /app

RUN echo `pwd`

RUN apt-get update
RUN apt-get install -y mariadb-client

CMD make run
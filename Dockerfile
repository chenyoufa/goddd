FROM alpine:latest


MAINTAINER CYF

RUN mkdir /app

WORKDIR /app

ADD consignment-service /app/consignment-servicce

CMD ["./consignment-service"]
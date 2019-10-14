FROM alpine:3.10.2

MAINTAINER inbox@sashko.me

RUN apk add git bash go musl-dev

RUN mkdir /storage

RUN go get github.com/boltdb/bolt/...

RUN cd /storage && git clone https://github.com/the-sashko/iMinistry tgBot && cd tgBot && /bin/bash scripts/install.sh

EXPOSE 80

CMD ["./storage/tgBot/scripts/run.sh"]

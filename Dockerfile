FROM golang:1.16-alpine

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN go build -o /godockerpostgres

EXPOSE 4000

CMD [ "/godockerpostgres" ]
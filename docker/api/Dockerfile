FROM golang:latest

COPY src/api /go/src/api

WORKDIR /go/src/api/

RUN apt-get update -y && apt-get install -y mariadb-client
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql

CMD ["sh", "./start_app.sh"]
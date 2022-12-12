FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

# change GOARCH for your os, yours will probably be amd64 - I am using a m1 pro macbook pro
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o app

EXPOSE 3000
ENTRYPOINT ["/var/www/app"]
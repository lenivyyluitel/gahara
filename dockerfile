FROM golang:1.17

WORKDIR /gahara
COPY . .

RUN go get -d -v ./...
RUN go build cmd/gahara/main.go

EXPOSE 8080

CMD [ "./main" ]

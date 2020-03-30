FROM golang:1.14-alpine

RUN mkdir /superheroesapi
WORKDIR /superheroesapi

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/superheroesapi

CMD ["/go/bin/superheroesapi"]
EXPOSE 8080
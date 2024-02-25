FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr server/*.go

EXPOSE 8000
CMD [ "./svr" ]
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./env/configuration.yaml ./env
RUN go mod download

COPY *.go ./

RUN go build -o /main 

EXPOSE 30

CMD [ "/main" ]


# FROM golang:latest
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . ./
# #ADD ./env/configuration.yaml ./src/rgc/env/configuration.yaml
# RUN go build -o main .
# #EXPOSE 30
# ENTRYPOINT [ "./main" ]
FROM golang:1.15.2-alpine3.12 as build

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN go build -o main .
FROM build as final

RUN mkdir /app

WORKDIR /app

COPY --from=build /build .

EXPOSE 3000

CMD ["/app/main"]
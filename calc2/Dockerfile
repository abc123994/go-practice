FROM docker.io/golang:1.21.6-alpine3.18 as base
WORKDIR /app

FROM docker.io/golang:1.21.6-alpine3.18 as build

WORKDIR /app

COPY . .

RUN go build main.go

FROM base as final

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/html/* /app/html/
CMD ["./main"]

EXPOSE 80
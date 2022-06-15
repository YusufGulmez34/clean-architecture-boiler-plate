FROM golang:1.18.3-alpine3.15 as builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY . .

RUN go mod download
RUN go build cmd/boiler-plate/main.go

FROM alpine

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.yaml /app/config.yaml
WORKDIR /app


CMD [ "./main" ]

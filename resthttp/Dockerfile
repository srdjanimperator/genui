# BUILD

FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /restapi ./resthttp/app.go

# DEPLOY 

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /restapi /restapi

EXPOSE 7001

USER nonroot:nonroot

ENTRYPOINT [ "/restapi" ]

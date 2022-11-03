# Build Stage
FROM golang:1.17-alpine3.16 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]
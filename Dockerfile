FROM golang:1.24 AS build
LABEL authors="Vitalie Balanici"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

FROM busybox

WORKDIR /app
COPY --from=build /app/app .
COPY --from=build /app/www ./www

EXPOSE 8080
CMD [ "/app/app" ]
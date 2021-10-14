FROM golang:1.17.2-alpine as build

WORKDIR /app
COPY . .
RUN go mod download && go build -o hackube

FROM golang:1.17.2-alpine as final
WORKDIR /app
COPY --from=build app/hackube .

ENTRYPOINT ["./hackube"]
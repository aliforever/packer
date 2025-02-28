FROM golang:alpine as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o packer ./cmd/packer

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/packer ./
COPY --from=builder /app/cmd/packer/public ./public

EXPOSE 8080

CMD [ "./packer" ]
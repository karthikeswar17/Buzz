FROM golang:1.21.4
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build
EXPOSE 8000
CMD ["./user"]
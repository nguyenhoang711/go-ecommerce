FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o dev.shoppenguin.com ./cmd/ecommerce

FROM scratch

COPY ./configs /configs

COPY --from=builder /build/dev.shoppenguin.com /

ENTRYPOINT [ "/dev.shoppenguin.com", "configs/local.yaml" ]
FROM golang:1.19-alpine AS builder

WORKDIR /usr/src/app

ENV GOPROXY=https://goproxy.cn

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY ./main.go .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server


FROM scratch as runner

COPY --from=builder /usr/src/app/server /opt/app/

CMD ["/opt/app/server"]


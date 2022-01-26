ARG GOLANG_VERSION=1.17
FROM golang:${GOLANG_VERSION} AS builder

RUN apt-get -qq update && apt-get -yqq install upx

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /src

COPY . .
RUN go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -tags 'osusergo netgo static_build' \
  -o /bin/with_coffee \
  ./main.go


RUN upx -q -9 /bin/with_coffee

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/with_coffee /bin/with_coffee
COPY lib/format/templates /lib/format/templates

ENTRYPOINT ["/bin/with_coffee"]
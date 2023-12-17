# https://github.com/GoogleContainerTools/distroless
FROM golang:1.21.5 AS builder
WORKDIR /go/src/app
COPY . .
RUN make setup
RUN make docker

FROM gcr.io/distroless/static-debian11
WORKDIR /root/
COPY --from=builder /go/src/app/bin/signet .
CMD [ "./signet" ]

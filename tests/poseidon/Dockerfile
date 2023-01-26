FROM golang:1.19-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git
ADD . /poseidon_cli
RUN go build -o /poseidon_cli/execute /poseidon_cli/execute.go


FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /poseidon_cli/execute /usr/local/bin/poseidon

ENTRYPOINT ["poseidon"]

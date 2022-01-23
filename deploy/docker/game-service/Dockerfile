FROM golang:1.17-alpine AS builder

ARG GITHUB_PATH=github.com/arttet/rock-paper-scissors-lizard-spock
ARG VERSION COMMIT_HASH

RUN apk add --update make git curl

COPY Makefile /home/${GITHUB_PATH}/Makefile
COPY go.mod /home/${GITHUB_PATH}/go.mod
COPY go.sum /home/${GITHUB_PATH}/go.sum
COPY pkg /home/${GITHUB_PATH}/pkg

WORKDIR /home/${GITHUB_PATH}
RUN make reqs
RUN make deps

COPY . /home/${GITHUB_PATH}
RUN make build


FROM alpine:latest as server

ARG GITHUB_PATH=github.com/arttet/rock-paper-scissors-lizard-spock
ARG SERVICE_NAME=game-service
LABEL org.opencontainers.image.source https://${GITHUB_PATH}

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/${SERVICE_NAME} .
COPY --from=builder /home/${GITHUB_PATH}/configs/${SERVICE_NAME}/config.yml .

RUN chown root:root ${SERVICE_NAME}
ENTRYPOINT ["./game-service"]

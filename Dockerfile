# Install development environment
FROM golang:1.12.2-alpine as builder
ENV GO111MODULE=on
RUN apk add --no-cache git nodejs npm make gcc musl-dev && \
	npm i -g typescript && \
	go install github.com/aerogo/pack && \
	go install github.com/aerogo/run && \
	go install golang.org/x/tools/cmd/goimports && \
	git clone --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn
WORKDIR /notify.moe

# Download dependencies when go.mod or go.sum changes
COPY go.mod go.sum ./
RUN go mod download

# Run Typescript compiler when scripts change
COPY ./tsconfig.json ./
COPY ./scripts ./scripts
RUN tsc

# Build
COPY . ./
RUN pack && \
	go build

# Production
FROM alpine:latest as production
RUN apk add --no-cache ca-certificates
COPY --from=builder /root/.aero /root/.aero
COPY --from=builder /notify.moe /notify.moe
ENV ARN_ROOT=/notify.moe
WORKDIR /notify.moe
ENTRYPOINT ["./notify.moe"]
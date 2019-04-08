# Development
FROM golang:1.12.2-alpine as builder
RUN apk add --no-cache git nodejs npm make gcc libc-dev
RUN npm i -g typescript
RUN git clone --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn
ENV GO111MODULE=on
RUN go install github.com/aerogo/pack && \
	go install github.com/aerogo/run && \
	go install golang.org/x/tools/cmd/goimports
RUN mkdir /notify.moe
ADD go.mod go.sum /notify.moe/
WORKDIR /notify.moe
RUN go mod download
ADD . /notify.moe
RUN tsc && \
	pack && \
	GOOS=linux go build

# Production
FROM alpine:latest as production
RUN apk add ca-certificates
COPY --from=builder /root/.aero /root/.aero
COPY --from=builder /notify.moe /notify.moe
ENV ARN_ROOT=/notify.moe
WORKDIR /notify.moe
CMD ["./notify.moe"]
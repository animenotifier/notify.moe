FROM golang:1.12.2-alpine as builder
RUN mkdir /notify.moe
ADD . /notify.moe
WORKDIR /notify.moe
ENV ARN_ROOT /notify.moe
RUN apk add git nodejs npm make gcc libc-dev
RUN npm i -g typescript
RUN go mod download
RUN make tools && make assets
RUN GOOS=linux go build
RUN git clone --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn

FROM alpine:latest as production
COPY --from=builder /root/.aero /root/.aero
COPY --from=builder /notify.moe /notify.moe
WORKDIR /notify.moe
CMD ["./notify.moe"]
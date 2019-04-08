FROM golang:1.12.2-alpine
RUN mkdir /notify.moe
ADD . /notify.moe
WORKDIR /notify.moe
ENV ARN_ROOT /notify.moe
RUN apk add --no-cache git nodejs npm make gcc libc-dev
RUN npm i -g typescript
RUN go mod download
RUN make all
RUN apk del git nodejs npm make gcc libc-dev
CMD ["/notify.moe/notify.moe"]
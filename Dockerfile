FROM golang:1.12.2-alpine
RUN mkdir /notify.moe
ADD . /notify.moe
WORKDIR /notify.moe
ENV ARN_ROOT /notify.moe
RUN apk add --no-cache git nodejs npm make gcc libc-dev
RUN npm i -g typescript
RUN go mod download
RUN make tools && make assets
RUN make server
RUN git clone --depth=1 https://github.com/animenotifier/database ~/.aero/db/arn
RUN apk del git nodejs npm make gcc libc-dev
CMD ["/notify.moe/notify.moe"]
FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/shuStudent
ENV GO111MODULE on
WORKDIR /go/src/shuStudent/cli
RUN go get && go build
WORKDIR /go/src/shuStudent/web
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/shuStudent/shuStudent /
WORKDIR /
CMD ./shuStudent
ENV PORT 8000
EXPOSE 8000
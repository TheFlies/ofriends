# build the server
FROM golang:alpine as gobuilder

RUN mkdir -p /build/src/github.com/TheFlies/ofriends
ENV GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV GOPATH=/build
WORKDIR /build/src/github.com/TheFlies/ofriends
ADD vendor ./vendor
ADD internal ./internal
ADD main.go .

RUN go build -v -o ofriends .
# build the web
FROM node:alpine as webbuilder
RUN mkdir /build
ADD ./web /build
WORKDIR /build
RUN yarn install
RUN yarn build
# now we are
FROM alpine
RUN adduser -S -D -H -h /app ofriends
COPY --from=gobuilder /build/src/github.com/TheFlies/ofriends/ofriends /app/
COPY --from=webbuilder /build/dist /app/web
WORKDIR /app
RUN chown -R ofriends /app
RUN chmod +x ofriends
EXPOSE 8080
USER ofriends
CMD ["./ofriends"]
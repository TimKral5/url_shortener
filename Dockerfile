FROM golang AS build
WORKDIR /src

COPY . .
RUN go build ./cmd/url_shortener

FROM alpine:latest AS alpine
WORKDIR /app

RUN apk add libc6-compat
COPY --from=build /src/url_shortener .
RUN chmod +x url_shortener

EXPOSE 3000
CMD ["/app/url_shortener"]


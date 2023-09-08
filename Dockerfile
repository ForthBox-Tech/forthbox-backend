FROM golang:1.17-alpine as builder

COPY . /opt/app/
RUN mkdir -p /opt/app/ && cd /opt/app && go mod tidy
RUN cd /opt/app/ && go build -o /opt/app/web cmd/server/main.go && chmod +x /opt/app/web


FROM alpine:edge

COPY --from=builder /opt/app/web /app/web
COPY --from=builder /opt/app/conf /app/conf

RUN apk add tzdata --no-cache && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone
ENV TZ=Asia/Shanghai

RUN cd /app/conf && cp app_dev.ini app.ini

WORKDIR /app
VOLUME /app/data
EXPOSE 8080

CMD ["/app/web"]


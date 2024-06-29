FROM alpine:latest

RUN mkdir /app

COPY exporterApp /app

CMD ["/app/exporterApp"]

# build a tiny  docker image
FROM alpine:latest
RUN mkdir /app
COPY mailApp /app
CMD ["/app/mailApp"]
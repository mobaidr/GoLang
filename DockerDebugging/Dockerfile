FROM golang:1.18-alpine
WORKDIr /app
EXPOSE 8090

COPY main.go ./
ENV GO111MODULE=off
RUN go build -o hello-app .

CMD [ "./hello-app" ]
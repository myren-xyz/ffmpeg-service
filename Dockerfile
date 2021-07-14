
FROM golang:alpine
RUN mkdir app
ADD . /app
WORKDIR /app
RUN go build .
EXPOSE 8080
CMD ["./ffmpeg-service"]
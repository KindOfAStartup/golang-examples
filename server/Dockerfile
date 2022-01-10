FROM golang:1.17.1
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8099
ARG APP_ENV local
RUN go build -o main .
CMD "/app/main"
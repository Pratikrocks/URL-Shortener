# Build Stage
# First pull Golang image
FROM golang:1.17-alpine as build-env

# Set environment variable
ENV APP_NAME url-shortener
ENV CMD_PATH main.go

# Copy application data into image
COPY . /src/$APP_NAME
WORKDIR /src/$APP_NAME

# Budild application
RUN CGO_ENABLED=0 go build -v -o $APP_NAME

# Run Stage
FROM alpine:3.14

# Set environment variable
ENV APP_NAME url-shortener

# Copy only required data into this image
COPY --from=build-env /src/$APP_NAME .

# Expose application port
EXPOSE 8082

# Start app
CMD ./$APP_NAME
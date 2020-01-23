# Build with command: docker build -t <appname> -f Dockerfile.m .
# Run with command: docker run -p 8080:8080 <appname>

### Build Stage ###

# Get latest golang version
FROM golang: latest as builder

# Make a new internal directory.
RUN mkdir /app

# Add files to the new directory.
ADD . /app

# Set the current internal working directory.
WORKDIR /app

# Build the app.
RUN go build -o main .

### END BUILD STAGE ###


### FINAL STAGE ###

# Get latest version of alpine.
FROM alpine:latest

# Set internal working directory.
WORKDIR /root/

# Copy the prebuilt binary from the previous stage.
COPY --from=builder /app/main .

# Open to the outside.
EXPOSE 8080

# Run the app.
CMD ["./main"]

### END FINAL STAGE ###

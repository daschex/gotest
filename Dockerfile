# build with command: docker build -t <appname> .
# run with command: docker run -p 8080:8080 <appname>

##### Build instructions

# Get latest version of golang
FROM golang:latest

# Make a new directory for our working files.
RUN mkdir /app

# Add the files from our root to the new directory.
ADD . /app

# Set the internal working directory to the newly created one.
WORKDIR /app

# Build the app.
RUN go build -o main .

# Run the built app.
CMD ["/app/main"]

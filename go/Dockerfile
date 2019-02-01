FROM golang:1.11.2

# Install packages
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp

# build the code
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 

# Expose the application on port 8081
EXPOSE 8081

# Set the entry point of the container 
CMD ["./main"]
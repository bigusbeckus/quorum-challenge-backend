FROM ubuntu:latest

# Install dependencies
RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y git curl make

# Download go
RUN curl -LO https://go.dev/dl/go1.21.3.linux-amd64.tar.gz

# Extract go
RUN tar -C /usr/local -xzvf go1.21.3.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
RUN go version

COPY . /app
WORKDIR /app
RUN go mod tidy
RUN make build

CMD [ "make", "run" ]

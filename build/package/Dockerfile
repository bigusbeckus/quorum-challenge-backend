FROM ubuntu:latest

# Install dependencies
RUN apt update -y
RUN apt upgrade -y
RUN apt install -y git curl make

# Download go
RUN curl -O https://go.dev/dl/go1.21.3.linux-amd64.tar.gz

# Extract go
RUN tar -xvf go1.21.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
RUN go version

COPY . /app
WORKDIR /app
RUN go mod tidy
RUN make build

CMD [ "make", "run" ]

FROM golang:1.4.2-wheezy


# -----------------
# Install dependencies
# -----------------

RUN apt-get update && apt-get install -y build-essential


# -----------------
# Install Go dependencies
# -----------------

RUN go get -u github.com/go-martini/martini

# -----------------
# Copy files over
# -----------------

RUN mkdir -p /src/unit/
ADD . /src/unit/
WORKDIR /src/unit/
RUN chmod +x /src/unit/start

# -----------------
# Set the server
# -----------------

EXPOSE 5000

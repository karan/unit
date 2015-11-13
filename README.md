# unit

![](https://raw.githubusercontent.com/karan/unit/master/_dna.jpg?token=ADHGIXpzysS9hkZ0sEW0XDuvDLHM8Tk3ks5WT3XPwA%3D%3D)

A highly modular, fast API framework in Go backed by [martini](https://github.com/go-martini/martini), and distributed as a Docker container for easy installation and deployment.

`unit` is designed to be a highly modular and extensible API. These modules or extensions are called *units* and live in `units/` folder. Each unit is part of the `package units`.

`units` are independent modules that can do anything you want them to do. Imagine having one that can draw an image, another that saves an image to s3, and another that notifies users - all exposed through your own HTTP API.

Simply drop your `units` in the `units/` directory, and restart (rebuild) the API. `unit` discovers and links the API, and installs the necessary dependencies automagically.

To get started, take a look the [installation](#installation) steps, and [units docs](#writing-units).

## Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Accessing the API](#accessing-the-api)
- [Writing Units](#writing-units)
  - [Examples](#examples)
- [Benchmark](#benchmark)

## Prerequisites

- [Docker](https://docs.docker.com/installation/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Docker Machine](https://docs.docker.com/machine/install-machine/) (Mac OS X)

## Installation

- `$ git clone git@github.com:karan/unit.git && cd unit/`
- `$ docker-compose up`

#### Mac & Docker Machine

If you're using Docker Machine, follow these instructions for installation:

```bash
# Provision the docker engine
$ docker-machine create --driver virtualbox unit

# Set the environment
$ eval "$(docker-machine env unit)"

# See the IP address of the host
$ docker-machine ip unit

# Start the container
$ docker-compose up
```

## Accessing the API

If using Docker Machine, run `$ docker-machine ls` to find out the VM IP address.

If your IP is `192.168.99.100`, load `192.168.99.100:5000` in the browser of the computer running docker to access the API.

## Writing units

TODO: A short tutorial on how to write unitss

#### Examples

TODO: Example units

## Benchmark

TODO: Add `wrk` benchmark

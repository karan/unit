# unit

A highly modular, extensible API framework in Go (Golang).

![](https://raw.githubusercontent.com/karan/unit/master/_dna.jpg?token=ADHGIXpzysS9hkZ0sEW0XDuvDLHM8Tk3ks5WT3XPwA%3D%3D)

- [Architecture](#architecture)
- [Installation](#installation)
- [Accessing the API](#accessing-the-api)
- [Writing Units](#writing-units)

## Architecture

`unit` is designed to be a highly modular and extensible API. These modules or extensions are called *units* and live in `units/` folder. Each unit is part of the `package units`.

Simply drop your units in the directory, and restart (rebuild) the API. `unit` discovers and links the API together automagically.

## Installation

- Install [Docker](https://docs.docker.com/installation/), [Docker Compose](https://docs.docker.com/compose/install/), and if using on a Mac, [Docker Machine](https://docs.docker.com/machine/install-machine/).
- Clone this repo.
- `docker-compose up`

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

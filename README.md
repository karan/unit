# unit
Modular API in Golang

### Installation

- Install [Docker](https://docs.docker.com/installation/), [Docker Compose](https://docs.docker.com/compose/install/), and if using on a Mac, [Docker Machine](https://docs.docker.com/machine/install-machine/).
- Clone this repo.
- `docker-compose up`

### Accessing the API

If using docker machine, run `$ docker-machine ls` to find out the VM IP address.

If your IP is `192.168.59.103`, load `192.168.59.103:5000` in the browser of the computer running docker to access the API.

### Architecture

`unit` is designed to be a highly modular and extensible API. These modules or extensions are called *units* and live in `units/` folder. Each unit is part of the `package units`.

Simply drop your units in the directory, and restart (rebuild) the API. `unit` discovers and links the API together automagically.

### Writing units

TODO: A short tutorial on how to write unitss

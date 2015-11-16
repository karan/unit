# unit

![](https://raw.githubusercontent.com/karan/unit/master/_dna.jpg?token=ADHGIXpzysS9hkZ0sEW0XDuvDLHM8Tk3ks5WT3XPwA%3D%3D)

A highly **modular**, **fast** API framework in Go backed by [martini](https://github.com/go-martini/martini), and distributed as a **Docker** container for easy installation and deployment.

`unit` is designed to be a highly modular and extensible API. These modules or extensions are called *units* and live in `units/` folder. Each unit is part of the `package units`.

`units` are independent modules that can do anything you want them to do. Imagine having one that can draw an image, another that saves an image to s3, and another that notifies users - all exposed through your own HTTP API.

Simply drop your `units` in the `units/` directory, and restart (rebuild) the API. `unit` discovers and links the API, and **installs the necessary dependencies** automagically.

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

Every unit must import at least these two:

```go
import "github.com/go-martini/martini"
import "./../unit"
```

For any third-party imports, mark them with the `//- unit-deps` comment so `unit` can discover them and install them for you.

Example:

```go
import (
  "github.com/go-martini/martini"

  "./../unit"

  "github.com/martini-contrib/render" //- unit-dep
)
```

The only thing needed after that is to register at least one group of routes like so:

```go
g := unit.Group(func(router martini.Router) {

  // Use martini like you would
  router.Get("/1", func() string {
    return "v2 - 1!"
  })

  router.Get("/2", func() string {
    return "v2 - 2!"
  })

})

// Register with the following path
g.Register("/v2")
```

The routes will then be available at `host/v2/1` and `host/v2/2`.

The most beautiful thing here is that you use martini to setup the routes like you normally would. You have full access to all features of martini - you can render a static page, or send JSON or send XML response. `unit` does not really care.

#### Example

[`units/plugin1.go`](https://github.com/karan/unit/blob/master/units/plugin1.go)
```go
package units

import (
  "github.com/go-martini/martini"

  "./../unit"
)

func init() {
  g := unit.Group(func(router martini.Router) {

    router.Get("/1", func() string {
      return "v2 - 1!"
    })

    router.Get("/2", func() string {
      return "v2 - 2!"
    })

  })

  g.Register("/v2")
}

```

## Benchmark

`unit` is fast as a cheetah.

```bash
$ ab -n 1000 -c 20 http://192.168.99.100:5000/v1/1
This is ApacheBench, Version 2.3 <$Revision: 1663405 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.99.100 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        192.168.99.100
Server Port:            5000

Document Path:          /v1/1
Document Length:        7 bytes

Concurrency Level:      20
Time taken for tests:   2.450 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      123000 bytes
HTML transferred:       7000 bytes
Requests per second:    408.15 [#/sec] (mean)
Time per request:       49.002 [ms] (mean)
Time per request:       2.450 [ms] (mean, across all concurrent requests)
Transfer rate:          49.03 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4  63.8      0    1167
Processing:     4   45  19.0     40     166
Waiting:        3   42  17.9     39     164
Total:          4   49  66.2     41    1213

Percentage of the requests served within a certain time (ms)
  50%     41
  66%     47
  75%     53
  80%     59
  90%     71
  95%     82
  98%     97
  99%    113
 100%   1213 (longest request)
```

*Inspired by the slower, much less user-friendly [modapi](https://github.com/csu/modapi)*

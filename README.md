### TweetoGo

TweetoGo is a simple Telegram Bot to fetch last 5 popular tweets for a given hashtag.


### Requirements

Before installing TweetoGo, you need to get one token from the `BotFather`, `access-token`, `access-token-secret`, `consumer-key` and `consumer-key` from Twitter APIs. 
Save theses token like this:

```shell
$ tree secrets
secrets
├── access-token
├── access-token-secret
├── consumer-key
├── consumer-secret
└── telegram-token
```

### Installation

You can install TweetoGo directly from sources or from Docker.

#### From binary: 

Get the latest [release](https://github.com/tormath1/tweetogo/releases)

#### From sources:

Please make sure that `go` and `dep` are installed: 

```shell
$ dep version
dep:
 version     : v0.4.1
 build date  : 2018-01-24
 git hash    : 37d9ea0a
 go version  : go1.9.1
 go compiler : gc
 platform    : linux/amd64
$ go version
go version go1.10.1 linux/amd64
```

Clone this repo and change your location: 

```shell
$ git clone https://github.com/tormath1/tweetogo.git
$ cd tweetogo/
$ dep ensure -vendor-only
$ mv /your/secrets /tmp
$ go run main.go
```

#### From Docker

Please make sure that `docker` and `docker-compose` are installed on your machine: 

```shell
$ docker version
Client:
 Version:       18.03.0-ce
 API version:   1.37
 Go version:    go1.10
 Git commit:    0520e24302
 Built: Fri Mar 23 01:47:41 2018
 OS/Arch:       linux/amd64
 Experimental:  false
 Orchestrator:  swarm

Server:
 Engine:
  Version:      18.03.0-ce
  API version:  1.37 (minimum version 1.12)
  Go version:   go1.10
  Git commit:   0520e24302
  Built:        Fri Mar 23 01:48:12 2018
  OS/Arch:      linux/amd64
  Experimental: false
$ docker-compose version
docker-compose version 1.21.0, build unknown
docker-py version: 3.2.1
CPython version: 3.6.4
OpenSSL version: OpenSSL 1.1.0h  27 Mar 2018
```

You can build directly on your machine, or simply pull image from Docker [Hub](https://hub.docker.com): 

```shell
$ docker build -t name:tag .
$ docker pull tormath1/tweetogo:0.0.1
```

Fire up your containers with `docker-compose` !

```shell
$ mv /your/secrets /tmp
$ docker-compose up -d 
```
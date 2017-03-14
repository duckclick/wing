[![CircleCI](https://circleci.com/gh/duckclick/wing.svg?style=svg)](https://circleci.com/gh/duckclick/wing)

# Wing

## Running

Install and start redis server
```sh
brew install redis
redis-server
```

```sh
go run main.go
# or
# CONFIG=path/to/application.yml go run main.go
```

## Running tests

`make test`

## Development

__Wing__ uses `glide` for dependency management, to setup the project do the following:

```sh
brew install glide
```

__Note__: Read here (https://github.com/Masterminds/glide) for a different OS

```sh
glide install
```

To add new dependecies do:

```sh
glide get <dependency>
```

Build a docker image:

```sh
docker build -t wing .
```

Run a docker image and mount the configuration file:

```sh
docker run --rm -t -p 7273 -v ${PWD}/application.yml:/go/src/github.com/duckclick/wing/application.yml wing
```

## Generate self signed certificate

Generate private key (.key)

```sh
# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

```sh
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

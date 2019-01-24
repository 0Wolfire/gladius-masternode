# Gladius Master Node

The Gladius Master Node software

## Overview
This application runs as a the main connection point in the beta network. All
connections are routed through one master node, and then the client is connected
to the geographically nearest edge node.


# Building

## Build the first stage of the docker container (proxygen environment)
```shell
docker build --target proxygen-env -t gladiusio/proxygen-env .
```

## Build the second stage of the docker container (builds the masternode binary)
```shell
docker build --target masternode-builder -t gladiusio/masternode-builder .
```

## Build the full container that runs the masternode
```shell
docker build -t gladiusio/masternode .
```

## Build the development container
```shell
docker build -f develop.Dockerfile -t gladiusio/masternode-develop
```

## Run the development container
```shell
docker-compose -f develop-compose.yml run --name devenv dev bash
```

## Debug tests with development container (once inside it)
```shell
make check
cd tests
gdb ./masternode_tests
```

## Copy library headers to your host machine for IntelliSense purposes (optional)
```shell
docker cp <container id>:/usr/local/include <path on host to put header files>
docker cp <container id>:/usr/include <path on host to put header files>
```
Then configure your IDE to include the path you copied the headers to.

## Build the Proxygen docs (optional)
```shell
docker exec -it <container id> /bin/bash
sudo apt-get install doxygen
cd /proxygen
doxygen Doxyfile

(from another host terminal session)
docker cp <container id>:/proxygen/html <path on host to put docs>
```

## Running the Software
```shell
docker run -p 80:80 -it gladiusio/masternode
```

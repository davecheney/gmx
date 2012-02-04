# Go management extensions

## Installation
  
	go get github.com/davecheney/gmx

## Getting started

Integrating gmx into your application is as simple as importing the package in your main package via the side effect operator.

	package main

	import _ "github.com/davecheney/gmx"

By default gmx opens a unix socket in /tmp, the name of the socket is

	/tmp/.gmx.$PID.$VERSION

## Protocol version 0

The current protocol version is 0, which is a simple JSON based protocol. You can communicate with the gmx socket using a tool like socat.

	socat UNIX-CONNECT:/tmp/.gmx.12345.0 stdin
     
Requests are the names of registered keys, the results are json encoded

	% socat UNIX-CONNECT:/tmp/.gmx.9328.0 stdin
	runtime.gomaxprocs 
	1

The names of registered keys are registered with the key `registry`. If there is no value registered then the json encoding of nil will be returned.

For convenience a client is included in the gmxc sub directory. Please consult the README in that directory for more details.

## Registering gmx keys

New keys can be registered using the `Register` function

	gmx.Register(key string, f func() interface{})

`f` can be any function that returns a json encodable result. `f` is executed whenever its key is invoked, responsibility for ensuring the function is thread safe is the responsibility of the programmer.

## Runtime instrumentation

By default gmx registers a number of values under the runtime key, refer to the runtime.go source for more details.

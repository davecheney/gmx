Go management extensions
========================

Installation
------------
  
     go get github.com/davecheney/gmx


Getting started
---------------

Integrating GMX into your application is as simple as importing the package in your main package

     package main

     import _ "github.com/davecheney/gmx"

By default GMX opens a unix socket in /tmp, the name of the socket is

     /tmp/.gmx.$PID.$VERSION


Protocol version 0
------------------

The current protocol version is 0, which is a simple JSON based protocol. This protocol only supports getting registered values. You can communicate with the GMX socket using a tool like socat.

     socat UNIX-CONNECT:/tmp/.gmx.12345.0 stdin
     
Requests are the names of registered keys, the results are json encoded

     % socat UNIX-CONNECT:/tmp/.gmx.9328.0 stdin
     runtime.gomaxprocs 
     1

The names of registered keys are registered with the key


     registry
     ["runtime.memstats","runtime.gomaxprocs","registry"]


Registering GMX entries
-----------------------

New entries can be registerd using the Register function

     gmx.Register(name string, f func() interface{})

f can be any function that returns a json encodable result. 

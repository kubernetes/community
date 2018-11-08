# Unix Socket Support for kubectl

Status: Pending

Version: GA

Implementation Owner: @choo-stripe

## Motivation

Currently kubectl does not expose a way to send the request over a unix socket. Curl provides a --unix-socket option, as do many other cli utilities. Stripe uses unix sockets as a secured proxy into some infrastructure and it would be great to be able to send reqs from kubectl over this unix socket.

## Proposal


A --unix-socket flag will be added to kubectl, much like in curl. The curl docs state:
(HTTP) Connect through this Unix domain socket, instead of using the network.

This is precisely what the kubectl unix-socket flag will be used for.

## User Experience

A new global option to kubectl will be added (--unix-socket)

### Use Cases

Issuing commands to any kube api server whereby unix socket is the preferred way to access it.

kubectl --unix-socket ~/.myunixsocket --server myserver.com get pods

## Implementation

Already wrote an implementation:
https://github.com/kubernetes/kubernetes/compare/master...choo-stripe:choo/unix-socket

### Client/Server Backwards/Forwards compatibility

There should be no compatibility issues. New clients will still be compatible with old servers and vice-versa.

## Alternatives considered

We could allow users to designate the network AND address that get passed into the net.Dial function used to connect to the server. That would clutter the cli interface. Taking this approach keeps it more consistant with curl and other familiar unix tools.

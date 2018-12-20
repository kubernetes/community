# Unix Socket Support for kubectl

Status: Pending

Version: GA

Implementation Owner: @choo-stripe

## Motivation

Currently kubectl does not expose a way to send the request over a unix socket. Curl provides a --unix-socket option, as do many other cli utilities. Stripe uses unix sockets as a secured proxy into some infrastructure and it would be great to be able to send reqs from kubectl over this unix socket.

## Proposal


A --via-unix-socket flag will be added to kubectl, much like the --unix-socket in curl. The curl docs state:
(HTTP) Connect through this Unix domain socket, instead of using the network.

This is precisely what the kubectl via-unix-socket flag will be used for.

## User Experience

A new global option to kubectl will be added (--via-unix-socket)

via-unix-socket will also be supported in kubeconfig files, e.g.:

    clusters:
    - cluster:
        server: mykubecluster.com
        via-unix-socket: /tmp/.unix-socket-proxy

### Use Cases

Issuing commands to any kube api server whereby unix socket is the preferred way to access it.

kubectl --via-unix-socket ~/.myunixsocket --server mykubecluster.com get pods

## Implementation

Already wrote an implementation:
[https://github.com/kubernetes/kubernetes/compare/master...choo-stripe:choo/unix-socket](https://github.com/kubernetes/kubernetes/compare/master...choo-stripe:choo/unix-socket)

### Client/Server Backwards/Forwards compatibility

There should be no compatibility issues. New clients will still be compatible with old servers and vice-versa.

## Alternatives considered

- We could allow users to designate the network AND address that get passed into the net.Dial function used to connect to the server. That would clutter the cli interface. Taking this approach keeps it more consistent with curl and other familiar unix tools.
- We could use a different name, other than "via-unix-socket". Perhaps "proxy-socket" would be better. Unfortunately, we cannot use --unix-socket because it is already a flag to kubectl proxy. We could also use "--via-unix-socket" on the cli and just "unix-socket" in kubeconfig files. I opted to use "via-unix-socket" to keep both config syntax unified.

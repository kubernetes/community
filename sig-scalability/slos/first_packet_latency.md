## Time To First Packet SLI details

### Definition

| Status | SLI |
| --- | --- |
| WIP | First Packet Latency in milliseconds (ms) from the client initiating the TCP connection to a Service (sending the SYN packet) to the client receiving the first packet from the Service backend (typically the SYN-ACK packet in the three-way handshake) measured as 99th percentile over last 5 minutes aggregated across all the node instances.|

### User stories

- As a user of vanilla Kubernetes, I want some guarantees on how quickly my pods can connect
to the service backends

### Other notes

First Packet Latency is a more user-centric metric than just the full connection establishment time. It reflects the initial perceived delay.  A fast First Packet Latency makes your application feel fast, even if the full handshake takes a bit longer.

### How to measure the SLI.

Requires precise timestamps for when the client sends the SYN packet and when it receives the first packet from the server. This can be done:

- Client-side: In the application code or using a benchmark application.
- Network devices: Packet inspection and analysis on nodes along the network datapath.

### Caveats

Important Considerations:

- Network Latency: geographic distance, routing, and network congestion.
- Other traffic on the network can delay the SYN-ACK, even if the server responds quickly.
- Client-side processing and network conditions on the client side can also introduce minor delays.
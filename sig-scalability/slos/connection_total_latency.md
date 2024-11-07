## Connection Total Latency SLI details

### Definition

| Status | SLI |
| --- | --- |
| WIP | The time elapsed in seconds (s) or minutes (min) from the successful establishment of a TCP connection to a Kubernetes service to the connection being closed measured as 99th percentile over last 5 minutes aggregated across all the node instances.|

### User stories

- As a user of vanilla Kubernetes, I want some visibility on how longs my pods are connected
to the services

### Other notes

The total connection duration can help to understand how clients interact with services, optimize resource usage, and identify potential issues like connection leaks or excessive short-lived connections.

### How to measure the SLI.

Requires precise timestamps for when the client sends the SYN packet and when it receives the last packet from the server. This can be done:

- Client-side: In the application code or using a benchmark application.
- Network devices: Packet inspection and analysis on nodes along the network datapath.

### Caveats

Important Considerations:

- Network Latency: geographic distance, routing, and network congestion.
- How quickly the server can process the SYN packet and send the SYN-ACK also contributes to the first packet latency.
- Other traffic on the network can delay the SYN-ACK, even if the server responds quickly.
- Client-side processing and network conditions on the client side can also introduce minor delays.
## Throughput

### Definition

| Status | SLI |
| --- | --- |
| WIP |The rate of successful data transfer over a TCP connection to services, measured in bits per second (bps), kilobits per second (kbps), megabits per second (Mbps), or gigabits per second (Gbps) measured as 99th percentile over last 5 minutes aggregated across all the connections to services in a node.|

### User stories

- As a user of vanilla Kubernetes, I want some visibility to ensure my applications meet my performance requirements when connection to services
- As a user of vanilla Kubernetes, I want to understan when my applications meet my performance requirements when connection to services

### Other notes

The aggregated throughput help to understand if the cluster network and applications can handle the required data transfer rates and to identify any bottlenecks limiting throughput.

### How to measure the SLI.

Requires tto collect both the time duration of the connection and the amount of data transferred during that time. This can be done:

- Client-side: In the application code or using a benchmark application.
- Network devices: Packet inspection and analysis on nodes along the network datapath.


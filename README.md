# Adaption of assignment_demo_2023

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This is an amended version of the demo and template shared for backend assignment of 2023 TikTok Tech Immersion.

## Installation

Requirement:

- golang 1.18+
- docker

To install dependency tools:

```bash
make pre
```

## Run

```bash
docker-compose up -d
```

Check if it's running:

```bash
curl localhost:8080/ping
```
## Implementation Approach 
_Implementaton is in progress._

Backend Database: Redis

**Rough Flow**
1. Client makes HTTP call to server. 
2. Server makes RPC call to rpc-server in local network.
3. rpc-server receives request (`SendRequest()` / `PullRequest()`) and processes it. 
4. Type of request determines read/write from/to Redis backend.
5. Data is sent back to HTTP server as a response to the client's request.

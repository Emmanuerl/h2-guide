# HTTP/2: A Comprehensive Guide for Developers

## 🚀 Introduction to HTTP/2

HTTP/2 is the second major version of the HTTP protocol, published in **May 2015** as [RFC 7540](https://datatracker.ietf.org/doc/html/rfc7540). It was developed by the IETF HTTP Working Group, based on **Google’s SPDY** protocol, to address performance limitations in HTTP/1.1.

HTTP/2 is **binary, multiplexed, and more efficient**. It improves web performance, reduces latency, and optimizes the way data is transferred between clients and servers.

---

## 🔁 HTTP/2 vs. HTTP/1.1 (and HTTP/1.0)

See Diagram [here](https://github.com/Emmanuerl/h2-guide/blob/main/COMPARE.md)

| Feature                    | HTTP/1.0 | HTTP/1.1             | HTTP/2                     |
| -------------------------- | -------- | -------------------- | -------------------------- |
| Protocol Type              | Text     | Text                 | **Binary**                 |
| Multiplexing               | ❌       | ❌ (pipelining only) | ✅ Yes (real multiplexing) |
| Head-of-line Blocking      | ❌       | ✅ Yes               | ✅ Reduced (at transport)  |
| Header Compression         | ❌       | ❌                   | ✅ HPACK                   |
| Connection Reuse           | ❌       | ✅ Persistent        | ✅ Single TCP connection   |
| Server Push                | ❌       | ❌                   | ✅ Supported               |
| TLS Requirement (Browsers) | ❌       | ❌                   | ✅ (effectively required)  |
| Prioritization             | ❌       | ❌                   | ✅ Supported               |

---

## ⚙️ Core Features of HTTP/2

### 1. **Binary Framing Layer**

HTTP/2 uses a binary format instead of human-readable text. This allows more efficient parsing and error handling.

### 2. **Multiplexing**

Multiple requests and responses are interleaved on a single TCP connection without blocking each other.

### 3. **Stream Prioritization**

Clients can assign weights and dependencies to streams, allowing better bandwidth utilization.

### 4. **Header Compression (HPACK Algo.)**

Reduces overhead by compressing HTTP headers using a dynamic table and indexed representation.

### 5. **Server Push**

The server can send assets (like CSS/JS) before the client explicitly requests them.

---

## 🌍 Applications of HTTP/2

- **Web Performance Optimization (arguably)**: Faster page load times due to multiplexing and header compression.
- **Mobile & Low-bandwidth Environments**: Improved performance in high-latency or constrained networks.
- **APIs & gRPC**: gRPC is built on HTTP/2, enabling multiplexed, bi-directional streaming RPCs.
- **CDNs & Reverse Proxies**: Better connection handling under high traffic loads.

---

## ⚠️ Challenges and Problems in HTTP/2

### 1. **Head-of-line Blocking at TCP Layer**

Although HTTP/2 allows multiple streams, they still share a single TCP connection. If one packet is lost, **all streams** are delayed. This is mitigated in **HTTP/3** using QUIC (UDP).

### 2. **Server Push Misuse**

While powerful, server push is complex to implement well and can waste bandwidth if used poorly.

### 3. **Debugging Complexity**

Binary framing makes debugging harder without proper tooling (e.g., `nghttp`, Wireshark).

### 4. **TLS Overhead**

While not mandated by the spec, all major browsers require HTTPS for HTTP/2, which increases handshake complexity (though mitigated by TLS 1.3).

### 5. **Proxy & Middleware Compatibility**

Some proxies, middleboxes, or firewalls do not support HTTP/2 well, which can lead to downgraded or broken connections.

---

## 📊 Performance Implications

- **Reduced Latency**: Especially noticeable on high-latency connections.
- **Improved Page Load Times**: Fewer connections and less blocking.
- **Fewer TCP Connections**: Reduces server load and improves scalability.

However, **over-optimization** (e.g., bundling too many resources) can backfire under HTTP/2 since multiplexing eliminates the need for those hacks.

---

## 🔐 HTTP/2 & Security

- **TLS with ALPN** is required by browsers.
- Many servers support **Opportunistic Encryption** for non-browser clients.
- HTTP/2 does not inherently offer more security than HTTP/1.1, but encourages secure transport.

---

## Stream States

In HTTP/2, communication happens over **streams** — independent, bidirectional channels within a single TCP connection. Each stream represents a single request/response exchange.

Each stream progresses through well-defined **states** during its lifetime. These states help the client and server manage the lifecycle and resources efficiently.

---

## The 7 Stream States (per [RFC 7540 §5.1](https://datatracker.ietf.org/doc/html/rfc7540#section-5.1))

| State                    | Description                                                                     |
| ------------------------ | ------------------------------------------------------------------------------- |
| **Idle**                 | Stream is created but no frames have been sent or received yet.                 |
| **Reserved (Local)**     | The local endpoint has reserved the stream (e.g., via `PUSH_PROMISE`).          |
| **Reserved (Remote)**    | The remote endpoint has reserved the stream (e.g., via `PUSH_PROMISE`).         |
| **Open**                 | The stream is active; both sides can send frames (`HEADERS`, `DATA`, etc.).     |
| **Half-Closed (Local)**  | Local endpoint has sent `END_STREAM` flag but can still receive frames.         |
| **Half-Closed (Remote)** | Remote endpoint has sent `END_STREAM` flag but local can still send frames.     |
| **Closed**               | The stream is closed; no further frames can be sent or received on this stream. |

---

## Frame Types

All frames start with a 9-byte header specifying length, type, flags, and stream identifier. Frames are typically stream specific. Some frames apply to the entire connection (e.g., SETTINGS, PING), others to specific streams (e.g., DATA, HEADERS).

| Frame Type    | Type Code | Description                                                                                  |
| ------------- | --------- | -------------------------------------------------------------------------------------------- |
| DATA          | 0x0       | Carries arbitrary data (e.g., HTTP request or response body).                                |
| HEADERS       | 0x1       | Contains header fields (HTTP metadata) for requests or responses.                            |
| PRIORITY      | 0x2       | Specifies stream priority and dependency information.                                        |
| RST_STREAM    | 0x3       | Immediately terminates a stream with an error code.                                          |
| SETTINGS      | 0x4       | Negotiates parameters affecting communication, e.g., max concurrent streams.                 |
| PUSH_PROMISE  | 0x5       | Reserved a stream for server push, promising to send headers later.                          |
| PING          | 0x6       | Used for measuring round-trip time and checking connection liveness.                         |
| GOAWAY        | 0x7       | Informs the peer to stop creating new streams; initiates connection shutdown.                |
| WINDOW_UPDATE | 0x8       | Implements flow control by increasing the window size for streams or connection.             |
| CONTINUATION  | 0x9       | Used to continue sending header block fragments if they don't fit in a single HEADERS frame. |

---

# 😎 Flaunting HTTP/2: gRPC

gRPC is a modern, high-performance RPC (Remote Procedure Call) framework. Developed by Google, it enables client and server applications to communicate transparently. Uses **Protocol Buffers** (protobuf) as the interface definition language and data serialization format.

## Key Concepts in gRPC

- **Service Definition**: Define services and methods in `.proto` files using protobuf syntax.
- **Stub/Client**: Client-side object that provides the methods to call the remote service.
- **Server**: Implements the service interface and handles incoming calls.
- **Streaming**: Supports unary RPCs (single request/response) and streaming RPCs (client, server, or bidirectional streams).
- **Deadlines/Timeouts**: Allows setting time limits on RPC calls.
- **Metadata**: Key-value pairs sent with calls for authentication, tracing, etc.

## How gRPC Relates to HTTP/2

gRPC **uses HTTP/2 as its transport protocol** by default!
gRPC’s design tightly couples with HTTP/2, enabling efficient, low-latency communication ideal for microservices, mobile, and web apps.

---

## ✅ Summary

HTTP/2 brings massive improvements to the web stack: faster, more efficient, and better suited to modern workloads. However, it introduces complexity — understanding frames, streams, and flow control is essential for those building performance-critical systems.

HTTP/2 is not a magic bullet but a powerful protocol upgrade when deployed correctly — especially when paired with TLS 1.3, good cache strategies, and modern server configurations.

---

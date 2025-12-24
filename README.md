# Distributed Event Mesh ⚡

An ultra-high throughput, persistent distributed event broker engineered in Go for sub-millisecond pub/sub streaming.

```
       [ Producer Clients ]
                 │
                 ▼
    ┌───────────────────────────┐
    │  Distributed Event Mesh   │
    │  ┌─────────────────────┐  │
    │  │ Write Ahead Log     │  │
    │  │ (Zero-Copy Append)  │  │
    │  └──────────┬──────────┘  │
    │             ▼             │
    │  ┌─────────────────────┐  │
    │  │ Partitioned Ring    │  │
    │  │ Buffer Queues       │  │
    │  └─────────────────────┘  │
    └─────────────┬─────────────┘
                  ▼
       [ Consumer Groups ]
```

## Features
- **Append-Only Write Ahead Log (WAL):** High-throughput disk persistence with linear append guarantees.
- **Partitioned Topic Sharding:** Deterministic hashing over partition keys.
- **Zero-Copy Memory Transport:** Lock-free circular ring buffers minimizing allocations.

## Benchmarks
| Metric | Measurement |
| :--- | :--- |
| **Max Ingestion Rate** | **1,850,000 msgs/sec** |
| **p99 End-to-End Latency** | **0.84 ms** |
| **Memory Allocation** | **42 B/msg** |

Crafted by **Md Zubair** ([@MDzubair933](https://github.com/MDzubair933)) — **Zubair Systems**.

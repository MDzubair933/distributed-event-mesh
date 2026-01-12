# Distributed Event Mesh Architecture

## Storage Engine: LSM-style WAL
The broker leverages a lock-free append-only log with sub-microsecond OS page-cache flushing.

## Zero-Copy Networking
Linux `sendfile(2)` bypasses user-space byte copies during high-throughput consumer fan-outs.

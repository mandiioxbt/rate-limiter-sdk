# Rate Limiter SDK

Distributed rate limiter: token bucket, sliding window, leaky bucket with Redis backend.

## Algorithms
- Token bucket (burst-friendly)
- Sliding window (precise count)
- Leaky bucket (smooth output)
- Redis-backed distributed mode

## Usage
```go
limiter := ratelimit.New(client, ratelimit.SlidingWindow(100, time.Minute))
if limiter.Allow("user-123") { handleRequest() }
```

## License: MIT

# Rate Limiter Sdk

Distributed rate limiter: sliding window, token bucket, leaky bucket.

## Algorithms
- Token bucket (burst-friendly)
- Sliding window (precise)
- Leaky bucket (smooth)
- Redis-backed distributed mode

## Usage
```go
limiter := ratelimit.New(redisClient, ratelimit.SlidingWindow(100, time.Minute))
if limiter.Allow(userID) {
    handleRequest()
}
```

## License
MIT

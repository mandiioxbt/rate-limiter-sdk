package ratelimiter
import ("sync"; "time")
type SlidingWindow struct {
    mu sync.Mutex; windows map[string][]time.Time; limit int; interval time.Duration
}
func New(limit int, interval time.Duration) *SlidingWindow {
    return &SlidingWindow{windows: make(map[string][]time.Time), limit: limit, interval: interval}
}
func (sw *SlidingWindow) Allow(key string) bool {
    sw.mu.Lock(); defer sw.mu.Unlock()
    now := time.Now(); cutoff := now.Add(-sw.interval)
    valid := []time.Time{}
    for _, t := range sw.windows[key] { if t.After(cutoff) { valid = append(valid, t) } }
    if len(valid) < sw.limit { valid = append(valid, now); sw.windows[key] = valid; return true }
    sw.windows[key] = valid; return false
}

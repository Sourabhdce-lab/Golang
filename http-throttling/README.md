# HTTP Requests Throttling
Use this package in HTTP server implementaion to rate limit HTTP requests. 

## Why rate limiting
- To save server going out of resources due to un-manageable number of requests. 
- To avoid DoS attacks

# Install
```bash
go get -u -v github.com/sourgoyal/golang/http-throttling
```

# Usage
```bash
import(throttling "github.com/sourgoyal/golang/http-throttling")

// Application must configure 'rate' and 'burstSize'
ConfigureLimiter(r rate.Limit, b int)

// To be called by application to check rate limit
LimitRate(next http.Handler) http.Handler
```

# North Star 
- Rate limiting per client on the basis of HTTP client IP address
- Rate limitting for particular methods
- Allow runtime rate limit configuration

# License
MIT License
let RateLimiter = require('limiter').RateLimiter

arn.cacheLimiter = new RateLimiter(1, 250)
arn.networkLimiter = new RateLimiter(1, 500)
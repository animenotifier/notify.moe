'use strict'

let RateLimiter = require('limiter').RateLimiter

arn.cacheLimiter = new RateLimiter(1, 350)
arn.networkLimiter = new RateLimiter(1, 1000)
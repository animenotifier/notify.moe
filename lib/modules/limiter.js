'use strict'

let RateLimiter = require('limiter').RateLimiter

arn.cacheLimiter = new RateLimiter(1, 500)
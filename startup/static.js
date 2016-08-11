// Service worker routes
app.sendFile('service-worker.js', 'worker/service-worker.js')
app.sendFile('cache-polyfill.js', 'worker/cache-polyfill.js')
package main

// loggerConfig := zap.NewDevelopmentConfig()
// loggerConfig.OutputPaths = append(loggerConfig.OutputPaths, "logs/server.log")

// logger, _ := loggerConfig.Build()

// logTime := func(ctx *aero.Context, next func()) {
// 	start := time.Now()
// 	next()
// 	responseTime := time.Since(start)

// 	if ctx.StatusCode == 200 {
// 		logger.Info(ctx.URI(), zap.Duration("responseTime", responseTime), zap.Int("statusCode", ctx.StatusCode))
// 	} else {
// 		logger.Warn(ctx.URI(), zap.Duration("responseTime", responseTime), zap.Int("statusCode", ctx.StatusCode))
// 	}
// }

// app.Use(logTime)

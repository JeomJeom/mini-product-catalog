package middleware

import (
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// MemoryUsageLogger is a Gin middleware that logs Go runtime memory statistics
// for each incoming HTTP request. It measures memory allocation, total allocation,
// system memory used, and the number of garbage collections triggered.
func MemoryUsageLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Collect memory statistics from the Go runtime
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		// Log key memory usage metrics after request is processed
		log.Printf("[MEMORY] Alloc = %v KB | TotalAlloc = %v KB | Sys = %v KB | NumGC = %v | Path = %s | Duration = %s",
			memStats.Alloc/1024,      // currently allocated heap objects
			memStats.TotalAlloc/1024, // total bytes allocated (even if freed)
			memStats.Sys/1024,        // total memory obtained from the OS
			memStats.NumGC,           // number of completed garbage collection cycles
			c.FullPath(),             // the route path for the request
			time.Since(start),        // how long the request took
		)
	}
}

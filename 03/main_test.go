package main

import (
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		main()
	}
	elapsed := time.Since(start)
	t.Logf("average speed %s", elapsed/100)
}

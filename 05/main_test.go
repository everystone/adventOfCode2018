package main

import (
	"testing"
	"time"
)

func TestM(t *testing.T) {
	main()
}
func TestPerformance(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		main()
	}
	elapsed := time.Since(start)
	t.Logf("average speed %s", elapsed/100)
}

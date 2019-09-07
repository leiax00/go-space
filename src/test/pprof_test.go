package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPprof(t *testing.T) {
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
	fmt.Print("Hello pprof...")
	count := 0
	for true {
		if count%2 == 0 {
			count++
		} else {
			count--
		}
	}
}

package restclient_test

import (
	rest "github.com/arielsrv/golang-toolkit/restclient"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp := Rb.Get("/user")
		if resp.StatusCode != http.StatusOK {
			log.Print("f[" + strconv.Itoa(i) + "] Status != OK (200)")
		}
	}
}

func BenchmarkCacheGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp := Rb.Get("/cache/user")

		if resp.StatusCode != http.StatusOK {
			log.Print("f[" + strconv.Itoa(i) + "] Status != OK (200)")
		}
	}
}

func BenchmarkSlowGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp := Rb.Get("/slow/user")

		if resp.StatusCode != http.StatusOK {
			log.Print("f[" + strconv.Itoa(i) + "] Status != OK (200)")
		}
	}
}

func BenchmarkSlowConcurrentGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rb.ForkJoin(func(cr *rest.Concurrent) {
			for j := 0; j < 100; j++ {
				cr.Get("/slow/user")
			}
		})
	}
}

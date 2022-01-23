package singleton

import (
	"testing"

	"github.com/golang-collections/lib.go/assert"
)

func TestLazySingletonGetInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("[BenchmarkGetLazyInstance] Fail")
			}
		}
	})
}

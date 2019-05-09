package test

import (
	"github.com/mercadolibre/myml/src/api/services/myml"
	"testing"
)

func BenchmarkAPI(b *testing.B) {

	for n := 0; n < b.N; n++ {
		myml.GetGeneralInfo(1234567)
	}
}

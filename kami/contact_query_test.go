package main

import "testing"
import "net/http/httptest"
import "encoding/json"

// func BenchmarkContactQueryAll(b *testing.B) {
// 	cq := NewContactQuery(1, 10)
// 	for i := 0; i < b.N; i++ {
// 		cq.All()
// 	}
// }

func BenchmarkContactQueryAllAndEncoding(b *testing.B) {
	cq := NewContactQuery(1, 10)

	for i := 0; i < b.N; i++ {
		res := cq.All()
		json.NewEncoder(httptest.NewRecorder()).Encode(res)
	}
}

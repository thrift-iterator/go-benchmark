package lstore

import (
	"testing"
	"github.com/thrift-iterator/go-benchmark/thrift/media"
	"encoding/json"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go-benchmark"
)

func Benchmark(b *testing.B) {
	m := &media.MediaContent{}
	err := json.Unmarshal([]byte(go_benchmark.M1), m)
	if err != nil {
		b.Error(err)
	}
	buf := thrift.NewTMemoryBuffer()
	proto := thrift.NewTBinaryProtocol(buf, true, true)
	m.Write(proto)
	bytes := buf.Bytes()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(bytes)
		m.Read(proto)
	}
}
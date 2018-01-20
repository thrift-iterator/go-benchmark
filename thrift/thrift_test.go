package thrift

import (
	"testing"
	"github.com/thrift-iterator/go-benchmark/thrift/media"
	"encoding/json"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/thrift-iterator/go-benchmark"
	"github.com/thrift-iterator/go"
	"github.com/v2pro/wombat/generic"
)

func Benchmark_thrift(b *testing.B) {
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

func Benchmark_thrifter_dynamic(b *testing.B) {
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
	decoder := thrifter.Config{
		Protocol:       thrifter.ProtocolBinary,
		DynamicCodegen: true,
	}.Froze().NewDecoder(nil, nil)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(bytes)
		decoder.Reset(nil, bytes)
		decoder.Decode(m)
	}
}

func Benchmark_thrifter_static(b *testing.B) {
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
	generic.DynamicCompilationEnabled = true
	decoder := api.NewDecoder(nil, nil)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		buf.Write(bytes)
		decoder.Reset(nil, bytes)
		decoder.Decode(m)
	}
}

package protobuf

import (
	"testing"
	"encoding/json"
	"github.com/thrift-iterator/go-benchmark"
	"github.com/stretchr/testify/require"
	"fmt"
)

func Test(t *testing.T) {
	should := require.New(t)
	m := &MediaContent{}
	err := json.Unmarshal([]byte(go_benchmark.M1), m)
	should.NoError(err)
	output, err := m.Marshal()
	should.NoError(err)
	fmt.Println(output)
}

func Benchmark(b *testing.B) {
	m := &MediaContent{}
	err := json.Unmarshal([]byte(go_benchmark.M1), m)
	if err != nil {
		b.Error(err)
	}
	output, err := m.Marshal()
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m.Unmarshal(output)
	}
}

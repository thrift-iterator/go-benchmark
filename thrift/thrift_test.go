package lstore

import (
	"testing"
	"github.com/thrift-iterator/go-benchmark/thrift/media"
	"encoding/json"
	"git.apache.org/thrift.git/lib/go/thrift"
)

var m1 = `{
	"media": {
		"uri": "http://javaone.com/keynote.mpg",
		"title": "Javaone Keynote",
		"width": 640,
		"height": 480,
		"format": "video/mpg4",
		"duration": 18000000,    
		"size": 58982400,        
		"bitrate": 262144,       
		"persons": ["Bill Gates", "Steve Jobs스"],
		"player": "JAVA",
		"copyright": null
	},
	"images": [
		{
			"uri": "http://javaone.com/keynote_large.jpg",
			"title": "Javaone Keynote",
			"width": 1024,
			"height": 768,
			"size": "LARGE"
		},
		{
			"uri": "http://javaone.com/keynote_small.jpg",
			"title": "Javaone Keynote",
			"width": 320,
			"height": 240,
			"size": "SMALL"
		}
	]
}`

func Benchmark(b *testing.B) {
	m := &media.MediaContent{}
	err := json.Unmarshal([]byte(m1), m)
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
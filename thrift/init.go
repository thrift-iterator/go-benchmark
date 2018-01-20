package thrift

import (
	"github.com/v2pro/wombat/generic"
	"github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go-benchmark/thrift/media"
)

var api = thrifter.Config{
	Protocol: thrifter.ProtocolBinary,
	IsFramed: false,
}.Froze()

//go:generate go install github.com/v2pro/wombat/cmd/wombat-codegen
//go:generate $GOPATH/bin/wombat-codegen -pkg github.com/thrift-iterator/go-benchmark/thrift
func init() {
	generic.Declare(func() {
		api.WillDecodeFromBuffer(
			(*media.MediaContent)(nil),
		)
	})
}
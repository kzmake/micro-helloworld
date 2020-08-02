// +build tools

package main

import (
	_ "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/micro/micro/v2"
	_ "github.com/micro/micro/v2/cmd/protoc-gen-micro"
)

//go:generate go install github.com/golang/protobuf/proto
//go:generate go install github.com/golang/protobuf/protoc-gen-go
//go:generate go install github.com/micro/micro/v2
//go:generate go install github.com/micro/micro/v2/cmd/protoc-gen-micro

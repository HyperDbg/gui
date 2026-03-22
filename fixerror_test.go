package main

import (
	"testing"

	"github.com/ddkwork/golibrary/std/stream"
)

func TestFixError(t *testing.T) {
	stream.Fmt(".")
	stream.Fix(".")
}

package main

import (
	"testing"
	"unsafe"
)

var code = "Hello World"
var slice = []byte("Hello World 2")

func BenchmarkStringToBytesSafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice = []byte(code)
	}
}

func BenchmarkStringToBytesUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice = *(*[]byte)(unsafe.Pointer(&code))
	}
}

func BenchmarkBytesToStringSafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		code = string(slice)
	}
}

func BenchmarkBytesToStringUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		code = *(*string)(unsafe.Pointer(&slice))
	}
}

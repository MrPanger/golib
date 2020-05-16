package util

import (
	"bytes"
	"testing"
)

func TestStringToSlice(t *testing.T) {
	str := "Hello World"
	if !bytes.Equal(StringToSlice(str), []byte("Hello World")) {
		t.Fail()
	}
}

func TestSliceToString(t *testing.T) {
	sli := []byte("Hello World")
	if SliceToString(sli) != "Hello World" {
		t.Fail()
	}
}

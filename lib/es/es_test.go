package es

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	err := PutMetadata("hello", 1, 12, "abc")
	if err != nil {
		panic(err)
	}

	meta, err := getMetadata("hello", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(meta)

	metas, err := SearchAllVersions("hello", 0, 1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(metas)
}

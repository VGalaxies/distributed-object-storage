package json

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

type hit struct {
	Source Metadata `json:"_source"`
}

type total struct {
	Relation string
	Value    int
}

type searchResult struct {
	Hits struct {
		Total total
		Hits  []hit
	}
}

func Test(t *testing.T) {
	bytes, err := ioutil.ReadFile("hits.json")
	if err != nil {
		panic(err)
	}
	var sr searchResult
	err = json.Unmarshal(bytes, &sr)
	if err != nil {
		panic(err)
	}

	bytes, err = ioutil.ReadFile("source.json")
	if err != nil {
		panic(err)
	}
	var hit hit
	err = json.Unmarshal(bytes, &hit)
	if err != nil {
		panic(err)
	}
}

package j

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

type codeResponse struct {
	Tree     *codeNode `json:"tree"`
	Username string    `json:"username"`
}

type codeNode struct {
	Name     string      `json:"name"`
	Kids     []*codeNode `json:"kids"`
	CLWeight float64     `json:"cl_weight"`
	Touches  int         `json:"touches"`
	MinT     int64       `json:"min_t"`
	MaxT     int64       `json:"max_t"`
	MeanT    int64       `json:"mean_t"`
}

var codeStruct codeResponse

func init() {
	f, err := os.Open("testdata/code.json.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(gz)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &codeStruct); err != nil {
		panic("unmarshal code.json: " + err.Error())
	}
}

func BenchmarkJSON(b *testing.B) {
	enc := json.NewEncoder(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		if err := enc.Encode(&codeStruct); err != nil {
			b.Fatal("Encode:", err)
		}
	}
}

func BenchmarkJSONBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&codeStruct); err != nil {
			b.Fatal("Encode:", err)
		}
		ioutil.Discard.Write(buf.Bytes())
	}
}

func BenchmarkJSONmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf, err := json.Marshal(&codeStruct)
		if err != nil {
			b.Fatal("Marshal:", err)
		}
		ioutil.Discard.Write(buf)
	}
}

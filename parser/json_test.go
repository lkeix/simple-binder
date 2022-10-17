package parser

import "testing"

func TestJsonParse(t *testing.T) {
	jp := NewJsonParser()
	src := `{ "test": "test" }`
	jp.Parse(([]byte)(src))
}

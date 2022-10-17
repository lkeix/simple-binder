package parser

import "fmt"

type tree struct {
	name     string
	value    interface{}
	parent   *tree
	children map[string]*tree
}

type JsonParser struct {
	src  []byte
	root *tree
}

type stack []byte

func (stk stack) top() byte {
	if len(stk) > 0 {
		return stk[len(stk)-1]
	}
	return 'a'
}

func (stk *stack) push(b byte) {
	*stk = append(*stk, b)
}

func (stk *stack) pop() byte {
	top := stk.top()
	*stk = (*stk)[:len(*stk)-1]
	return top
}

func NewJsonParser() Parser {
	return &JsonParser{}
}

func (jp *JsonParser) Parse(src []byte) {
	var stk stack
	start := 0
	isValue := false
	for i, byt := range src {
		if byt == '{' {
			stk.push(byt)
			continue
		}

		if byt == ':' {
			// TODO
			isValue = true
		}

		if byt == '"' && stk.top() == '"' {
			fmt.Printf("isValue = %v\n", isValue)
			fmt.Println(string(src[start:i]))
			stk.pop()
			continue
		} else if byt == '"' && stk.top() != '"' {
			start = i + 1
			stk.push(byt)
			continue
		}

		if stk.top() == '"' {

		}

		if byt == '}' && stk.top() != '{' {
			panic("error")
		}

		if byt == '}' && stk.top() == '}' {
			stk.pop()
			isValue = false
			continue
		}

		if byt == ',' {
			// TODO
			isValue = false
		}
	}
}

func (jp *JsonParser) Marshal(src interface{}) ([]byte, error) {
	return []byte("aa"), nil
}

func (jp *JsonParser) Unmarshal(src []byte, dst interface{}) error {
	return nil
}

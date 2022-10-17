package parser

type Parser interface {
	Parse([]byte)
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

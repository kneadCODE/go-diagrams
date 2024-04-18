package diagrams

import (
	"github.com/kneadCODE/go-diagrams/diagram"
)

func New(opts ...diagram.Option) (*diagram.Diagram, error) {
	return diagram.New(opts...)
}

package parser

import (
	"github.com/hawoond/hawoond/internal/utils/parser/json"
	"github.com/hawoond/hawoond/internal/utils/parser/structure"
)

type Parser struct {
	Json   json.Json
	Struct structure.Struct
}

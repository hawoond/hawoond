package hawoond

import (
	"github.com/hawoond/hawoond/internal/constant"
	"github.com/hawoond/hawoond/internal/fiber/fiberkit"
	"github.com/hawoond/hawoond/internal/utils/converter"
	"github.com/hawoond/hawoond/internal/utils/converter/checker"
	"github.com/hawoond/hawoond/internal/utils/converter/time"
	"github.com/hawoond/hawoond/internal/utils/parser"
)

var Fiberkit fiberkit.Fiber = fiberkit.Fiber{}
var Converter converter.Converter
var Checker checker.Checker
var TimeTo time.TimeTo
var Constant constant.Constant
var Parser parser.Parser

func init() {
	Fiberkit = fiberkit.Fiber{}
	Converter = converter.Converter{
		ByteTo:    converter.ByteTo{},
		ComplexTo: converter.ComplexTo{},
		Float32To: converter.Float32To{},
		Float64To: converter.Float64To{},
		IntTo:     converter.IntTo{},
		Int8To:    converter.Int8To{},
		Int16To:   converter.Int16To{},
		Int32To:   converter.Int32To{},
		Int64To:   converter.Int64To{},
		UintTo:    converter.UintTo{},
		Uint8To:   converter.Uint8To{},
		Uint16To:  converter.Uint16To{},
		Uint32To:  converter.Uint32To{},
		Uint64To:  converter.Uint64To{},
		RuneTo:    converter.RuneTo{},
		StringTo:  converter.StringTo{},
	}
	Checker = checker.Checker{}
	Constant = constant.Constant{}
	TimeTo = time.TimeTo{}
	Parser = parser.Parser{}
}

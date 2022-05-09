package utils

import (
	"github.com/davecgh/go-spew/spew"
)

const prettyPrintIndentation = "\t"

func PrettyPrint(arg any) string {
	cfg := spew.NewDefaultConfig()
	cfg.Indent = prettyPrintIndentation
	cfg.SortKeys = true
	cfg.DisableMethods = false
	cfg.MaxDepth = 20
	cfg.DisablePointerAddresses = true
	return cfg.Sdump(arg)
}

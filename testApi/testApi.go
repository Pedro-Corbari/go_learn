package main

import (
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
)

func main() {

	config.Default.Secret = apiSecret
	convertapi.ConvDef("pdf", "txt",
		param.NewPath("File", "C:\\Users\\pedro\\Downloads\\Path.pdf", nil),
	).ToPath("C:\\Users\\pedro\\Documents")
}

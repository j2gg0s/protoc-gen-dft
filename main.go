package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"github.com/j2gg0s/protoc-gen-dft/module"
)

func main() {
	pgs.
		Init().
		RegisterModule(module.NewDefaultor()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}

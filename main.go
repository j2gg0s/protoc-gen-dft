package main

import (
	"github.com/j2gg0s/protoc-gen-dft/module"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.
		Init().
		RegisterModule(module.NewDefaultor()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}

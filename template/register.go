package template

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func New(params pgs.Parameters) *template.Template {
	tpl := template.New("dft.gotpl")

	fns := goSharedFuncs{pgsgo.InitContext(params)}
	tpl.Funcs(map[string]interface{}{
		"context": dftContext,
		"render":  render(tpl),

		"msgTyp":   fns.msgTyp,
		"name":     fns.Name,
		"oneof":    fns.oneofTypeName,
		"accessor": fns.accessor,
		"pkg":      fns.PackageName,
	})

	template.Must(tpl.ParseFiles("template/dft.gotpl"))

	return tpl
}

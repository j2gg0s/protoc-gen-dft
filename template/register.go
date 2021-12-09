package template

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func New(params pgs.Parameters) *template.Template {
	tpl := template.New("go")

	fns := goSharedFuncs{pgsgo.InitContext(params)}
	tpl.Funcs(map[string]interface{}{
		"context": rulesContext,
		"render":  render(tpl),

		"msgTyp":   fns.msgTyp,
		"name":     fns.Name,
		"oneof":    fns.oneofTypeName,
		"accessor": fns.accessor,
		"pkg":      fns.PackageName,
	})

	template.Must(tpl.Parse(fileTpl))

	template.Must(tpl.New("none").Parse(noneTpl))
	template.Must(tpl.New("number").Parse(numTpl))
	template.Must(tpl.New("string").Parse(strTpl))
	template.Must(tpl.New("numbers").Parse(numsTpl))
	template.Must(tpl.New("strings").Parse(strsTpl))

	return tpl
}

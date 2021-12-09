package module

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"github.com/j2gg0s/protoc-gen-dft/template"
)

type Defaultor struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func NewDefaultor() pgs.Module {
	return &Defaultor{ModuleBase: &pgs.ModuleBase{}}
}

const (
	defaultorName = "defaultor"
)

func (m *Defaultor) Name() string { return defaultorName }

func (m *Defaultor) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
	m.ctx = pgsgo.InitContext(ctx.Parameters())
}

func (m *Defaultor) Execute(
	targets map[string]pgs.File,
	pkgs map[string]pgs.Package,
) []pgs.Artifact {
	module := m.Parameters().Str("module")

	tpl := template.New(m.Parameters())
	for _, f := range targets {
		m.Push(f.Name().String())

		out := m.ctx.OutputPath(f).SetExt(".dft.go")
		outPath := strings.TrimLeft(
			strings.ReplaceAll(out.String(), module, ""), "/")
		m.AddGeneratorTemplateFile(outPath, tpl, f)

		m.Pop()
	}
	return m.Artifacts()
}

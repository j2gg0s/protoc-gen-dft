package template

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"github.com/j2gg0s/protoc-gen-dft/dft"
)

type DftContext struct {
	Field pgs.Field
	Typ   string
	Val   interface{}
}

func dftContext(f pgs.Field) (DftContext, error) {
	var dftVal dft.FieldDft
	_, err := f.Extension(dft.E_Dft, &dftVal)
	if err != nil {
		return DftContext{}, err
	}
	out := DftContext{
		Field: f,
	}

	sep, trim := ",", " "
	switch r := dftVal.GetDft().(type) {
	case *dft.FieldDft_Number:
		out.Typ, out.Val = "number", r.Number
	case *dft.FieldDft_String_:
		out.Typ, out.Val = "string", r.String_
	case *dft.FieldDft_Numbers:
		if len(r.Numbers.GetSep()) > 0 {
			sep = r.Numbers.GetSep()
		}
		if len(r.Numbers.GetTrim()) > 0 {
			trim = r.Numbers.GetTrim()
		}
		numbers := []float64{}
		for _, number := range strings.Split(r.Numbers.GetValues(), sep) {
			n, err := strconv.ParseFloat(strings.Trim(number, trim), 64)
			if err != nil {
				return out, err
			}
			numbers = append(numbers, n)
		}
		out.Typ, out.Val = "numbers", numbers
	case *dft.FieldDft_Strings:
		if len(r.Strings.GetSep()) > 0 {
			sep = r.Strings.GetSep()
		}
		if len(r.Strings.GetTrim()) > 0 {
			trim = r.Strings.GetTrim()
		}
		strs := []string{}
		for _, str := range strings.Split(r.Strings.GetValues(), sep) {
			strs = append(strs, strings.Trim(str, trim))
		}
		out.Typ, out.Val = "strings", strs
	case *dft.FieldDft_Json:
		out.Typ, out.Val = "json", r.Json
	default:
		out.Typ = "none"
	}

	return out, nil
}

func render(tpl *template.Template) func(ctx DftContext) (string, error) {
	return func(ctx DftContext) (string, error) {
		var b bytes.Buffer
		err := tpl.ExecuteTemplate(&b, ctx.Typ, ctx)
		return b.String(), err
	}
}

type goSharedFuncs struct{ pgsgo.Context }

func (fns goSharedFuncs) accessor(ctx DftContext) string {
	return fmt.Sprintf("m.Get%s()", fns.Name(ctx.Field))
}

func (fns goSharedFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}
func (fns goSharedFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f)).Pointer()
}

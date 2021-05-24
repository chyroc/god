package god_template

import (
	"bytes"
	"text/template"
)

type Template struct {
	tmpl string
}

func New(tmpl string) *Template {
	return &Template{
		tmpl: tmpl,
	}
}

func (r *Template) BuildMap(vars map[string]interface{}) (string, error) {
	t, err := template.New("tmpl").Parse(r.tmpl)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, vars)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (r *Template) BuildInterface(vars interface{}) (string, error) {
	t, err := template.New("tmpl").Parse(r.tmpl)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, vars)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

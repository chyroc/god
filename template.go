package god

import (
	"bytes"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

type Template struct {
	tmpl         string
	isText       bool
	textTemplate *textTemplate.Template
	htmlTemplate *htmlTemplate.Template
	err          error
}

func NewTemplate(tmpl string) *Template {
	textTemplate, err := textTemplate.New("tmpl").Parse(tmpl)
	if err != nil {
		return &Template{err: err}
	}
	htmlTemplate, err := htmlTemplate.New("tmpl").Parse(tmpl)
	if err != nil {
		return &Template{err: err}
	}
	return &Template{
		tmpl:         tmpl,
		isText:       true,
		textTemplate: textTemplate,
		htmlTemplate: htmlTemplate,
	}
}

func (r *Template) WithHtml() *Template {
	r.isText = false
	return r
}

func (r *Template) Build(vars interface{}) (string, error) {
	if r.err != nil {
		return "", r.err
	}

	buf := &bytes.Buffer{}
	var err error
	if r.isText {
		err = r.textTemplate.Execute(buf, vars)
	} else {
		err = r.htmlTemplate.Execute(buf, vars)
	}
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

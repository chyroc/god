package test

import (
	"testing"

	"github.com/chyroc/god/god_template"
)

func Test_god_template_text(t *testing.T) {
	as := NewAssert(t)

	t.Run("", func(t *testing.T) {
		r := god_template.New("{{ .Name }}")
		s, err := r.BuildMap(map[string]interface{}{
			"Name": "a",
		})
		as.Nil(err)
		as.Equal("a", s)
	})

	t.Run("", func(t *testing.T) {
		r := god_template.New("{{ .Name }}")
		s, err := r.BuildMap(map[string]interface{}{
			"name": "a",
		})
		as.Nil(err)
		as.Equal("<no value>", s)
	})

	t.Run("", func(t *testing.T) {
		r := god_template.New("{{ .Name }}")
		s, err := r.BuildMap(nil)
		as.Nil(err)
		as.Equal("<no value>", s)
	})

	t.Run("", func(t *testing.T) {
		r := god_template.New("{{ Name }}")
		_, err := r.BuildMap(map[string]interface{}{
			"Name": "a",
		})
		as.NotNil(err)
		as.Equal(`template: tmpl:1: function "Name" not defined`, err.Error())
	})

	t.Run("", func(t *testing.T) {
		r := god_template.New("{{ .Name }}")
		s, err := r.BuildInterface(struct {
			Name string
		}{
			Name: "a",
		})
		as.Nil(err)
		as.Equal("a", s)
	})

	t.Run("", func(t *testing.T) {
		r := god_template.New("{{ Name }}")
		_, err := r.BuildInterface(struct {
			Name string
		}{
			Name: "a",
		})
		as.NotNil(err)
		as.Equal(`template: tmpl:1: function "Name" not defined`, err.Error())
	})
}

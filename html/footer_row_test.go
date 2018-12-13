package html_test

import (
	"testing"
	"github.com/elliotchance/gedcom/html"
	"github.com/stretchr/testify/assert"
)

func TestFooterRow_String(t *testing.T) {
	component := html.NewFooterRow().String()

	assert.Contains(t, component, "<div class=\"row\">")
	assert.Contains(t, component,
		"Generated with <a href=\"https://github.com/elliotchance/gedcom\">github.com/elliotchance/gedcom</a>")
}
package template_pattern

import "testing"

func TestTemplate(t *testing.T) {
	xihonsgji := &Xihongshi{}
	doCook(xihonsgji)

	dojidan := &Chaojidan{}
	doCook(dojidan)
}

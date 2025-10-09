package specw

import (
	"encoding/json"
	"fmt"

	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v3"

	"image/color"
	"strings"
)

type Color struct {
	Color color.Color

	Raw string
}

func (c *Color) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, got %q", n.Kind)
	}

	return c.UnmarshalString(n.Value)
}

func (c *Color) UnmarshalString(value string) error {
	c.Raw = value

	if strings.HasPrefix(value, "#") {
		hc, err := hexToRGBA(value)
		if err != nil {
			return err
		}
		c.Color = *hc
		return nil
	}

	colorName := strings.ToLower(value)

	cc, ok := colornames.Map[colorName]
	if ok {
		c.Color = cc
		return nil
	}

	return fmt.Errorf("unknown color %q", colorName)
}

func (c *Color) UnmarshalBinary(data []byte) error {
	return c.UnmarshalString(string(data))
}

func (c *Color) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return c.UnmarshalString(s)
}

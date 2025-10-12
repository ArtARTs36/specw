package specw

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v3"

	"image/color"
	"strings"
)

type Color struct {
	Color color.RGBA
}

func (c *Color) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, got %q", n.Kind)
	}

	return c.UnmarshalString(n.Value)
}

func (c *Color) UnmarshalString(value string) error {
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

func (c Color) MarshalYAML() (interface{}, error) {
	return c.Hex(), nil
}

func (c Color) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(c.Hex())), nil
}

func (c *Color) MarshalBinary() ([]byte, error) {
	return []byte(c.Hex()), nil
}

func (c *Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.Color.R, c.Color.G, c.Color.B)
}

// AsEEE sets color #eeeeee.
func (c *Color) AsEEE() {
	c.Color = color.RGBA{
		R: 238, //nolint:mnd // not need
		G: 238, //nolint:mnd // not need
		B: 238, //nolint:mnd // not need
		A: 255, //nolint:mnd // not need
	}
}

func hexToRGBA(hex string) (*color.RGBA, error) {
	const (
		minHexLength = 3
		maxHexLength = 6
	)

	hex = strings.TrimPrefix(hex, "#")

	if len(hex) < minHexLength {
		return nil, errors.New("color is too short")
	}
	if len(hex) > maxHexLength {
		return nil, errors.New("color is too long")
	}

	if len(hex) == minHexLength {
		hex = fmt.Sprintf("%s%s%s%s%s%s",
			string(hex[0]), string(hex[0]),
			string(hex[1]), string(hex[1]),
			string(hex[2]), string(hex[2]),
		)
	}

	values, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return nil, err
	}

	return &color.RGBA{
		R: uint8(values >> 16),         //nolint:mnd // not need
		G: uint8((values >> 8) & 0xFF), //nolint:mnd // not need
		B: uint8(values & 0xFF),        //nolint:mnd // not need
		A: 255,                         //nolint:mnd // not need
	}, nil
}

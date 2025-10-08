package specw

import (
	"errors"
	"fmt"

	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v3"

	"image/color"
	"strconv"
	"strings"
)

type Color struct {
	Color color.Color

	Raw string
}

type HexColor struct {
	Color color.Color

	Raw string
}

func (c *Color) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, got %q", n.Kind)
	}

	c.Raw = n.Value

	if strings.HasPrefix(n.Value, "#") {
		hc, err := hexToRGBA(n.Value)
		if err != nil {
			return err
		}
		c.Color = *hc
		return nil
	}

	colorName := strings.ToLower(n.Value)

	cc, ok := colornames.Map[colorName]
	if ok {
		c.Color = cc
		return nil
	}

	return fmt.Errorf("unknown color %q", colorName)
}

func (c *HexColor) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, got %q", n.Kind)
	}

	return c.FromHex(n.Value)
}

func (c *HexColor) FromHex(hex string) error {
	rgba, err := hexToRGBA(hex)
	if err != nil {
		return err
	}

	c.Raw = hex
	c.Color = *rgba

	return nil
}

func hexToRGBA(hex string) (*color.RGBA, error) {
	const (
		minHexLength = 4
		maxHexLength = 7
	)

	if len(hex) < minHexLength {
		return nil, errors.New("short hex string")
	}
	if len(hex) > maxHexLength {
		return nil, errors.New("long hex string")
	}

	values, err := strconv.ParseUint(hex[1:], 16, 32)
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

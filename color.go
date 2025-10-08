package specw

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"image/color"
	"strconv"
)

type HexColor struct {
	Color color.RGBA
}

func (c *HexColor) UnmarshalYAML(n *yaml.Node) error {
	var v string

	if err := n.Decode(&v); err != nil {
		return err
	}

	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, got %q", n.Kind)
	}

	return c.FromHex(v)
}

func (c *HexColor) FromHex(hex string) error {
	const (
		minHexLength = 4
		maxHexLength = 7
	)

	if len(hex) < minHexLength {
		return errors.New("short hex string")
	}
	if len(hex) > maxHexLength {
		return errors.New("long hex string")
	}

	values, err := strconv.ParseUint(hex[1:], 16, 32)
	if err != nil {
		return err
	}

	c.Color = color.RGBA{
		R: uint8(values >> 16),         //nolint:mnd // not need
		G: uint8((values >> 8) & 0xFF), //nolint:mnd // not need
		B: uint8(values & 0xFF),        //nolint:mnd // not need
		A: 255,                         //nolint:mnd // not need
	}

	return nil
}

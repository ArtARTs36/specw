package specw

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"image/color"
	"strconv"
	"strings"
)

type HexColor struct {
	Color color.Color

	Raw string
}

func (c *HexColor) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode {
		return fmt.Errorf("expected scalar node, got %q", n.Kind)
	}

	return c.UnmarshalString(n.Value)
}

func (c *HexColor) UnmarshalString(value string) error {
	rgba, err := hexToRGBA(value)
	if err != nil {
		return err
	}

	c.Raw = value
	c.Color = *rgba

	return nil
}

func (c *HexColor) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	return c.UnmarshalString(s)
}

func hexToRGBA(hex string) (*color.RGBA, error) {
	const (
		minHexLength = 3
		maxHexLength = 6
	)

	hex = strings.TrimPrefix(hex, "#")

	if len(hex) < minHexLength {
		return nil, errors.New("short hex string")
	}
	if len(hex) > maxHexLength {
		return nil, errors.New("long hex string")
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

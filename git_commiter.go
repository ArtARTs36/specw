package specw

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type GitCommitter struct {
	Name  string
	Email string
}

func (c *GitCommitter) String() string {
	return fmt.Sprintf("%s <%s>", c.Name, c.Email)
}

func (c *GitCommitter) UnmarshalYAML(n *yaml.Node) error {
	switch n.Kind { //nolint:exhaustive // other yaml types not supported
	case yaml.MappingNode:
		var objectable struct {
			Name  Env[string] `yaml:"name"`
			Email Env[string] `yaml:"email"`
		}

		if err := n.Decode(&objectable); err != nil {
			return err
		}

		c.Name = objectable.Name.Value
		c.Email = objectable.Email.Value

		return nil
	case yaml.ScalarNode:
		return c.UnmarshalString(n.Value)
	}

	return errors.New("unexpected yaml node, expects mapping or scalar")
}

func (c *GitCommitter) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(data, []byte{'"'}) && len(data) > 2 {
		return c.UnmarshalString(string(data[1 : len(data)-1]))
	}

	var objectable struct {
		Name  Env[string] `json:"name"`
		Email Env[string] `json:"email"`
	}

	if err := json.Unmarshal(data, &objectable); err != nil {
		return err
	}

	c.Name = objectable.Name.Value
	c.Email = objectable.Email.Value

	return nil
}

func (c *GitCommitter) UnmarshalString(val string) error {
	name := strings.Builder{}
	email := strings.Builder{}

	curr := &name

	for _, char := range val {
		switch char {
		case ' ':
			continue
		case '<':
			if name.Len() == 0 {
				return errors.New("found '<' before name found")
			}

			curr = &email
		case '>':
			if email.Len() == 0 {
				return errors.New("found '>' before email ends")
			}
		default:
			curr.WriteRune(char)
		}
	}

	if name.Len() == 0 {
		return errors.New("name not found")
	}
	if email.Len() == 0 {
		return errors.New("email not found")
	}

	c.Name = name.String()
	c.Email = email.String()

	return nil
}

func (c *GitCommitter) Valid() bool {
	return c != nil && c.Name != "" && c.Email != ""
}

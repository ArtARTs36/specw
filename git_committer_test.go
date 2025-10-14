package specw

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitCommitterUnmarshalYAML(t *testing.T) {
	cases := []struct {
		Title    string
		Content  string
		Expected *GitCommitter
	}{
		{
			Title:   "parse from scalar node",
			Content: `{committer: name <user@mail.ru>}`,
			Expected: &GitCommitter{
				Name:  "name",
				Email: "user@mail.ru",
			},
		},
		{
			Title:   "parse from mapping node",
			Content: `{committer: {name: name, email: user@mail.ru}}`,
			Expected: &GitCommitter{
				Name:  "name",
				Email: "user@mail.ru",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			var actual GitCommitter
			err := yaml.Unmarshal([]byte(c.Content), &actual)
			require.NoError(t, err)
		})
	}
}

func TestGitCommitterUnmarshalJSON(t *testing.T) {
	cases := []struct {
		Title    string
		Content  string
		Expected *GitCommitter
	}{
		{
			Title:   "parse from scalar node",
			Content: `{"committer": "name <user@mail.ru>"}`,
			Expected: &GitCommitter{
				Name:  "name",
				Email: "user@mail.ru",
			},
		},
		{
			Title:   "parse from mapping node",
			Content: `{"committer": {"name": "name", "email": "user@mail.ru"}}`,
			Expected: &GitCommitter{
				Name:  "name",
				Email: "user@mail.ru",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Title, func(t *testing.T) {
			var actual GitCommitter
			err := json.Unmarshal([]byte(c.Content), &actual)
			require.NoError(t, err)
		})
	}
}

func TestGitCommitterUnmarshalString(t *testing.T) {
	cases := []struct {
		Identity string
		Expected *GitCommitter
		Err      error
	}{
		{
			Identity: "simple <simple@mail.ru>",
			Expected: &GitCommitter{
				Name:  "simple",
				Email: "simple@mail.ru",
			},
		},
		{
			Identity: "github-actions[bot] <github-actions[bot]@users.noreply.github.com>",
			Expected: &GitCommitter{
				Name:  "github-actions[bot]",
				Email: "github-actions[bot]@users.noreply.github.com",
			},
		},
		{
			Identity: " <github-actions[bot]@users.noreply.github.com>",
			Err:      errors.New("found '<' before name found"),
		},
		{
			Identity: "name <>",
			Err:      errors.New("found '>' before email ends"),
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Identity, func(t *testing.T) {
			committer := &GitCommitter{}

			err := committer.UnmarshalString(tCase.Identity)
			if tCase.Err != nil {
				require.Equal(t, tCase.Err, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tCase.Expected, committer)
			}
		})
	}
}

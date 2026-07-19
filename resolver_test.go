package wikilink

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type resolverFunc func(*Node) ([]byte, []string, error)

func (f resolverFunc) ResolveWikilink(n *Node) ([]byte, []string, error) {
	return f(n)
}

func TestDefaultResolver(t *testing.T) {
	t.Parallel()

	tests := []struct {
		target   string
		fragment string
		want     string
	}{
		{
			target: "foo",
			want:   "foo.html",
		},
		{
			target: "foo bar",
			want:   "foo bar.html",
		},
		{
			target: "foo/bar",
			want:   "foo/bar.html",
		},
		{
			target: "foo bar.pdf",
			want:   "foo bar.pdf",
		},
		{
			target: "foo/bar.png",
			want:   "foo/bar.png",
		},
		{
			target:   "foo",
			fragment: "bar",
			want:     "foo.html#bar",
		},
		{
			target:   "foo/bar",
			fragment: "baz",
			want:     "foo/bar.html#baz",
		},
		{
			fragment: "foo",
			want:     "#foo",
		},
	}

	for _, tt := range tests {
		tt := tt
		name := fmt.Sprintf("%v#%v", tt.target, tt.fragment)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _, err := DefaultResolver.ResolveWikilink(&Node{
				Target:   []byte(tt.target),
				Fragment: []byte(tt.fragment),
			})
			require.NoError(t, err, "resolve failed")
			assert.Equal(t, tt.want, string(got), "result mismatch")
		})
	}
}

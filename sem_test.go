package semver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []*testVersion{
		tv("1.2.3").Expect(New(1, 2, 3)),
		tv("1.2.3-alpha").Expect(New(1, 2, 3).WithPreRelease("alpha")),
		tv("1.2.3+build").Expect(New(1, 2, 3).WithMetadata("build")),
		tv("1.2.3-alpha+build").Expect(New(1, 2, 3).WithPreRelease("alpha").WithMetadata("build")),
		tv("v1.2.3").Expect(New(1, 2, 3)),
		tv("v1.2.3-alpha").Expect(New(1, 2, 3).WithPreRelease("alpha")),
		tv("v1.2.3+build").Expect(New(1, 2, 3).WithMetadata("build")),
		tv("v1.2.3-alpha+build").Expect(New(1, 2, 3).WithPreRelease("alpha").WithMetadata("build")),
		tv("-1.0.0").Err(),
		tv("1.0.0-").Err(),
		tv("1.0.0-+").Err(),
		tv("1.0.0-.").Err(),
		tv("1.0.0-..").Err(),
		tv("a.b.c").Err(),
		tv("1.2").Err(),
		tv("").Err(),
	}
	for _, c := range cases {
		v, err := Parse(c.Input)
		if c.Error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.Expected, v)
		}
	}
}

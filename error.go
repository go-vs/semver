package semver

import (
	"errors"
	"fmt"
)

var ErrEmptyStr = errors.New("empty string")
var ErrInvalidStr = errors.New("invalid semver")
var ErrParseMajor = func(err error) error {
	return fmt.Errorf("parse major: %v", err)
}
var ErrParseMinor = func(err error) error {
	return fmt.Errorf("parse minor: %v", err)
}
var ErrParsePatch = func(err error) error {
	return fmt.Errorf("parse patch: %v", err)
}

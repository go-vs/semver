package semver

import (
	"regexp"
	"strconv"
)

const semRegexStr = `^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

var semRegex = regexp.MustCompile(semRegexStr)

func Parse(v string) (*Version, error) {
	if v == "" {
		return nil, ErrEmptyStr
	}
	var err error
	m := semRegex.FindStringSubmatch(v)
	if m == nil {
		return nil, ErrInvalidStr
	}
	sem := &Version{
		PreRelease: m[4],
		Metadata:   m[5],
	}
	majorStr, minorStr, patchStr := m[1], m[2], m[3]
	sem.Major, err = strconv.ParseUint(majorStr, 10, 64)
	if err != nil {
		return nil, ErrParseMajor(err)
	}
	sem.Minor, err = strconv.ParseUint(minorStr, 10, 64)
	if err != nil {
		return nil, ErrParseMinor(err)
	}
	sem.Patch, err = strconv.ParseUint(patchStr, 10, 64)
	if err != nil {
		return nil, ErrParsePatch(err)
	}

	return sem, nil
}

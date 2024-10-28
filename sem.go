package semver

import (
	"fmt"
	"strings"
)

type Version struct {
	Major, Minor, Patch uint64
	PreRelease          string
	Metadata            string
}

func New(major, minor, patch uint64) *Version {
	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (s *Version) WithPreRelease(preRelease string) *Version {
	s.PreRelease = preRelease
	return s
}

func (s *Version) WithMetadata(metadata string) *Version {
	s.Metadata = metadata
	return s
}

func (s *Version) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprint(s.Major))
	sb.WriteByte('.')
	sb.WriteString(fmt.Sprint(s.Minor))
	sb.WriteByte('.')
	sb.WriteString(fmt.Sprint(s.Patch))
	if s.PreRelease != "" {
		sb.WriteByte('-')
		sb.WriteString(s.PreRelease)
	}
	if s.Metadata != "" {
		sb.WriteByte('+')
		sb.WriteString(s.Metadata)
	}
	return sb.String()
}

func (s *Version) Compare(other *Version) int {
	return Compare(s, other)
}

func Compare(a, b *Version) int {
	if a.Major != b.Major {
		return cmpUint64(a.Major, b.Major)
	}
	if a.Minor != b.Minor {
		return cmpUint64(a.Minor, b.Minor)
	}
	if a.Patch != b.Patch {
		return cmpUint64(a.Patch, b.Patch)
	}
	return strings.Compare(a.PreRelease, b.PreRelease)
}

func cmpUint64(a, b uint64) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

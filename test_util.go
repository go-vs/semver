package semver

type testVersion struct {
	Input    string
	Expected *Version
	Error    bool
}

func tv(v string) *testVersion {
	return &testVersion{
		Input: v,
	}
}

func (t *testVersion) In(input string) *testVersion {
	t.Input = input
	return t
}

func (t *testVersion) Expect(v *Version) *testVersion {
	t.Expected = v
	return t
}

func (t *testVersion) Err() *testVersion {
	t.Error = true
	return t
}

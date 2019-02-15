package kroki

import (
	"testing"
)

func TestCreatePayload(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			in:   `digraph G {Hello->World}`,
			want: "eNpKyUwvSizIUHBXqPZIzcnJ17ULzy_KSakFBAAA__9sQAjG",
		},
	}
	for _, c := range cases {
		result, err := CreatePayload(c.in)
		if err != nil {
			t.Errorf("CreatePayload error:\n%+v", err)
		}
		if result != c.want {
			t.Errorf("CreatePayload error\nexpected: %s\nactual:   %s", c.want, result)
		}
	}
}

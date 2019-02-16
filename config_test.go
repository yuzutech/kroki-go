package kroki

import (
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

func TestConfiguration(t *testing.T) {
	cases := []struct {
		in   string
		want Configuration
	}{
		{
			in: "{}",
			want: Configuration{
				URL:     "https://demo.kroki.io",
				Timeout: time.Second * 20,
			},
		},
		{
			in: `
url: "https://kroki.io"
timeout: 30
`,
			want: Configuration{
				URL:     "https://kroki.io",
				Timeout: time.Second * 30,
			},
		},
	}

	for _, tc := range cases {
		var got Configuration
		err := yaml.Unmarshal([]byte(tc.in), &got)
		if err != nil {
			t.Errorf("Unmarshal(%q) error:\n%+v", tc.in, err)
		}
		if got != tc.want {
			t.Errorf("error\ngot:\n%s\nwant:\n%s", got, tc.want)
		}
	}

}

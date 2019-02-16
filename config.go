package kroki

import (
	"time"

	"github.com/pkg/errors"
)

// ImageFormat the image format returned by Kroki
type ImageFormat string

// GraphFormat the format of the graph passed to Kroki
type GraphFormat string

const (
	// Svg is the svg format
	Svg ImageFormat = "svg"
)

const (
	// Graphviz is the graphviz graph format
	Graphviz GraphFormat = "graphviz"
	// BlockDiag is the blockdiag graph format
	BlockDiag GraphFormat = "blockdiag"
	// SeqDiag is the seqdiag graph format
	SeqDiag GraphFormat = "seqdiag"
	// Mermaid is the mermaid graph format
	Mermaid GraphFormat = "mermaid"
	// Nomnoml is the nomnoml graph format
	Nomnoml GraphFormat = "nomnoml"
	// Plantuml is the plantuml graph format
	Plantuml GraphFormat = "plantuml"
	// Svgbob is the svgbob graph format
	Svgbob GraphFormat = "svgbob"
	// C4plantuml is the c4plantuml graph format
	C4plantuml GraphFormat = "c4plantuml"
	// Umlet is the umlet graph format
	Umlet GraphFormat = "umlet"
)

// Configuration contains the configuration for the Kroki client
type Configuration struct {
	URL     string
	Timeout time.Duration
}

// UnmarshalYAML parses a reporter configuration from YAML.
func (configuration *Configuration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawConfiguration struct {
		URL     string
		Timeout uint64
	}
	// this is the default configuration
	rawConfig := rawConfiguration{
		URL:     "https://kroki.io",
		Timeout: 20,
	}

	if err := unmarshal(&rawConfig); err != nil {
		return errors.Wrap(err, "fail to decode the yaml configuration")
	}
	*configuration = Configuration{
		URL:     rawConfig.URL,
		Timeout: time.Second * time.Duration(rawConfig.Timeout),
	}
	return nil
}

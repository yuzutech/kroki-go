package kroki

import (
	"time"

	"github.com/pkg/errors"
)

// ImageFormat the image format returned by Kroki
type ImageFormat string

// DiagramType the type of diagram sent to Kroki
type DiagramType string

const (
	// Svg is the svg format
	Svg ImageFormat = "svg"
)

const (
	// GraphViz is the graphviz diagram type
	GraphViz DiagramType = "graphviz"
	// BlockDiag is the blockdiag diagram type
	BlockDiag DiagramType = "blockdiag"
	// SeqDiag is the seqdiag diagram type
	SeqDiag DiagramType = "seqdiag"
	// Mermaid is the mermaid diagram type
	Mermaid DiagramType = "mermaid"
	// Nomnoml is the nomnoml diagram type
	Nomnoml DiagramType = "nomnoml"
	// PlantUML is the plantuml diagram type
	PlantUML DiagramType = "plantuml"
	// Svgbob is the svgbob diagram type
	Svgbob DiagramType = "svgbob"
	// C4PlantUML is the c4plantuml diagram type
	C4PlantUML DiagramType = "c4plantuml"
	// UMlet is the umlet diagram type
	UMlet DiagramType = "umlet"
	// Erd is the erd diagram type
	Erd DiagramType = "erd"
	// NwDiag is the nwdiag diagram type
	NwDiag DiagramType = "nwdiag"
	// ActDiag is the actdiag diagram type
	ActDiag DiagramType = "actdiag"
	// Ditaa is the ditaa diagram type
	Ditaa DiagramType = "ditaa"
)

// Configuration contains the configuration for the Kroki client
type Configuration struct {
	URL     string
	Timeout time.Duration
}

// UnmarshalYAML parses a kroki configuration from YAML
func (configuration *Configuration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawConfiguration struct {
		URL     string
		Timeout uint64
	}
	// default configuration
	rawConfig := rawConfiguration{
		URL:     "https://demo.kroki.io",
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

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
	// SVG is the svg format
	SVG ImageFormat = "svg"
	// PNG is the png format
	PNG ImageFormat = "png"
	// JPEG is the jpeg format
	JPEG ImageFormat = "jpeg"
	// PDF is the pdf format
	PDF ImageFormat = "pdf"
	// Base64 is the base64 format
	Base64 ImageFormat = "base64"
)

const (
	// ActDiag is the actdiag diagram type
	ActDiag DiagramType = "actdiag"
	// BlockDiag is the blockdiag diagram type
	BlockDiag DiagramType = "blockdiag"
	// BPMN is the bpmn diagram type
	BPMN DiagramType = "bpmn"
	// BPMN is the bpmn diagram type
	Bytefield DiagramType = "bytefield"
	// C4PlantUML is the c4plantuml diagram type
	C4PlantUML DiagramType = "c4plantuml"
	// Ditaa is the ditaa diagram type
	Ditaa DiagramType = "ditaa"
	// Erd is the erd diagram type
	Erd DiagramType = "erd"
	// Excalidraw is the excalidraw diagram type
	Excalidraw DiagramType = "excalidraw"
	// GraphViz is the graphviz diagram type
	GraphViz DiagramType = "graphviz"
	// Mermaid is the mermaid diagram type
	Mermaid DiagramType = "mermaid"
	// Nomnoml is the nomnoml diagram type
	Nomnoml DiagramType = "nomnoml"
	// NwDiag is the nwdiag diagram type
	NwDiag DiagramType = "nwdiag"
	// PacketDiag is the packetdiag diagram type
	PacketDiag DiagramType = "packetdiag"
	// PlantUML is the plantuml diagram type
	PlantUML DiagramType = "plantuml"
	// RackDiag is the rackdiag diagram type
	RackDiag DiagramType = "rackdiag"
	// SeqDiag is the seqdiag diagram type
	SeqDiag DiagramType = "seqdiag"
	// Svgbob is the svgbob diagram type
	Svgbob DiagramType = "svgbob"
	// UMlet is the umlet diagram type
	UMlet DiagramType = "umlet"
	// Vega is the vega diagram type
	Vega DiagramType = "vega"
	// VegaLite is the vegalite diagram type
	VegaLite DiagramType = "vegalite"
	// WaveDrom is the wavedrom diagram type
	WaveDrom DiagramType = "wavedrom"
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

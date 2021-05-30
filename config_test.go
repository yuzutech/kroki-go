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
				URL:     "https://kroki.io",
				Timeout: time.Second * 20,
			},
		},
		{
			in: `
url: "https://a.kroki.io"
timeout: 30
`,
			want: Configuration{
				URL:     "https://a.kroki.io",
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

func TestGetSupportedImageFormats(t *testing.T) {
	supportedImageFormats := GetSupportedImageFormats()
	checkContainsImageFormat(t, supportedImageFormats, Base64)
	checkContainsImageFormat(t, supportedImageFormats, SVG)
	checkContainsImageFormat(t, supportedImageFormats, JPEG)
	checkContainsImageFormat(t, supportedImageFormats, PDF)
	checkContainsImageFormat(t, supportedImageFormats, PNG)
}

func checkContainsImageFormat(t *testing.T, list []ImageFormat, imageFormat ImageFormat) {
	if !containsImageFormat(list, imageFormat) {
		t.Errorf("error\n%s should be a supported image format", imageFormat)
	}
}

func containsImageFormat(s []ImageFormat, imageFormat ImageFormat) bool {
	for _, v := range s {
		if v == imageFormat {
			return true
		}
	}
	return false
}

func TestGetSupportedDiagramTypes(t *testing.T) {
	supportedDiagramTypes := GetSupportedDiagramTypes()
	checkContainsDiagramType(t, supportedDiagramTypes, ActDiag)
	checkContainsDiagramType(t, supportedDiagramTypes, BlockDiag)
	checkContainsDiagramType(t, supportedDiagramTypes, BPMN)
	checkContainsDiagramType(t, supportedDiagramTypes, Bytefield)
	checkContainsDiagramType(t, supportedDiagramTypes, C4PlantUML)
	checkContainsDiagramType(t, supportedDiagramTypes, Ditaa)
	checkContainsDiagramType(t, supportedDiagramTypes, Erd)
	checkContainsDiagramType(t, supportedDiagramTypes, Excalidraw)
	checkContainsDiagramType(t, supportedDiagramTypes, Mermaid)
	checkContainsDiagramType(t, supportedDiagramTypes, Nomnoml)
	checkContainsDiagramType(t, supportedDiagramTypes, NwDiag)
	checkContainsDiagramType(t, supportedDiagramTypes, PacketDiag)
	checkContainsDiagramType(t, supportedDiagramTypes, Pikchr)
	checkContainsDiagramType(t, supportedDiagramTypes, PlantUML)
	checkContainsDiagramType(t, supportedDiagramTypes, RackDiag)
	checkContainsDiagramType(t, supportedDiagramTypes, SeqDiag)
	checkContainsDiagramType(t, supportedDiagramTypes, Svgbob)
	checkContainsDiagramType(t, supportedDiagramTypes, UMlet)
	checkContainsDiagramType(t, supportedDiagramTypes, Vega)
	checkContainsDiagramType(t, supportedDiagramTypes, VegaLite)
	checkContainsDiagramType(t, supportedDiagramTypes, WaveDrom)
}

func checkContainsDiagramType(t *testing.T, list []DiagramType, diagramType DiagramType) {
	if !containsDiagramType(list, diagramType) {
		t.Errorf("error\n%s should be a supported diagram type", diagramType)
	}
}

func containsDiagramType(s []DiagramType, diagramType DiagramType) bool {
	for _, v := range s {
		if v == diagramType {
			return true
		}
	}
	return false
}

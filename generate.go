package kroki

import (
	"fmt"
	"os"
)

const MAX_URI_LENGTH = 4096

// FromString takes a string and returns the image generated by Kroki
func (c *Client) FromString(input string, diagramType DiagramType, imageFormat ImageFormat) (string, error) {
	payload, err := CreatePayload(input)
	if err != nil {
		return "", err
	}
	if len(payload) > MAX_URI_LENGTH {
		return c.PostRequest(input, diagramType, imageFormat)
	} else {
		return c.GetRequest(payload, diagramType, imageFormat)
	}
}

// FromFile takes a file path and returns the image generated by Kroki
func (c *Client) FromFile(path string, diagramType DiagramType, imageFormat ImageFormat) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("fail to read file '%s': %w", path, err)
	}
	input := string(content)
	payload, err := CreatePayload(input)
	if err != nil {
		return "", err
	}
	if len(payload) > MAX_URI_LENGTH {
		return c.PostRequest(input, diagramType, imageFormat)
	} else {
		return c.GetRequest(payload, diagramType, imageFormat)
	}
}

// WriteToFile takes a file path and a string
// write the string into the file
func (c *Client) WriteToFile(path string, result string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("fail to create file '%s': %w", path, err)
	}
	defer file.Close()
	_, err = file.Write([]byte(result))
	if err != nil {
		return fmt.Errorf("fail to write to file '%s': %w", path, err)
	}
	err = file.Sync()
	if err != nil {
		return fmt.Errorf("fail to sync file '%s': %w", path, err)
	}
	return nil
}

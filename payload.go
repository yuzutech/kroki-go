package kroki

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
)

// CreatePayload takes a string and returns a payload in deflate + base64 format
func CreatePayload(input string) (string, error) {

	var buffer bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buffer, 9)
	if err != nil {
		return "", fmt.Errorf("fail to create the writer: %w", err)
	}
	_, err = writer.Write([]byte(input))
	writer.Close()
	if err != nil {
		return "", fmt.Errorf("fail to create the payload: %w", err)
	}
	result := base64.URLEncoding.EncodeToString(buffer.Bytes())
	return result, nil
}

package kroki

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFromString(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := "eNpKyUwvSizIUHBXqPZIzcnJ17ULzy_KSakFBAAA__9sQAjG"
		uri := strings.Split(r.RequestURI, "/")
		payload := uri[len(uri)-1]
		if payload != expected {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", expected, payload)
		}
		imageFormat := uri[len(uri)-2]
		if imageFormat != string(SVG) {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", string(SVG), imageFormat)
		}
		diagramType := uri[len(uri)-3]
		if diagramType != string(GraphViz) {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", string(GraphViz), diagramType)
		}
	}))
	defer ts.Close()
	port, err := strconv.ParseUint(strings.Split(ts.URL, ":")[2], 10, 16)
	if err != nil {
		t.Errorf("error getting the port :\n%+v", err)
	}
	client := New(Configuration{
		URL:     fmt.Sprintf("http://localhost:%d", port),
		Timeout: time.Second * 10,
	})
	_, err = client.FromString("digraph G {Hello->World}", GraphViz, SVG)
	if err != nil {
		t.Errorf("FromString error:\n%+v", err)
	}
}

func TestFromStringWithServerError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := "eNpczEEKwkAMheF9TxGz7wmKXkRcxPoYho6JZEQZxLtLVHDs8vE-_mOxeTllSfQYiH5r3FGCwuWKGoPf1xify7nytNb5w27wRpDaNjwNf2Y_WzGnLXFyQBtKsTsfIrSOd_aSdfmiLt0Bc9GEIM9XAAAA__9Z7UER"
		uri := strings.Split(r.RequestURI, "/")
		payload := uri[len(uri)-1]
		if payload != expected {
			t.Errorf("FromStringWithServerError error\nexpected: %s\nactual:   %s", expected, payload)
		}
		imageFormat := uri[len(uri)-2]
		if imageFormat != string(JPEG) {
			t.Errorf("FromStringWithServerError error\nexpected: %s\nactual:   %s", string(JPEG), imageFormat)
		}
		diagramType := uri[len(uri)-3]
		if diagramType != string(BlockDiag) {
			t.Errorf("FromStringWithServerError error\nexpected: %s\nactual:   %s", string(BlockDiag), diagramType)
		}
		w.WriteHeader(400)
		_, err := w.Write([]byte("Error 400: Unsupported output format: jpeg. Must be one of blockdiag for png, svg or pdf"))
		if err != nil {
			t.Errorf("error writting response:\n%+v", err)
		}
	}))
	defer ts.Close()
	port, err := strconv.ParseUint(strings.Split(ts.URL, ":")[2], 10, 16)
	if err != nil {
		t.Errorf("error getting the port :\n%+v", err)
	}
	client := New(Configuration{
		URL:     fmt.Sprintf("http://localhost:%d", port),
		Timeout: time.Second * 10,
	})
	_, err = client.FromString(`blockdiag {
  blockdiag -> generates -> "block-diagrams";
  blockdiag -> is -> "very easy!";

  blockdiag [color = "greenyellow"];
  "block-diagrams" [color = "pink"];
  "very easy!" [color = "orange"];
}`, BlockDiag, JPEG)
	if err == nil {
		t.Error("FromStringWithServerError must return an error")
	}
	expectedErrorMessage := "fail to generate the image {status: 400, body: Error 400: Unsupported output format: jpeg. Must be one of blockdiag for png, svg or pdf}"
	errorMessage := err.Error()
	if errorMessage != expectedErrorMessage {
		t.Errorf("FromStringWithServerError error\nexpected: %s\nactual:   %s", expectedErrorMessage, errorMessage)
	}

}

func TestFromFile(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := "eNpKyUwvSizIUHBXqPZIzcnJ17ULzy_KSakFBAAA__9sQAjG"
		uri := strings.Split(r.RequestURI, "/")
		payload := uri[len(uri)-1]
		if payload != expected {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", expected, payload)
		}
		imageFormat := uri[len(uri)-2]
		if imageFormat != string(SVG) {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", string(SVG), imageFormat)
		}
		diagramType := uri[len(uri)-3]
		if diagramType != string(GraphViz) {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", string(GraphViz), diagramType)
		}
	}))
	defer ts.Close()
	port, err := strconv.ParseUint(strings.Split(ts.URL, ":")[2], 10, 16)
	if err != nil {
		t.Errorf("error getting the port :\n%+v", err)
	}
	client := New(Configuration{
		URL:     fmt.Sprintf("http://localhost:%d", port),
		Timeout: time.Second * 10,
	})
	_, err = client.FromFile("tests/hello.dot", GraphViz, SVG)
	if err != nil {
		t.Errorf("FromString error:\n%+v", err)
	}
}

func TestWriteToFile(t *testing.T) {
	client := New(Configuration{})
	expected := "clojure"
	filePath := "tests/test_write.ignore.test"
	err := client.WriteToFile(filePath, expected)
	if err != nil {
		t.Errorf("WriteToFile error:\n%+v", err)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Errorf("read file error:\n%+v", err)
	}
	if string(content) != expected {
		t.Errorf("WriteToFile error\nexpected: %s\nactual:   %s", expected, string(content))
	}
}

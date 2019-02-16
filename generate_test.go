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
		if imageFormat != string(Svg) {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", string(Svg), imageFormat)
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
	_, err = client.FromString("digraph G {Hello->World}", GraphViz, Svg)
	if err != nil {
		t.Errorf("FromString error:\n%+v", err)
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
		if imageFormat != string(Svg) {
			t.Errorf("FromString error\nexpected: %s\nactual:   %s", string(Svg), imageFormat)
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
	_, err = client.FromFile("tests/hello.dot", GraphViz, Svg)
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

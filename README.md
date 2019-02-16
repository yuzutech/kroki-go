# Kroki Go

[![Build Status](https://travis-ci.org/yuzutech/kroki-go.svg?branch=master)](https://travis-ci.org/yuzutech/kroki-go)

A Golang library for [https://kroki.io/](https://kroki.io/).

## Usage

Create a client:

```golang
client := kroki.New(kroki.Configuration{
		URL:     "https://demo.kroki.io",
		Timeout: time.Second * 20,
	})
```

### String to diagram

Use the `FromString` function to convert a string to a diagram:

```golang
result, err := client.FromString("digraph G {Hello->World}", kroki.Graphviz, kroki.Svg)
```

Here, `result` contains the image returned by Kroki as a string.

### File to diagram

Use the `FromFile` function to convert a file to a diagram:

```bash
echo "digraph G {Hello->World}" > hello.dot
```

```golang
result, err := client.FromFile("./hello.dot", kroki.Graphviz, kroki.Svg)
```

### Write to a file

Use the `WriteToFile` function to write a result to a file:

```golang
err = client.WriteToFile("./result.svg", result)
```

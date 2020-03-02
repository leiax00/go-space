package test

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestXmlConstruct(t *testing.T) {
	obj := &Item{
		Attr1:  "a1",
		Attr2:  "",
		Attr3:  "",
		Attr4:  []string{"a41", "a42"},
		Child1: "c1",
		Child2: "",
		Child: &Child{
			Name: "leiax00",
			Age:  0,
		},
	}
	if bytes, err := xml.MarshalIndent(obj, "", "    "); err == nil {
		os.Stdout.Write(bytes)
	}
	os.Stdout.Write([]byte("\n"))
}

type Item struct {
	Attr1  string   `xml:"attr1,attr,omitempty"`
	Attr2  string   `xml:"attr2,attr,omitempty"`
	Attr3  string   `xml:"attr3,attr"`
	Attr4  []string `xml:"attr4,attr,omitempty"`
	Child1 string   `xml:"child_1"`
	Child2 string   `xml:"child_2"`
	Child  *Child   `xml:"child"`
}

type Child struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

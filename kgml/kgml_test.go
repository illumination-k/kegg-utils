package kgml_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"kegg/kgml"
	"os"
	"testing"
)

func TestKGMLParse(t *testing.T) {
	r, err := os.Open("./tests/ko00010.kgml")

	if err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadAll(r)

	if err != nil {
		t.Error(err)
	}

	var pathway kgml.KGMLPathway

	if err := xml.Unmarshal(content, &pathway); err != nil {
		t.Error(err)
	}

	for _, e := range pathway.Entries {
		fmt.Printf("%#v\n", e)
	}

	for _, e := range pathway.Reactions {
		fmt.Printf("%#v\n", e)
	}

	for _, e := range pathway.Relations {
		fmt.Printf("%#v\n", e)
	}

	t.Error()
}

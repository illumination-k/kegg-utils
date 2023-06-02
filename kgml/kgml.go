package kgml

import (
	"encoding/xml"
	"strings"
)

type KGMLEntry struct {
	ID       int
	Names    []string
	Type     string
	Reaction string
	Link     string
}

func (e *KGMLEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	ee := struct {
		ID       int    `xml:"id,attr"`
		Names    string `xml:"name,attr"`
		Type     string `xml:"type,attr"`
		Reaction string `xml:"reaction,attr"`
		Link     string `xml:"link,attr"`
	}{}

	if err := d.DecodeElement(&ee, &start); err != nil {
		return err
	}

	*e = KGMLEntry{ID: ee.ID, Names: strings.Split(ee.Names, " "), Type: ee.Type, Reaction: ee.Reaction, Link: ee.Link}

	return nil
}

type SubType struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Relation struct {
	Entry1  int     `xml:"entry1,attr"`
	Entry2  int     `xml:"entry2,attr"`
	Type    string  `xml:"type,attr"`
	SubType SubType `xml:"subtype"`
}

type SubsTrate struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Product struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type KGMLReaction struct {
	ID        int          `xml:"id,attr"`
	Name      string       `xml:"name,attr"`
	Type      string       `xml:"type,attr"`
	Subtrates []*SubsTrate `xml:"substrate"`
	Products  []*Product   `xml:"product"`
}

type KGMLPathway struct {
	XMLName   xml.Name        `xml:"pathway"`
	Name      string          `xml:"name,attr"`
	Title     string          `xml:"title,attr"`
	Entries   []*KGMLEntry    `xml:"entry"`
	Reactions []*KGMLReaction `xml:"reaction"`
	Relations []*Relation     `xml:"relation"`
}

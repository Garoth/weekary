package main

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func NewElement(tagname, innerHTML string) *goquery.Selection {
	node := &html.Node{}

	node.Type = html.ElementNode
	// I'm unsure what this field is for, and goquery seems happy w/o it
	// node.DataAtom = atom.Div
	node.Data = tagname

	return goquery.NewDocumentFromNode(node).First().AppendHtml(innerHTML)
}

func SetAttr(selection *goquery.Selection, attr string, value string) {
	for i := 0; i < selection.Size(); i++ {
		node := selection.Get(i)
		attrs := make([]html.Attribute, 0)

		for _, a := range node.Attr {
			if a.Key != attr {
				newAttr := new(html.Attribute)
				newAttr.Key = a.Key
				newAttr.Val = a.Val
				attrs = append(attrs, *newAttr)
			}
		}

		newAttr := new(html.Attribute)
		newAttr.Key = attr
		newAttr.Val = value
		attrs = append(attrs, *newAttr)
		node.Attr = attrs
	}
}

package main

import (
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html"
)

func NewDiv() *html.Node {
	node := &html.Node{}

	node.Type = html.ElementNode
	node.DataAtom = atom.Div
	node.Data = "foo"

	return node
}

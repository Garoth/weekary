package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/codeskyblue/go-sh"
	"github.com/yosssi/gohtml"
)

// Generates a section for a given person, with their name an weekly status
func makeNameHeader(name string) *goquery.Selection {
	selection := &goquery.Selection{}

	nameDiv := NewElement("div", name)
	SetAttr(nameDiv, "style", "font-size: 18px;")
	selection = selection.AddSelection(nameDiv)

	statusDiv := NewElement("div", "Working normal hours.")
	SetAttr(statusDiv, "style", "font-size: 14px; color: rgb(121, 121, 121);")
	selection = selection.AddSelection(statusDiv)

	return selection
}

func main() {
	file, err := os.Open("/Users/athorp/Configs/weekary/template.html")
	if err != nil {
		log.Fatalln("Couldn't read template file", err)
	}

	referenceDoc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalln("Couldn't parse template document", err)
	}

	newDoc := goquery.CloneDocument(referenceDoc)
	body := newDoc.Find("body").First()
	body.Empty()

	body.AppendSelection(makeNameHeader("Andrei Thorp"))

	finalHtml, err := newDoc.Html()
	if err != nil {
		log.Fatalln("Couldn't produce final HTML!", err)
	}

	// Creating temporary file, writing HTML to it, opening browser
	temporaryFile, err := ioutil.TempFile("", "weekary-")
	defer temporaryFile.Close()
	if err != nil {
		log.Fatalln("Couldn't create temporary file to write to!", err)
	}

	if _, err := temporaryFile.WriteString(gohtml.Format(finalHtml)); err != nil {
		log.Fatalln("Couldn't write HTML to temporary file!", err)
	}

	fmt.Println("Successfully generated. Filename is:")
	fmt.Println(temporaryFile.Name())
	fmt.Println("Opening in browser...")

	command := sh.Command("open", temporaryFile.Name())
	command.Start()
	if err := command.Wait(); err != nil {
		log.Fatalln("Couldn't open browser to", temporaryFile.Name(), "--", err)
	}
}

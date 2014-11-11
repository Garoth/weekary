package main

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

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
	newDoc.Find("body").First().Empty()

	log.Println("Text is", referenceDoc.Find("#testdiv").First().Text())
	tmpHTml, err := newDoc.Find(".nothingtofind").AddNodes(NewDiv()).First().Html()
	log.Println("div has", tmpHTml)

	docStr, err := newDoc.Html()
	if err != nil {
		log.Fatalln("Couldn't produce final HTML!", err)
	}
	log.Println("Full dom is", docStr)
}

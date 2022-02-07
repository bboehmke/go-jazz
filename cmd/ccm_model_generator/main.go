package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/go-wordwrap"
)

//go:embed type.go.tpl
var tplStr string

// map of known models based on the type ID
var modelTypeRef = make(map[string]*Model)

// regex to parse field definition
var fieldRegEx = regexp.MustCompile(`(.+?) \(type: (.+?)\)\.(.+)`)

// parse fields of the given selection
func parseFields(selection *goquery.Selection, typeId string) []Field {
	var fields []Field

	selection.Find("li").Each(func(i int, selection *goquery.Selection) {
		match := fieldRegEx.FindStringSubmatch(strings.TrimSpace(selection.Text()))

		if len(match) == 4 {
			if badFields, ok := skipFields[typeId]; ok {
				if _, ok = badFields[strings.TrimSpace(match[1])]; ok {
					return // skip this one
				}
			}

			fields = append(fields, Field{
				Name:        strings.TrimSpace(match[1]),
				Type:        strings.TrimSpace(match[2]),
				Description: strings.Split(wordwrap.WrapString(strings.TrimSpace(match[3]), 75), "\n"),
			})
		}
	})

	return fields
}

// parseHeadLine of RTC CCM object definition
func parseHeadLine(headline string) (string, string) {
	// has type & element id
	if strings.Contains(headline, "(type:") {
		split := strings.Split(strings.ReplaceAll(headline, ")", ""), "(type:")
		return strings.TrimSpace(split[0]), strings.TrimSpace(split[1])
	}
	// has only type
	if strings.Contains(headline, "com.") {
		return "", strings.TrimSpace(headline)
	}
	// has only element id
	return strings.TrimSpace(headline), ""
}

func main() {
	// get latest docu from wiki
	response, err := http.Get("https://bugmenot:bugmenot@jazz.net/wiki/bin/view/Main/ReportsRESTAPI")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// list of resources to handle
	var resources = map[string]struct{}{
		"foundation": {},
		"scm":        {},
		"build":      {},
		"workitem":   {},
	}

	var models []Model
	var resource string
	var typeHeadline string
	var description string
	var linkRef string

	// find content
	doc.Find("div.patternTopic").Children().Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		node := s.Get(0)

		if node.Data == "h3" { // resource headline
			if _, ok := resources[text]; ok {
				resource = text
				typeHeadline = ""
				description = ""
				linkRef = ""
			} else {
				resource = ""
				typeHeadline = ""
				description = ""
				linkRef = ""
			}
		} else if resource != "" { // in resource section
			if node.Data == "h4" { // object headline
				typeHeadline = text
				description = ""
				linkRef, _ = s.Find("a").Attr("name")
			} else if typeHeadline != "" {

				// search for property definition
				if node.Data == "ul" {
					// get element ID and type ID from headline
					elementID, typeId := parseHeadLine(strings.TrimSpace(typeHeadline))

					// check if type is in missing list
					if typeId == "" {
						if t, ok := missingTypes[elementID]; ok {
							typeId = t
						}
					}

					// fix element IDs
					if e, ok := invalidElementIDs[typeId]; ok {
						elementID = e
					}

					// create model and parse fields
					model := Model{
						LinkRef:     linkRef,
						Description: strings.Split(wordwrap.WrapString(description, 75), "\n"),

						ResourceID: resource,
						ElementID:  elementID,
						TypeID:     typeId,
						Fields:     parseFields(s, typeId),
					}

					if model.TypeID != "" {
						modelTypeRef[model.TypeID] = &model
					}

					models = append(models, model)

				} else if description == "" { // first paragraph = description
					description = text
				}
			}
		}
	})

	// use template to generate model definition
	tpl, err := template.New("").Parse(tplStr)
	if err != nil {
		panic(err)
	}

	goFile, err := os.Create("ccm_model_gen.go")
	if err != nil {
		panic(err)
	}
	defer goFile.Close()

	// write header
	_, err = goFile.WriteString("package jazz\n\n")
	if err != nil {
		panic(err)
	}
	_, err = goFile.WriteString("// Code generated! DO NOT EDIT\n\n")
	if err != nil {
		panic(err)
	}
	_, err = goFile.WriteString("import \"time\"\n")
	if err != nil {
		panic(err)
	}

	// write models
	for _, model := range models {
		err = tpl.Execute(goFile, model)
		if err != nil {
			panic(err)
		}
	}
}

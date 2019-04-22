/*
  | Response genreator is intended to provide the functionality
  | that are necessary to render a response for
  | incoming requests ...
  | HTML response havily realies on Jet Templating engine ...
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet"
	"github.com/gorilla/csrf"
)

var root, _ = os.Getwd()

// Jet is sington of the jetTemplate Engine and used by
// this file
var Jet = jet.NewHTMLSet(filepath.Join(root, "view"))

func init() {
	Jet.SetDevelopmentMode(true)
}

// JSON Returns the data in form of "JSON" for the incoming
// request
func JSON(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to generate json ")
	}
	fmt.Fprint(w, string(response))
}

// View Returns a jet view in response of the in coming request
// with the data supplied as parameter
func View(w http.ResponseWriter, r *http.Request, data interface{}, viewName string) {

	templateName := viewName
	t, err := Jet.GetTemplate(templateName)
	if err != nil {
		log.Println("Failed to get template ")
	}
	dataMap := make(map[string]interface{})
	if data != nil {
		dataMap = data.(map[string]interface{})
	}
	vars := make(jet.VarMap)
	dataMap["config"] = Config
	dataMap["token"] = csrf.Token(r)
	dataMap["currentURL"] = r.URL.Path
	if err = t.Execute(w, vars, dataMap); err != nil {
		log.Println("Failed to execute view tepmlate ")
	}
}

// HTMLString Returns a jet view in response of the in coming request
// with the data supplied as parameter
// Used at email templates
func HTMLString(data interface{}, viewName string) string {
	var html bytes.Buffer
	templateName := viewName
	t, err := Jet.GetTemplate(templateName)
	if err != nil {
		log.Println("Failed to  get tempalte ")
	}
	vars := make(jet.VarMap)
	if err = t.Execute(&html, vars, data); err != nil {
		log.Println("Failed to execute view tepmlate ")
	}
	return html.String()
}

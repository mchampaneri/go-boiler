package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet"
	"github.com/fatih/color"
	"github.com/gorilla/csrf"
)

/*
  | Render Package is intended to provide the functionality
  | that are necessary to render a response for any
  | in coming request
*/

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
	DefaultLogger.Info("Json Rendered")
	if err != nil {
		DefaultLogger.Error("Error " + err.Error() + " occured during rendering Json Response")
		color.Red(" - respnose-generator.go  Json : %s", err.Error())
	}
	fmt.Fprint(w, string(response))
}

// View Returns a jet view in response of the in coming request
// with the data supplied as parameter
func View(w http.ResponseWriter, r *http.Request, data interface{}, viewName string) {
	session, err := UserSession.Get(r, "mvc-user-session")
	if err != nil || session.IsNew {
		// Just Ignore
	}
	templateName := viewName
	t, err := Jet.GetTemplate(templateName)
	if err != nil {
		DefaultLogger.Error("Error " + err.Error() + " occured during rendering View Response")
		color.Red(" - respnose-generator.go  View : %s", err.Error())
	}
	dataMap := make(map[string]interface{})
	if data != nil {
		dataMap = data.(map[string]interface{})
	}
	vars := make(jet.VarMap)
	dataMap["AppUrl"] = Config.AppURL
	// vars.Set("Auth", "true")
	if session.Values["auth"] == true {
		dataMap["Auth"] = true
		dataMap["Name"] = session.Values["name"]
		dataMap["NickName"] = session.Values["nickname"]
		dataMap["Email"] = session.Values["email"]
		dataMap["ProfilePic"] = session.Values["profile_pic"]
		dataMap["CoverPic"] = session.Values["cover_pic"]
		if session.Values["active"] == true {
			dataMap["Active"] = session.Values["active"]
		}
	}

	dataMap["Message"] = session.Values["message"]
	dataMap["Token"] = csrf.Token(r)
	dataMap["Url"] = r.URL.Path

	// Resetting the Session Message
	session.Options.MaxAge = 0
	session.Values["message"] = nil
	session.Save(r, w)
	if err = t.Execute(w, vars, dataMap); err != nil {
		DefaultLogger.Error("Error " + err.Error() + "occured during executing View Render")
		color.Red(" - respnose-generator.go  View  : %s", err.Error())
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
		DefaultLogger.Error("Error " + err.Error() + "occured during rendering View Response")
		color.Red(" - respnose-generator.go  HtmlString : %s", err.Error())
	}
	vars := make(jet.VarMap)
	if err = t.Execute(&html, vars, data); err != nil {
		DefaultLogger.Error("Error " + err.Error() + "occured during executing View Render")
		color.Red(" - respnose-generator.go  HtmlString : %s", err.Error())
	}
	return html.String()
}

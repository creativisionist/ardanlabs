package main

import (
  "bytes"
  "log"
  "net/http"
  "text/template"

  "github.com/gorilla/mux"
)

// Basic struct to hold basic page data variables
type PageData struct {
  Title string
  Body string
}

func main() {
  // Create a router
  rt := mux.NewRouter().StrictSlash(true)

  // Add the "index" or root path
  rt.HandleFunc("/", Index)

  // Fire up the server
  log.Println("Starting server on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", rt))
}

// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
  // Fille out page data for index
  pd := PageData{
    Title: "Index Page",
    Body: "Page body...",
  }

  // Render a template with our page data
  tmpl, err := htmlTemplate(pd)

  // If we got an error, write it out and exit
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }

  // All went well -- write out template
  w.Write([]byte(tmpl))
}

func htmlTemplate(pd PageData) (string, error) {
  // Define basic HTML template
  html := `<HTML>
  <head><title>{{.Title}}</title></head>
  <body>
  {{.Body}}
  </body>
  </HTML>`

  // Parse template
  tmpl, err := template.New("index").Parse(html)
  if err != nil {
    return "", err
  }

  // Need location to write executed template to
  var out bytes.Buffer

  // Render template w/data we passed in
  if err := tmpl.Execute(&out, pd); err != nil {
    // If we couldn't render, return an error
    return "", err
  }

  // Return the template
  return out.String(), nil
}

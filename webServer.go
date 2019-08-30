package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
  "net/url"
)

func HandlerGetEnv(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, time.Now().Format("2006-01-02 15:04:05"))
  
  queryForm, err := url.ParseQuery(r.URL.RawQuery)
  if err == nil && len(queryForm["env"]) > 0 {
      fmt.Fprintf(w, "ENV %s: %s", queryForm["env"][0], os.Getenv(queryForm["env"][0]))
      fmt.Fprintln(w, )
  } else {
      fmt.Fprintln(w, "Usage: http://{ip}:8080?env={envName}")
  }
  
}

func main() {
  http.HandleFunc("/", HandlerGetEnv)
  http.ListenAndServe(":8080", nil)
}

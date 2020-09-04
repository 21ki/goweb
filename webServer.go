package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
  "net/url"
)

func HandlerUsage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, time.Now().Format("2006-01-02 15:04:05"))
  fmt.Fprintln(w, "Usage:")
  fmt.Fprintln(w, "Show usage: http://{ip}:8080")
  fmt.Fprintln(w, "Get env by name: http://{ip}:8080/getenv?env={envName}")
  fmt.Fprintln(w, "Get hostname: http://{ip}:8080/hostname")
  fmt.Fprintln(w, "Show request info: http://{ip}:8080/requestinfo")
}

func HandlerGetEnv(w http.ResponseWriter, r *http.Request) {
  
  queryForm, err := url.ParseQuery(r.URL.RawQuery)
  if err == nil && len(queryForm["env"]) > 0 {
      fmt.Fprintf(w, "ENV %s: %s", queryForm["env"][0], os.Getenv(queryForm["env"][0]))
  } else {
      HandlerUsage(w, r)
  }
  
}

func HandleRequestInfo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "remoteAddr:",r.RemoteAddr)
  for k,v := range r.Header{
    fmt.Fprintln(w,k,":", v)
  }
}


func HandleHostname(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, os.Getenv("HOSTNAME"))
}

func main() {
  http.HandleFunc("/", HandlerUsage)
  http.HandleFunc("/getenv", HandlerGetEnv)
  http.HandleFunc("/requestinfo", HandleRequestInfo)
  http.HandleFunc("/hostname", HandleHostname)

  http.ListenAndServe(":8080", nil)
}

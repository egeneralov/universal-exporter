package main

import (
  "fmt"
  "flag"
  "strconv"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  
  "github.com/gosimple/slug"

)


var (
  workMode string
  bind string
  scrapeURL string
  itemsDelimiter string
  kvDelimiter string
)


func getRawData (url string) ([]byte) {
  resp, err := http.Get(url)
  if err != nil { panic(err) }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil { panic(err) }
  return body
}

func parseBody (body []byte) (interface{}) {
  var payload interface{}
  payload = map[string]interface{}{}

  for _, line := range strings.Split(string(body), itemsDelimiter) {
    key := ""
    value := ""
    for _, kv := range strings.Split(line, kvDelimiter) {
      if key == "" { key = kv } else { value = kv }
    }
    if key == "" { continue }
    
    key = slug.Make(key)
    
    if isInt(value) {
      intValue , _ := strconv.Atoi(value)
      payload.(map[string]interface{})[key] = intValue
    } else {
      payload.(map[string]interface{})[key] = value
    }
  }
  return payload
}

func isInt(payload string) (bool) {
  _, err := strconv.Atoi(payload)
  return err == nil
}

func isStartsWithInt(payload string) (bool) {
  return isInt(string([]byte(payload)[0]))
}

func interfaceToJson(payload interface{}) (string) {
  resultJson, err := json.Marshal(payload)
  if err != nil { panic(err) }
  return string(resultJson)
}

func Keys(m map[string]interface{}) (keys []string) {
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}

func getData () (map[string]interface {}) {
  body := getRawData(scrapeURL)
  payload := parseBody(body)
  md := payload.(map[string]interface{})
  return md
}

func toPrometheus (md map[string]interface {}) (string) {
  var answer string
  for _, k := range Keys(md) {
    key := strings.Replace(k, "-" , "_", -1)
    if isStartsWithInt(key) {
      key = fmt.Sprintf("s%s", key)
    }
    
    switch md[k].(type) {
      case string:
        answer = fmt.Sprintf("%s# HELP %s %s \n# TYPE %s gauge\n%s{value=\"%s\"} 1\n", answer, key, key, key, key, md[k])
      case int:
        answer = fmt.Sprintf("%s# HELP %s %s \n# TYPE %s gauge\n%s %d\n", answer, key, key, key, key, md[k])
      default:
    }
    
  }
  return answer
}

func toJson (payload map[string]interface {}) (string) {
  return interfaceToJson(payload)
}



func main() {
  
  flag.StringVar(&workMode, "mode", "json", "what to do [json,prometheus,exporter]")
  flag.StringVar(&bind, "bind", "0.0.0.0:9682", "bind to (golang http server)")
  flag.StringVar(&scrapeURL, "url", "http://10.0.0.1/cgi-bin/sysconf.cgi?page=ajax&action=get_status", "url to parse")
  flag.StringVar(&itemsDelimiter, "itemsDelimiter", "\n", "items (or lines) delimiter")
  flag.StringVar(&kvDelimiter, "kvDelimiter", "=", "key/value delimiter")
  flag.Parse()
  
  md := getData()
  
  switch workMode {
    case "prometheus":
      fmt.Println(toPrometheus(md))
    case "exporter":
      http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, toPrometheus(getData()))
      })
      fmt.Printf("binding exporter to %s\n", bind)
      err := http.ListenAndServe(bind, nil)
      if err != nil { panic(err) }
    default:
      fmt.Println(toJson(md))
  }
  
}
























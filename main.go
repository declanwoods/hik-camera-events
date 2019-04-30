package main

import "fmt"
import "net/http"
import "bufio"
import "strings"

func main() {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "http://192.168.0.131:80/ISAPI/Event/notification/alertStream",nil)
  req.SetBasicAuth("admin","hikvision1")
  res, _ := client.Do(req)

  if res != nil {
    reader := bufio.NewReader(res.Body)

    etype := ""
    datetime := ""
    state := ""

    for {
      line, _ := reader.ReadBytes('\n')
      if string(line) == "" {
        break
      }

      s := strings.Split(string(line), "<eventType>")
      if len(s) > 1 {
        x := strings.Split(s[1], "</eventType>")
        etype = strings.TrimSpace(x[0])
      }  
      s = strings.Split(string(line), "<dateTime>")
      if len(s) > 1 {
        x := strings.Split(s[1], "</dateTime>")
        datetime = strings.TrimSpace(x[0])
      }  
      s = strings.Split(string(line), "<eventState>")
      if len(s) > 1 {
        x := strings.Split(s[1], "</eventState>")
        state = strings.TrimSpace(x[0])
      }        

      if strings.Index(string(line), "--boundary") > -1 {
        if etype != "videoloss" && etype != "" {
          fmt.Println("New event,", etype, datetime, state)
        }
      }
    }
  }
}
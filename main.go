package main

import (
  "fmt"
  "time"
  "net/http"
  "bytes"
  "log"
)

var (
  arr []string
  status1 string
  status2 string
  token string
)

func ClearConsole() {
  fmt.Print("\033[H\033[2J")
}

func main() {
  fmt.Println("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mDiscord status changer \u001b[32;5;34m| \u001b[32;5;37mDevloped by github.com/vertionn\n")
  
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mStatus one: \u001b[32;5;34m")
  fmt.Scanln(&status1)
  
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mStatus two: \u001b[32;5;34m")
  fmt.Scanln(&status2)

  arr = append(arr, status1, status2, "github.com/vertionn")
   
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mToken: \u001b[32;5;34m")
  fmt.Scanln(&token)
  
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mStarting\u001b[32;5;34m? ")
  fmt.Scanln()

  ClearConsole()

  fmt.Println("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mDiscord status changer \u001b[32;5;34m| \u001b[32;5;37mDevloped by github.com/vertionn\n")
  
	for true {
    for i, v := range arr {
      time.Sleep(5 * time.Second)
      text := arr[i]
      var jsonStr = []byte(`{"custom_status":{"text":"` + text + `"}}`)
      requests, err := http.NewRequest("PATCH", "https://discord.com/api/v9/users/@me/settings",bytes.NewBuffer(jsonStr))
      if err != nil {
        log.Fatalln(err)
      }
      requests.Header.Set("authorization",token)
      requests.Header.Set("content-type","application/json")
      client := &http.Client{}
      response, err := client.Do(requests)
      if err != nil {
        log.Fatalln(err)
      }
  
      defer response.Body.Close()

      if response.StatusCode == 200 {
        fmt.Printf(`%v[%v+%v] %vUpdated Status To "%v%v%v"%v`,"\u001b[32;5;34m","\u001b[32;5;37m","\u001b[32;5;34m","\u001b[32;5;37m","\u001b[32;5;34m",v,"\u001b[32;5;37m","\n")
      } else {
        fmt.Printf(`%v[%v-%v] %va error has occured%v`, "\033[31m", "\033[37m", "\033[31m", "\033[37m", "\n")
      }
      
    }
  }
}

package main

import (
	"fmt"
  "time"
  "math/rand"
  "net/http"
  "bytes"
  "log"
)

func main() {
  fmt.Println("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mDiscord status changer \u001b[32;5;34m| \u001b[32;5;37mDevloped by github.com/vertionn\n")
  var status1 string
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mStatus one: \u001b[32;5;34m")
  fmt.Scanln(&status1)
  
  var status2 string
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mStatus two: \u001b[32;5;34m")
  fmt.Scanln(&status2)

  arr := []string{}
  arr = append(arr, status1, status2, "github.com/vertionn")
   
  var token string
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mToken: \u001b[32;5;34m")
  fmt.Scanln(&token)
  
  fmt.Printf("\u001b[32;5;34m[\u001b[32;5;37m?\u001b[32;5;34m] \u001b[32;5;37mStarting\u001b[32;5;34m? ")
  fmt.Scanln()
  
	for true {
    time.Sleep(5 * time.Second)
    randomIndex := rand.Intn(len(arr))
    ooo := arr[randomIndex]
    var jsonStr = []byte(`{"custom_status":{"text":"` + ooo + `"}}`)
    requests, err := http.NewRequest("PATCH", "https://discord.com/api/v9/users/@me/settings",bytes.NewBuffer(jsonStr))
    if err != nil {
      log.Fatalln(err)
    }
    requests.Header.Set("authorization",token)
    requests.Header.Set("content-type","application/json")
    client := &http.Client{}
    respopnse, err := client.Do(requests)
    if err != nil {
      log.Fatalln(err)
    }

    defer respopnse.Body.Close()


    fmt.Printf(`%v[%v+%v] %vUpdated Status To "%v%v%v"%v`,"\u001b[32;5;34m","\u001b[32;5;37m","\u001b[32;5;34m","\u001b[32;5;37m","\u001b[32;5;34m",ooo,"\u001b[32;5;37m","\n")
    
  }
}

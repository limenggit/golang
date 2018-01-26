package main
import "net/http"
import "io/ioutil"
import "fmt"
import "time"
func loop() {
  fmt.Println("loop start")
  client := &http.Client{
     }
     req,_  := http.NewRequest("GET", "http://www.baidu.com", nil)
     req.Header.Add("Cookie", "X_SESSION_ID=aa")
     req.Header.Add("Content-Type", "application/json; charset=utf-8")
     resp, _ := client.Do(req)
     body,_ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
func main() {
   go loop()
   go loop()
   time.Sleep(3 * time.Second)
}
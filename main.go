package main


import "fmt"

func main(){
   client := NewClient()
   weather,_:=client.Request("chongqing")
   fmt.Printf("Weather in %s: %.1f\u2103 \n", "chongqing", weather.Temperature)
}

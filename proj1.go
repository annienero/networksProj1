package main

import (
    "net"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
   var port string = "27993"
  // var port string = "80"
  if (len(os.Args) == 5 && os.Args[1] == "-p") {
    port = os.Args[2]
  }
  tcpAddr, err := net.ResolveTCPAddr("tcp4", "cbw.sh:" + port)
  if err != nil {
      fmt.Println("error1")
  }
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  if err != nil {
      fmt.Println("error2")
  }
  fmt.Println(conn)

  _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
  if err != nil {
      fmt.Println("error3")
  }
  result, err := ioutil.ReadAll(conn)
  fmt.Println(result)
}

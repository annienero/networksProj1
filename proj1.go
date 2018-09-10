package main

import (
    "net"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
  var port string = "27993"
  //var port string = "80"
  if (len(os.Args) == 5 && os.Args[1] == "-p") {
    port = os.Args[2]
  }
  tcpAddr, err := net.ResolveTCPAddr("tcp4", "cbw.sh:" + port)
  checkError(err)
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)
  fmt.Println("right before printing conn")
  fmt.Println(conn)

  _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
  checkError(err)
  result, err := ioutil.ReadAll(conn)
  fmt.Println(result)
}

func checkError(err) {
  if err != nil {
      fmt.Println(err)
  }
}

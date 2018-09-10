package main

import (
    "net"
    "fmt"
    "os"
    "time"
)

func main() {
  var port string = "27993"
  var helloMsg string = "cs3700fall2018 HELLO 001698478\n"

  if (len(os.Args) == 5 && os.Args[1] == "-p") {
    port = os.Args[2]
  }

  tcpAddr, err := net.ResolveTCPAddr("tcp4", "cbw.sh:" + port)
  checkError(err)
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)
  conn.SetReadDeadline(time.Now().Add(time.Second))
  _, err = conn.Write([]byte(helloMsg))
  checkError(err)
  
  var tmp = make([]byte, 256)
  _, err = conn.Read(tmp)
  checkError(err)
  var message = ""
  for tmp[0] != 0 {
    var str = string(tmp[:256])
    message += str
    tmp = make([]byte, 256)
    conn.Read(tmp)
  }
}

func checkError(err error) {
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
}

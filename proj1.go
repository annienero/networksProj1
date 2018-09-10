package main

import (
    "net"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
  //var port string = "27993"
  //var port string = "80" //Does string addition with '+' really work in go?

  //var host string = "login-faculty.ccs.neu.edu"
  //var helloMsg string = "cs3700fall2018 HELLO 001698478\n"
  //var testMsg string = "HEAD / HTTP/1.0\r\n\r\n"
  if (len(os.Args) == 5 && os.Args[1] == "-p") {
    //port = os.Args[2]
  }
  tcpAddr, err := net.ResolveTCPAddr("tcp4", "cbw.sh:27993")
  checkError(err)
  fmt.Println("got tcpAddr")
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)
  fmt.Println(conn)

  _, err = conn.Write([]byte("cs3700fall2018 HELLO 001698478\n"))
  checkError(err)
  result, err := ioutil.ReadAll(conn)
  fmt.Println(result)
}

func checkError(err error) {
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
}

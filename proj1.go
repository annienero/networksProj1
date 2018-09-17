package main

import (
    "net"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
  port := "27993"
  helloMsg := "cs3700fall2018 HELLO 001698478\n"
  countMsg := "cs3700fall2018 COUNT "


  if (len(os.Args) == 5 && os.Args[1] == "-p") {
    port = os.Args[2]
  }

  tcpAddr, err := net.ResolveTCPAddr("tcp4", "cbw.sh:" + port)
  checkError(err)
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)
  _, err = conn.Write([]byte(helloMsg))
  checkError(err)

  tokens := strings.Split(readFromConnection(conn), " ")
  messageType := tokens[1]
  arguments := tokens[2:len(tokens)]

  for messageType != "BYE" {
    count := countMsg + strconv.Itoa(evalFind(arguments)) + "\n"
    fmt.Println("countMessage = " + count)
    _, err = conn.Write([]byte(count))
    checkError(err)

    tokens = strings.Split(readFromConnection(conn), " ")
    messageType = tokens[1]
    arguments = tokens[2:len(tokens)]
  }

  // print secret flag
  fmt.Println(arguments[0])

  conn.Close()
}

func checkError(err error) {
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
}

func readFromConnection(conn *net.TCPConn) string {
  tmp := make([]byte, 256)
  bytes_read, err := conn.Read(tmp)
  checkError(err)
  message := ""
  for bytes_read == 256 {
    str := string(tmp[:256])
    message += str
    tmp = make([]byte, 256)
    bytes_read, err = conn.Read(tmp)
    checkError(err)
  }
  fmt.Println("bytes read " + strconv.Itoa(bytes_read))
  message += string(tmp[:256])
  fmt.Println("full message is " + message)
  return message
}

func evalFind(arguments []string) int {
  toFind := arguments[0]
  toSearch := arguments[1]
  return strings.Count(toSearch, toFind)
}

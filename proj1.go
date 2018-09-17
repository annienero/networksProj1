package main

import (
    "net"
    "fmt"
    "os"
    "time"
    "strings"
    "strconv"
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
  _, err = conn.Write([]byte(executeAndGetReply(message)))
  checkError(err)
}

func checkError(err error) {
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
}

//A message is a one of the 4 possible messages in this protocol, as described by the assignment
//A message's type is one of: [HELLO, FIND, COUNT, BYE]
func executeAndGetReply(messageString string) string {

  tokens := strings.Split(messageString, " ")
  messageType := tokens[1]
  arguments := tokens[2:len(tokens)]

  if messageType == "FIND" {
    return "cs3700fall2018 COUNT " + strconv.Itoa(evalFind(arguments)) + "\n"
  } else {  
    return "Bye case not implemented yet."
  }
}

func evalFind(arguments []string) int {
  toFind := arguments[0]
  toSearch := arguments[1]
  return strings.Count(toSearch, toFind)
}

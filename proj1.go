package main

import (
    "net"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
  //constants
  port := "27993"
  countMsg := "cs3700fall2018 COUNT "
  hostname := os.Args[len(os.Args) - 2] + ":"
  id := os.Args[len(os.Args) - 1]
  helloMsg := "cs3700fall2018 HELLO " + id + "\n"

  // if a port has been specified
  if (os.Args[1] == "-p") {
    port = os.Args[2]
  }

  //set up TCP connection
  tcpAddr, err := net.ResolveTCPAddr("tcp4", hostname + port)
  checkError(err)
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)

  //Send the initial HELLO message
  _, err = conn.Write([]byte(helloMsg))
  checkError(err)

  //tokenize the message into its important parts
  tokens := strings.Split(readFromConnection(conn), " ")
  messageType := tokens[1]
  arguments := tokens[2:len(tokens)]

  //loop executing and replying to the messages until we receive a BYE message
  for messageType != "BYE" {
    count := countMsg + strconv.Itoa(evalFind(arguments)) + "\n"
    _, err = conn.Write([]byte(count))
    checkError(err)

    tokens = strings.Split(readFromConnection(conn), " ")
    messageType = tokens[1]
    arguments = tokens[2:len(tokens)]
  }

  // print secret flag
  fmt.Println(string(arguments[0]))

  conn.Close() //close connection
}

func checkError(err error) {
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
}

//Read from the given connection until we read a newline character
func readFromConnection(conn *net.TCPConn) string {
  tmp := make([]byte, 256) //we read in 256 byte chunks -- consider making this a bit bigger?
  _, err := conn.Read(tmp)
  checkError(err)
  message := "" //we'll accumulate the message here
  for {
    str := string(tmp[:256]) //parse our chunk of bytes into a string
    message += str //add it to the message
    if strings.Contains(message, "\n") { //if it has a newline, we read the whole message, and we can stop grabbing chunks.
      break
    }
    //grab the next byte chunk...
    tmp = make([]byte, 256)
    _, err = conn.Read(tmp)
    checkError(err)
  }
  return message
}

//Returns the count, given the arguments of a FIND message
func evalFind(arguments []string) int {
  toFind := arguments[0]
  toSearch := arguments[1]
  return strings.Count(toSearch, toFind)
}

package main

import (
    "net"
    "fmt"
    "os"
    "strings"
    "strconv"
    "time"
)

func main() {
  //constants
  port := "27993"
  helloMsg := "cs3700fall2018 HELLO 001254621\n"
  countMsg := "cs3700fall2018 COUNT "


  if (len(os.Args) == 5 && os.Args[1] == "-p") {
    port = os.Args[2]
  }

  //set up TCP connection
  tcpAddr, err := net.ResolveTCPAddr("tcp4", "cbw.sh:" + port)
  checkError(err)
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkError(err)
  conn.SetReadDeadline(time.Now().Add(time.Second * 2)) //does this actually help us? SetReadDeadline just sets a 1-time deadline it doesn't refresh every time we Read()?
  
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
  fmt.Println("secret flag = " + string(arguments[0]))

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

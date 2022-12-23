package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	TwitchHost = "irc.chat.twitch.tv:6667"
)

func main() {
	itest := 1
	for true {
		println("running: " + strconv.Itoa(itest))
		time.Sleep(5 * time.Second)
		itest++
	}
	//conn := connect()
	//logon(conn)
	//
	//tp := textproto.NewReader(bufio.NewReader(conn))
	//
	//for {
	//	status, err := tp.ReadLine()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(status)
	//}

}
func connect() net.Conn {
	conn, err := net.Dial("tcp", TwitchHost)
	if err != nil {
		panic(err)
	}

	return conn
}
func logon(conn net.Conn) {
	sendData(conn, "PASS oauth:gwtrd8hyz8vtdyvipntsqwbiawseo7")
	sendData(conn, "NICK CookieNinja")
	sendData(conn, "CAP REQ :twitch.tv/commands twitch.tv/tags twitch.tv/membership")
	sendData(conn, "JOIN #xemtek")
}
func sendData(conn net.Conn, message string) {
	fmt.Fprintf(conn, "%s\r\n", message)
}

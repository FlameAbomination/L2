package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var timeFlag = flag.Duration("timeout", time.Duration(timeout), "таймаут на подключение к серверу")

var timeout int

func main() {
	flag.Parse()

	conn, _ := net.DialTimeout("tcp", flag.Args()[0]+":"+flag.Args()[1], *timeFlag)
	inReader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer func() {
		conn.Close()
	}()
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
loop:
	for {
		select {
		case <-cancelChan:
			break loop
		default:
			fmt.Print("Text to send: ")
			text, err := inReader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			_, err = connReader.WriteString(text)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			message, err := connReader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Print("Message from server: " + message)

		}
	}
}

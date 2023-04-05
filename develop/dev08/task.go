package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var current_path string

func execute(str string) (bool, error) {
	str = strings.ReplaceAll(str, "\r\n", "")
	cmd := strings.Split(str, " ")[0]
	args := strings.Split(str, " ")[1:]
	switch cmd {
	case "cd":
		if len(args) != 1 {
			return false, errors.New("too many arguments")
		}
		current_path = filepath.Join(current_path, args[0])
	case "pwd":
		fmt.Fprintln(os.Stdout, current_path)
	case "echo":
		fmt.Fprintln(os.Stdout, strings.Join(args, " "))
	case "kill":
		if len(args) != 1 {
			return false, errors.New("too many arguments")
		}
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			return false, err
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return false, err
		}
		proc.Kill()
		if err != nil {
			return false, err
		}
	case "exit":
		return true, nil
	case "ps":
		processList, err := ps.Processes()
		if err != nil {
			return false, err
		}

		for x := range processList {
			process := processList[x]
			fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
		}
	case "fork":
		//syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	case "exec":
		if len(args) != 1 {
			return false, errors.New("too many arguments")
		}
		cmd := exec.Command(args[0])
		if errors.Is(cmd.Err, exec.ErrDot) {
			cmd.Err = nil
		}
		if err := cmd.Run(); err != nil {
			return false, nil
		}
	default:
		fmt.Fprintln(os.Stderr, cmd, len(cmd))
		return false, errors.New("unknown command")
	}
	return false, nil
}

func main() {
	var err error
	current_path, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	done := false
	scanner := bufio.NewReader(os.Stdin)
	for !done {
		text, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		done, err = execute(text)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if done {
			break
		}
	}
}

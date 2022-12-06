// main.go

package main

import (
	"bufio"
	"fmt"
	"goscript/lexer"
	"goscript/repl"
	"goscript/token"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) == 2 {
		//读取本地脚本
		scriptName := os.Args[1]
		readFile, err := os.Open(scriptName)
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		var lines []string
		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}
		readFile.Close()
		for _, line := range lines {
			l := lexer.New(line)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				if tok.Type == token.ILLEGAL {
					panic("脚本解析错误")
				}
				fmt.Printf("%+v\n", tok)
			}
		}

	} else {
		//进入repl shell
		//let add = fn(x, y) { x + y; };
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}

}

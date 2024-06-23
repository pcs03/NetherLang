package main

import (
	"fmt"
    "os"
    "os/user"
    "netherlang/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Welkom %s! Dit is de NetherLang Programeertaal!\n", user.Username)
    repl.Start(os.Stdin, os.Stdout)
}

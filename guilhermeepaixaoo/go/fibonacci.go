package main

import (
    "os"
    "fmt"
    "strconv"
    "reflect"
)   

func main() {
    
    args := os.Args[1:]

    if len(args) == 1 {
        n, err := strconv.ParseInt(args[0], 10, 32)
        if err != nil {
            fmt.Println("Texto não é parâmetro para Fibonacci")
            return
        }

        f, l := 0
        for i := 0; i < int(n); i++ {
            
        }
        return
    }
    fmt.Println("Passe algum parâmetro para Fibonacci")
    return

}
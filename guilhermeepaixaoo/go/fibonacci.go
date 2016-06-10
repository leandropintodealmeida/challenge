package main

import (
    "os"
    "fmt"
    "strconv"
)   

func main() {
    
    args := os.Args[1:]

    if len(args) == 1 {
        n, err := strconv.ParseInt(args[0], 10, 32)
        if err != nil {
            fmt.Println("Texto não é parâmetro para Fibonacci")
            return
        }

        a := 0
        b := 1
        r := 0
        for i := 1; i < int(n); i++ {
            r = a + b
            a = b
            b = r
        }
        fmt.Println(r)
        return
    }
    fmt.Println("Passe algum parâmetro para Fibonacci")
    return

}
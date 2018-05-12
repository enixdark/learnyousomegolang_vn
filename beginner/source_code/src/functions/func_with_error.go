package main

import (
	"fmt"
	"errors"
)

func divide(a float32, b float32) (error, float32) {
    if b == 0 {
        err := errors.New("Không thể  thực hiện được phép chia")
        return err, 0.0
    }
    return nil, a / b
}

func main() {
    err, result := divide(1, 0)
    if err != nil {
        fmt.Println(err)
    }
    err, result = divide(1, 1)

    fmt.Println(result)
}
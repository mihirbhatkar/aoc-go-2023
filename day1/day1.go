package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Println("Error reading file:", err)
    }

    scanner := bufio.NewScanner(file)

    var slice []string

    for scanner.Scan() {
        line := scanner.Text()
        slice = append(slice, line)
    }

    var score int = 0

    for _, str := range slice {

        firstNumChecked := false
        number := []int{-1, -1}
        for _, char := range str {
            if '0' <= char && char <= '9' {
                if firstNumChecked == false {
                    firstNumChecked = true
                    intVal, _ := strconv.Atoi(string(char))
                    number[0] = intVal
                } else {
                    intVal, _ := strconv.Atoi(string(char))
                    number[1] = intVal
                }
            }
        }
        if number[1] == -1 {
            number[1] = number[0]
        }

        score += number[0]*10 + number[1]
    }

    fmt.Println(score)
}

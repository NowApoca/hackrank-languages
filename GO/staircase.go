package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the staircase function below.
func staircase(n int32) {
    var numIndex, numIndex2 int32 = 0,0
    for numIndex = 0 ; numIndex < n ; numIndex++ {
        for numIndex2 = 0 ; numIndex2 < n ; numIndex2++ {
                if numIndex >= (n - numIndex2 -1) {
                fmt.Printf("#")
            }else{
                fmt.Printf(" ")
            }
        }
        fmt.Printf("\n")
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}



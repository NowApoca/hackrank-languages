package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
    arrayLength := len(arr)
    var numPositive, numNegative, numZero int = 0,0,0
    for numIndex := 0 ; numIndex < arrayLength ; numIndex ++ {
        num := arr[numIndex]
        switch true {
            case num > 0:
                numPositive++
            break
            case num < 0:
                numNegative++
            break
            case num == 0:
                numZero++
            break
        }
    }
    fmt.Printf("%.6f\n", float64(numPositive) / float64(arrayLength))
    fmt.Printf("%.6f\n", float64(numNegative) / float64(arrayLength))
    fmt.Printf("%.6f\n", float64(numZero) / float64(arrayLength))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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


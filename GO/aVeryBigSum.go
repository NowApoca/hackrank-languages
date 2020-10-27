package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	arrayLength := len(ar);
	var sum int64;
	for numIndex := 0 ; numIndex < arrayLength; numIndex++ {
		sum += ar[numIndex]
	}
	return sum;
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int64

    for i := 0; i < int(arCount); i++ {
        arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        ar = append(ar, arItem)
    }

    result := aVeryBigSum(ar)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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


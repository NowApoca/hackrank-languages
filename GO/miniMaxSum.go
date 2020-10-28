package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
    var arrayLength = len(arr)
    var min,max int32 = 0,0
    var sortedArr []int32 = arrayMergeSorter(arr) 
    for numIndex := 0; numIndex < (arrayLength -1) ; numIndex++{
        min += sortedArr[numIndex]
        max += sortedArr[arrayLength - 1 - numIndex]
    }
    fmt.Printf("%d %d", min, max)
}

func arrayMergeSorter(arr []int32) []int32{
    var arrayLength int = len(arr)
    var middleIndex int = arrayLength/2
    var sortedArr1 []int32 = sortArrayRecursively(arr[:middleIndex]);
    var sortedArr2 []int32 = sortArrayRecursively(arr[middleIndex:arrayLength]);
    var sortedArr = append(sortedArr1, sortedArr2...)
    return sortArrayRecursively(sortedArr);
}

func sortArrayRecursively (arr []int32) []int32{
    var sortedArr []int32 = arr;
    for numIndex := 0 ; numIndex < len(sortedArr) ; numIndex++ {
        for numIndex2 := 0 ; numIndex2 < (len(sortedArr) - 1) ; numIndex2++ {
            if sortedArr[numIndex2] > sortedArr[numIndex2 + 1] {
                sortedArr[numIndex2], sortedArr[numIndex2 + 1] = sortedArr[numIndex2 + 1] , sortedArr[numIndex2] 
            }
        }
    }
    return sortedArr;
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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



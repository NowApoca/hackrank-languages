#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'diagonalDifference' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY arr as parameter.
#

def diagonalDifference(arr):
    # Write your code here
    rangeMatrice = len(arr)
    firstDiagonalArr = []
    for x in range(rangeMatrice):
        firstDiagonalArr.append([x,x])
    secondDiagonalArr = []
    for y in range(rangeMatrice):
        secondDiagonalArr.append([y,rangeMatrice -y])
    sum1 = 0
    for x1 in firstDiagonalArr:
        sum1 += arr[x1[0]][x1[1]];
    sum2 = 0
    for y2 in secondDiagonalArr:
        sum2 += arr[y2[0]][y2[1]-1];
    return abs(sum1 - sum2)

# arr = [[1,4,6],[2,1,3],[7,5,3]]
# print(diagonalDifference(arr))

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    arr = []

    for _ in range(n):
        arr.append(list(map(int, input().rstrip().split())))

    result = diagonalDifference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()

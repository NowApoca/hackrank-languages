#!/bin/python3
# note : check print

import math
import os
import random
import re
import sys

# Complete the plusMinus function below.
def plusMinus(arr):
    negativeCount = 0
    positiveCount = 0
    sum = len(arr)
    for x in arr:
        if x > 0:
            positiveCount += 1
        if x < 0:
            negativeCount += 1
    positiveRate = (positiveCount*1.0)/sum
    negativeRate = (negativeCount*1.0)/sum
    print('{0:.6f}'.format(positiveRate))
    print('{0:.6f}'.format(negativeRate))
    print('{0:.6f}'.format(1-positiveRate-negativeRate))
    return ['{0:.6f}'.format(positiveRate),'{0:.6f}'.format(negativeRate), '{0:.6f}'.format(1-positiveRate-negativeRate)]
    
# print(plusMinus([-4, 3, -9, 0, 4, 1]))

if __name__ == '__main__':
    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)

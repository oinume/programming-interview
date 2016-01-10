# http://www.ardendertat.com/2011/09/17/programming-interview-questions-1-array-pair-sum/

def pairSum1(arr, k):
    if len(arr)<2:
        return
    arr.sort()
    left, right = (0, len(arr)-1)
    while left<right:
        currentSum=arr[left]+arr[right]
        if currentSum==k:
            print arr[left], arr[right]
            left+=1 #or right-=1
        elif currentSum<k:
            left+=1
        else:
            right-=1

pairSum1([ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 ], 10)

# http://www.ardendertat.com/2011/09/24/programming-interview-questions-3-largest-continuous-sum/

def largestContinuousSum(array):
    if len(array) == 0:
        return 0
    maxSum = currentSum = array[0]
    for num in array[1:]:
        currentSum = max(currentSum+num, num)
        maxSum = max(currentSum, maxSum)
    return maxSum

if __name__ == '__main__':
    print(largestContinuousSum([100, 6, -46])) # 106
    print(largestContinuousSum([100, 6, -46, 10, 100])) # 170

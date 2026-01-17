def steps(number):
    if number <= 0:
        raise ValueError("Only positive integers are allowed")

    n = number
    count = 0

    while n != 1:
        if n & 1:          # odd
            n = n * 3 + 1
        else:              # even
            n >>= 1        # divide by 2
        count += 1

    return count

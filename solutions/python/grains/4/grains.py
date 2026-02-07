MAX = 64
TOTAL = 18446744073709551615

def square(num):
    if num < 1 or num > MAX:
        raise ValueError("square must be between 1 and 64")
    return 1 << (num - 1)

def total():
    return TOTAL
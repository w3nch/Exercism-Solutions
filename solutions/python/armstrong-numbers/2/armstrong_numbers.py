def is_armstrong_number(n):
    s = str(n)
    return n == sum(int(d) ** len(s) for d in s)

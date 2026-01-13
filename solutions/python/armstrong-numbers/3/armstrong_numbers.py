from functools import lru_cache

@lru_cache(None)
def p(d, k):
    return d ** k

def is_armstrong_number(n):
    s = str(n)
    k = len(s)
    return n == sum(p(int(d), k) for d in s)

def is_armstrong_number(num):
    s = str(num)
    power = len(s)
    return num == sum(int(d) ** power for d in s)
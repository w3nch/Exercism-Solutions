def is_armstrong_number(number):
    count = str(number)
    powers = len(count)
    armstrong_number = 0
    for i in count:
        armstrong_number += int(i) ** powers
    return armstrong_number == number
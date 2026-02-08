def is_armstrong_number(num):
    list_num = list(map(int, str(num)))
    num_len = len(list_num)
    numbers = 0
    for i , v in enumerate(list_num):
        numbers += v ** num_len
    return numbers == num 
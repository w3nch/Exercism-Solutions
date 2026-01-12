import re

def is_isogram(string):
    new_replaced_string = string.replace("-", "").replace(" ", "").lower()
    return len(new_replaced_string) == len(set(new_replaced_string))
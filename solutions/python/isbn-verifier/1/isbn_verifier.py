def is_valid(isbn):
    cleaned = str(isbn).replace("-", "")
    if len(cleaned) != 10:
        return False

    total = 0
    weight = 10

    for idx, i in enumerate(cleaned):
        if i == "X":
            if idx != 9:
                return False
            value = 10
        elif i.isdigit():
            value = int(i)
        else:
            return False

        total += value * weight
        weight -= 1

    return total % 11 == 0

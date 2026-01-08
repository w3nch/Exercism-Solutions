def classify(number):
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """
    if number <= 0:
        raise ValueError("Classification is only possible for positive integers.")
    if number == 1:
        return "deficient"

    is_perfect = 0

    for i in range(1, int(number ** 0.5) + 1):
        if number % i == 0:
            is_perfect += i
            if i != number // i and i != 1:
                is_perfect += number // i

    if is_perfect == number:
        return "perfect"
    elif is_perfect > number:
        return "abundant"
    else:
        return "deficient"
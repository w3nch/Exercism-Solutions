def classify(number):
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """
    if number <= 0:
        raise ValueError("Classification is only possible for positive integers.")
    if number == 1:
        return "deficient"

    is_perfect = 1
    i = 2

    while i * i <= number:
        if number % i == 0:
            is_perfect += i
            if i != number // i:
                is_perfect += number // i

            if is_perfect > number:
                return "abundant"
        i += 1

    if is_perfect == number:
        return "perfect"
    elif is_perfect > number:
        return "abundant"
    else:
        return "deficient"
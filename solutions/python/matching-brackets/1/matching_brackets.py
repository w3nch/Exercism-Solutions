def is_paired(input_string):
    stack = []
    mapping = {
        '(': ')',
        '[': ']',
        '{': '}'
    }

    for ch in input_string:
        if ch in mapping:
            # push what we EXPECT to see later
            stack.append(mapping[ch])

        elif ch in ')]}':
            if not stack:
                return False

            expected = stack.pop()
            if ch != expected:
                return False

    return not stack

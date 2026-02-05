def commands(binary):
    actions = ["wink", "double blink", "close your eyes", "jump"]
    result = []

    for i in range(len(actions)):
        if len(binary) > i and binary[-(i + 1)] == "1":
            result.append(actions[i])

    if len(binary) >= 5 and binary[-5] == "1":
        result.reverse()

    return result

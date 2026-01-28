BAND_COLOR = [
    "black",
    "brown",
    "red",
    "orange",
    "yellow",
    "green",
    "blue",
    "violet",
    "grey",
    "white"
]

def value(colors):
    first = second = None

    for i, color in enumerate(BAND_COLOR):
        if color == colors[0]:
            first = i
        if color == colors[1]:
            second = i

    return first * 10 + second
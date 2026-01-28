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
    color1 = colors[0]
    color2 = colors[1]

    first = BAND_COLOR.index(color1)
    second = BAND_COLOR.index(color2)

    return first * 10 + second
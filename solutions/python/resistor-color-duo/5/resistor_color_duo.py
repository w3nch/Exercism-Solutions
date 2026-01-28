BAND_COLOR = [
    "black", "brown", "red", "orange", "yellow",
    "green", "blue", "violet", "grey", "white"
]

def value(colors):
    result = 0
    for color in colors[:2]:
        result = result * 10 + BAND_COLOR.index(color)
    return result
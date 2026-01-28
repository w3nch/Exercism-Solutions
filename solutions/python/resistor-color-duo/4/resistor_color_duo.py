BAND_COLOR = [
    "black", "brown", "red", "orange", "yellow",
    "green", "blue", "violet", "grey", "white"
]

def value(colors):
    digits = [BAND_COLOR.index(c) for c in colors]
    return digits[0] * 10 + digits[1]
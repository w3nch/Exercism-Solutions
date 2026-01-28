BANDS = [
    "black", "brown", "red", "orange", "yellow",
    "green", "blue", "violet", "grey", "white"
]

def label(colors):
    first = BANDS.index(colors[0])
    second = BANDS.index(colors[1])
    multiplier = 10 ** BANDS.index(colors[2])

    value = ((first * 10) + second) * multiplier

    if value < 1_000:
        return f"{value} ohms"
    elif value < 1_000_000:
        return f"{value // 1_000} kiloohms"
    elif value < 1_000_000_000:
        return f"{value // 1_000_000} megaohms"
    else:
        return f"{value // 1_000_000_000} gigaohms"
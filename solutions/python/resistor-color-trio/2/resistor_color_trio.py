BANDS = {
    "black": 0, "brown": 1, "red": 2, "orange": 3,
    "yellow": 4, "green": 5, "blue": 6,
    "violet": 7, "grey": 8, "white": 9
}

UNITS = (
    (1_000_000_000, "gigaohms"),
    (1_000_000, "megaohms"),
    (1_000, "kiloohms"),
    (1, "ohms"),
)

def label(colors):
    value = (
        (BANDS[colors[0]] * 10 + BANDS[colors[1]])
        * (10 ** BANDS[colors[2]])
    )
    if value == 0:
        return "0 ohms"

    for factor, unit in UNITS:
        if value >= factor:
            return f"{value // factor} {unit}"
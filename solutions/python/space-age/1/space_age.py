class SpaceAge:
    ORBITAL_PERIODS = {
        "earth": 1.0,
        "mercury": 0.2408467,
        "venus": 0.61519726,
        "mars": 1.8808158,
        "jupiter": 11.862615,
        "saturn": 29.447498,
        "uranus": 84.016846,
        "neptune": 164.79132,
    }

    EARTH_YEAR = 31557600

    def __init__(self, seconds):
        self.seconds = seconds

    def _age_on_planet(self, planet):
        return round(
            self.seconds / (self.EARTH_YEAR * self.ORBITAL_PERIODS[planet]),
            2
        )

    def on_earth(self):
        return self._age_on_planet("earth")

    def on_mercury(self):
        return self._age_on_planet("mercury")

    def on_venus(self):
        return self._age_on_planet("venus")

    def on_mars(self):
        return self._age_on_planet("mars")

    def on_jupiter(self):
        return self._age_on_planet("jupiter")

    def on_saturn(self):
        return self._age_on_planet("saturn")

    def on_uranus(self):
        return self._age_on_planet("uranus")

    def on_neptune(self):
        return self._age_on_planet("neptune")

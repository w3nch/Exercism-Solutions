EXPECTED_BAKE_TIME = 40

bake_time_remaining = lambda elapsed: 40 - elapsed; bake_time_remaining.__doc__="Remaining bake time"
preparation_time_in_minutes = lambda layers: layers * 2; preparation_time_in_minutes.__doc__="Prep time"
elapsed_time_in_minutes = lambda layers, elapsed: layers * 2 + elapsed; elapsed_time_in_minutes.__doc__="Total elapsed"
object Bob {
    fun hey(input: String): String {
        var hasLetter = false
        var isYelling = true
        var lastNonWhitespace: Char? = null

        for (c in input) {
            if (!c.isWhitespace()) {
                lastNonWhitespace = c
            }
            if (c.isLetter()) {
                hasLetter = true
                if (!c.isUpperCase()) {
                    isYelling = false
                }
            }
        }

        if (lastNonWhitespace == null) {
            return "Fine. Be that way!"
        }

        val isQuestion = lastNonWhitespace == '?'
        val yelling = hasLetter && isYelling

        return when {
            yelling && isQuestion -> "Calm down, I know what I'm doing!"
            yelling -> "Whoa, chill out!"
            isQuestion -> "Sure."
            else -> "Whatever."
        }
    }
}

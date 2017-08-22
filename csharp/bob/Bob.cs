using System;

public static class Bob {
    private const string silence = "Fine. Be that way!",
        yelling = "Whoa, chill out!",
        question = "Sure.",
        whatever = "Whatever.";

    public static string Response (string statement) {
        statement = statement.Trim ();
        if (isSilence (statement)) {
            return silence;
        }
        if (isYelling (statement)) {
            return yelling;
        }
        if (isQuestion (statement)) {
            return question;
        }
        return whatever;
    }

    private static bool isSilence (string statement) {
        if (string.IsNullOrWhiteSpace (statement)) {
            return true;
        }
        return false;
    }

    private static bool isYelling (string statement) {
        bool foundUpper = false;
        foreach (char c in statement) {
            if (char.IsLower (c)) {
                return false;
            }
            if (char.IsUpper (c)) {
                foundUpper = true;
            }
        }
        return foundUpper;
    }
    private static bool isQuestion (string statement) {
        if (statement.EndsWith ("?")) {
            return true;
        }
        return false;
    }
}
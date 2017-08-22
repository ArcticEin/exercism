using System;

public static class Raindrops {
    public static string Convert (int number) {
        bool fac3 = number % 3 == 0;
        bool fac5 = number % 5 == 0;
        bool fac7 = number % 7 == 0;
        if (fac3 || fac5 || fac7) {
            string rain = "";
            if (fac3) {
                rain += "Pling";
            }
            if (fac5) {
                rain += "Plang";
            }
            if (fac7) {
                rain += "Plong";
            }
            return rain;
        }
        return number.ToString ();
    }
}
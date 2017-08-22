using System;
using System.Collections.Generic;
using System.Linq;

public static class SumOfMultiples {
    public static int To (IEnumerable<int> multiples, int max) {
        List<int> mult = new List<int> ();
        foreach (var i in multiples) {
            mult.AddRange (getMultiples (i, max));
        }
        return mult.Distinct ().Sum ();
    }

    private static List<int> getMultiples (int num, int max) {
        List<int> mult = new List<int> ();
        for (int i = num; i < max; i += num) {
            mult.Add (i);
        }
        return mult;
    }
}
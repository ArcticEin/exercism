using System;
using System.Collections.Generic;

public class DNA {
    private Dictionary<char, int> dict;
    public DNA (string sequence) {
        dict = new Dictionary<char, int> { { 'A', 0 }, { 'T', 0 }, { 'C', 0 }, { 'G', 0 } };
        foreach (char c in sequence) {
            dict[c]++;
        }
    }

    public IDictionary<char, int> NucleotideCounts {
        get {
            return dict;
        }
    }

    public int Count (char nucleotide) {
        if (!"ACGT".Contains (nucleotide.ToString ())) {
            throw new InvalidNucleotideException ();
        }
        return dict[nucleotide];
    }
}

public class InvalidNucleotideException : Exception { }
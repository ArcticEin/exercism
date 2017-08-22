using System;
using System.Collections.Generic;
using System.Linq;

public static class AccumulateExtensions
{
    // Might be a too easy solution...
    public static IEnumerable<U> Accumulate<T, U>(this IEnumerable<T> collection, Func<T, U> func) => collection.Select(c => func(c));
}
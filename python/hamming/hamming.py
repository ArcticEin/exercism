def distance(s1, s2):
    if len(s1) != len(s2):
        raise ValueError
    diff = 0
    for idx, val in enumerate(s1):
        if val != s2[idx]:
            diff += 1
    return diff

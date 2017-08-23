def is_pangram(string):
    letters = []
    for char in string.lower():
        if char.isalpha() and letters.count(char) == 0:
            letters.append(char)
    if len(letters) == 26:
        return True
    return False

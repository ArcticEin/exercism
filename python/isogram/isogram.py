def is_isogram(string):
    letters = []
    for char in string.lower():
        if char.isalpha() and letters.count(char) > 0:
            return False
        letters.append(char)
    return True

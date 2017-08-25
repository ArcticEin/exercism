import re

def word_count(string):
    words = {}
    for word in re.sub('[\W+_]', ' ', string.lower()).split():
        if not word in words:
            words[word] = 1
        else:
            words[word] += 1
    return words

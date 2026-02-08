def encode(strs):
    result = ""
    for word in strs:
        result += str(len(word)) + "#" + word
    return result


def decode(str):
    result = []
    i = 0

    while i < len(str):
        # Start j where i is at.
        j = i
        while str[j] != "#":
            # Find the end of the word.
            j += 1
        # i is the start of the number and j is the index of the hashtag.
        length = int(str[i:j])
        # j + 1 would be the char after the hashtag.
        i = j + 1
        # now the i is the first char plus the length is the end of the world.
        j = i + length
        # Now slice from i:j to grab the word.
        result.append(str[i:j])
        # Start the i at the end of word so you can capture the next word.
        i = j
    return result

def valid_anagram(s, t):
    if len(s) != len(t):
        return False

    count = {}
    for char in s:
        count[char] = count.get(char, 0) + 1

    for char in t:
        # If char is not in the count of S string or if the count is already
        # zero for this char (too many instances of char in T compared to S).
        if char not in count or count[char] == 0:
            return False
        count[char] -= 1
    return True

from collections import defaultdict


def group_anagrams(strs):
    result = defaultdict(list)
    for s in strs:
        count = [0] * 26  # An array with zeros to represent a-z.
        for c in s:
            count[ord(c) - ord("a")] += 1
            print(count)
        result[tuple(count)].append(s)
    return list(result.values())

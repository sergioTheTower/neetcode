def isPalindrome(s):
    left_pointer, right_pointer = 0, len(s) - 1
    s = str.lower(s)
    while left_pointer < right_pointer:
        if not alphaNum(s[left_pointer]):
            left_pointer += 1
            continue
        if not alphaNum(s[right_pointer]):
            right_pointer -= 1
            continue
        if s[left_pointer] != s[right_pointer]:
            return False
        left_pointer += 1
        right_pointer -= 1
    return True


def alphaNum(c):
    return (
        ord("A") <= ord(c) <= ord("Z")
        or ord("a") <= ord(c) <= ord("z")
        or ord("0") <= ord(c) <= ord("9")
    )

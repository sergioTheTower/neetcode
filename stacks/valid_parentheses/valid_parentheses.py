def valid_paren(s):
    closeToOpen = {")": "(", "]": "[", "}": "{"}
    stack = []
    for char in s:
        # Is this paren a close paren and in the dict.
        if char in closeToOpen:
            # The stack is not empty and check if the last element
            # is a open paren.
            if stack and stack[-1] == closeToOpen[char]:
                # If match, remove the open bracket out of the stack.
                stack.pop()
            else:
                # this means a mismatch between the stacks.
                return False
        else:
            # The current paren is a open bracket so push to stack.
            stack.append(char)
    return True if not stack else False

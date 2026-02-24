def eval_rpn(tokens):
    stack = []
    for token in tokens:
        if token == "+":
            stack.append(stack.pop() + stack.pop())
        elif token == "-":
            a = stack.pop()
            b = stack.pop()
            stack.append(b - a)
        elif token == "*":
            a = stack.pop()
            b = stack.pop()
            stack.append(a * b)
        elif token == "/":
            a = stack.pop()
            b = stack.pop()
            stack.append(int(b / a))
        else:
            stack.append(int(token))
    return stack[0]

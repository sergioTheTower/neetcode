import pytest
from eval_polish_notation import eval_rpn


@pytest.mark.parametrize(
    "tokens, expected",
    [
        # Basic addition and multiplication: ((2 + 1) * 3) = 9
        (["2", "1", "+", "3", "*"], 9),
        # Division and addition: (4 + (13 / 5)) = 6
        (["4", "13", "5", "/", "+"], 6),
        # Complex expression: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
        (["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"], 22),
        # Truncation toward zero: (6 / -132) should be 0, not -1
        (["6", "-132", "/"], 0),
        # Simple subtraction
        (["10", "5", "-"], 5),
    ],
)
def test_eval_rpn(tokens, expected):
    assert eval_rpn(tokens) == expected


def test_single_token():
    """Test case for a single number with no operators."""
    assert eval_rpn(["42"]) == 42

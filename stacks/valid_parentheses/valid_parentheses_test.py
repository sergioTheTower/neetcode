import pytest
from valid_parentheses import valid_paren


# @pytest.mark.parametrize takes the variable names as a string,
# a list of tuples containing the test data, and an optional list of 'ids' for naming the tests in the output.
@pytest.mark.parametrize(
    "s, expected",
    [
        ("[]", True),
        ("([{}])", True),
        ("[(])", False),
        ("[", False),
        ("(((", False),
        ("]()", False),
    ],
    ids=[
        "simple_pair",
        "nested_and_sequential",
        "wrong_order",
        "single_bracket",
        "only_open_brackets",
        "starts_with_closing",
    ],
)
def test_is_valid(s, expected):
    assert valid_paren(s) == expected

import pytest
from valid_palindrome import isPalindrome


def test_simple_palindrome():
    assert isPalindrome("aba")


def test_complex_sentence():
    # "Was it a car or a cat I saw?" -> "wasitacaroracatisaw"
    input_str = "Was it a car or a cat I saw?"
    assert isPalindrome(input_str)


def test_not_palindrome():
    # "race a car" -> "raceacar" (not a palindrome)
    assert not isPalindrome("race a car")


def test_empty_string():
    # An empty string is a valid palindrome
    assert isPalindrome("")


def test_single_character():
    # A single character is always a palindrome
    assert isPalindrome("x")


def test_only_symbols():
    # "!!!" becomes "" which is a palindrome
    assert isPalindrome("!!!")


def test_mixed_case():
    # "Aa" -> "aa"
    assert isPalindrome("Aa")


def test_numbers():
    assert isPalindrome("12321")
    assert not isPalindrome("12345")


# Optional: Parameterized test (Best Practice in Pytest)
@pytest.mark.parametrize(
    "input_str, expected",
    [
        ("A man, a plan, a canal: Panama", True),
        ("race a car", False),
        (" ", True),
        ("0P", False),
    ],
)
def test_parametrized_cases(input_str, expected):
    assert isPalindrome(input_str) == expected

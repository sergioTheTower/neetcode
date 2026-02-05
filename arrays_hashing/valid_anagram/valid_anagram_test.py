import pytest
from valid_anagram import valid_anagram


@pytest.mark.parametrize(
    "s, t, expected",
    [
        ("racecar", "carrace", True),  # Example 1: Standard anagram
        ("jar", "jam", False),  # Example 2: Different characters
        ("rat", "car", False),  # Same length, different letters
        ("anagram", "nagaram", True),  # Classic test case
        ("a", "ab", False),  # Different lengths
        ("", "", True),  # Edge case: Empty strings
        ("a", "a", True),  # Edge case: Single character
        ("aabb", "bbaa", True),  # Multiple occurrences
        ("pcc", "pc", False),  # Subset length mismatch
    ],
)
def test_valid_anagram(s, t, expected):
    """
    Verifies valid_anagram identifies anagrams correctly
    across various string lengths and character distributions.
    """
    assert valid_anagram(s, t) == expected

import pytest
from group_anagrams import group_anagrams


def sort_result(result):
    """
    Helper to make the test order-agnostic.
    Sorts the strings inside sublists, then sorts the sublists themselves.
    """
    return sorted([sorted(group) for group in result])


@pytest.mark.parametrize(
    "input_strs, expected_output",
    [
        # Example 1: Standard case
        (
            ["act", "pots", "tops", "cat", "stop", "hat"],
            [["hat"], ["act", "cat"], ["stop", "pots", "tops"]],
        ),
        # Example 2: Single character
        (["x"], [["x"]]),
        # Example 3: Empty string
        ([""], [[""]]),
        # Extra Case: Multiple different anagram groups
        (
            ["eat", "tea", "tan", "ate", "nat", "bat"],
            [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
        ),
    ],
)
def test_group_anagrams_logic(input_strs, expected_output):
    # Call the function directly
    actual = group_anagrams(input_strs)
    assert sort_result(actual) == sort_result(expected_output)

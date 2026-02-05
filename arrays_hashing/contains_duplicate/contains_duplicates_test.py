import pytest
from contains_duplicates import contains_duplicates


@pytest.mark.parametrize(
    "nums, expected",
    [
        ([1, 2, 3, 3], True),  # Case 1: Simple duplicate
        ([1, 2, 3, 4], False),  # Case 2: No duplicates
        ([], False),  # Case 3: Empty list
        ([1, 1, 1, 1], True),  # Case 4: All same elements
        ([-1, -2, -1], True),  # Case 5: Negative duplicates
        ([1], False),  # Case 6: Single element
    ],
)
def test_contains_duplicates(nums, expected):
    """
    Test the contains_duplicates function with various scenarios
    using pytest parameterization for clean, dry code.
    """
    assert contains_duplicates(nums) == expected

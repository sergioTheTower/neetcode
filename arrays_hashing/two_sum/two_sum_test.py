import pytest
from two_sum import two_sum


@pytest.mark.parametrize(
    "nums, target, expected",
    [
        ([3, 4, 5, 6], 7, [0, 1]),  # Example 1
        ([4, 5, 6], 10, [0, 2]),  # Example 2
        ([5, 5], 10, [0, 1]),  # Example 3: Duplicates
        ([1, 2, 3, 4], 3, [0, 1]),  # Target at start
        ([1, 3, 4, 2], 6, [2, 3]),  # Target at end
        ([1000, 2000], 3000, [0, 1]),  # Large numbers
    ],
)
def test_two_sum(nums, target, expected):
    """
    Ensures two_sum returns the correct pair of indices
    in the required order (smaller index first).
    """
    result = two_sum(nums, target)
    assert result == expected

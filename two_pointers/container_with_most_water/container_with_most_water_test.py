import pytest
from container_with_most_water import contains_water


@pytest.mark.parametrize(
    "height, expected",
    [
        ([1, 7, 2, 5, 4, 7, 3, 6], 36),  # Example 1
        ([2, 2, 2], 4),  # Example 2
        ([1, 1], 1),  # Minimum constraints
        ([4, 3, 2, 1, 4], 16),  # Wide container is better
        ([1, 2, 1], 2),  # Middle dip
        ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),  # Complex case
        ([1, 1000], 1),  # Very tall but narrow
    ],
)
def test_max_area(height, expected):
    assert contains_water(height) == expected


def test_empty_input():
    """Edge case: Although constraints say length >= 2, good to handle graceful failure or 0."""
    assert contains_water([]) == 0
    assert contains_water([5]) == 0

import pytest
from car_fleet import carFleet


# Using parametrize to run multiple test cases through a single test function
@pytest.mark.parametrize(
    "target, position, speed, expected, test_name",
    [
        (10, [1, 4], [3, 2], 1, "Example 1"),
        (10, [4, 1, 0, 7], [2, 2, 1, 1], 3, "Example 2"),
        (10, [3], [3], 1, "Single Car"),
        (100, [0, 2, 4], [4, 2, 1], 1, "All become one fleet"),
        (10, [], [], 0, "Empty input edge case"),
    ],
)
def test_car_fleet(target, position, speed, expected, test_name):
    """
    Test the carFleet function against various scenarios.
    """
    result = carFleet(target, position, speed)
    assert result == expected, (
        f"Failed on {test_name}: expected {expected}, got {result}"
    )

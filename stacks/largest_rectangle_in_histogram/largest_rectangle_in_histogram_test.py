import pytest

# Assuming your function is named largestRectangleArea in the .py file
from largest_rectangle_in_histogram import largestRectangleArea


# Define all test cases in a clean, parameterized format
@pytest.mark.parametrize(
    "name, heights, expected",
    [
        ("example_1", [7, 1, 7, 2, 2, 4], 8),
        ("example_2", [1, 3, 7], 7),
        ("single_element", [5], 5),
        ("all_same_heights", [2, 2, 2, 2], 8),
        ("ascending_heights", [1, 2, 3, 4], 6),
        ("descending_heights", [4, 3, 2, 1], 6),
        ("empty_array", [], 0),
    ],
)
def test_largest_rectangle_area(name, heights, expected):
    """
    Tests the largestRectangleArea function with various edge cases.
    The 'name' parameter helps identify which specific test failed in the output.
    """
    assert largestRectangleArea(heights) == expected

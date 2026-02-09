import pytest

from product_array_except_self import product_except_self


@pytest.mark.parametrize(
    "name, nums, expected",
    [
        ("Example NC", [1, 2, 3, 4], [24, 12, 8, 6]),
        ("Example 1", [1, 2, 4, 6], [48, 24, 12, 8]),
        ("Example 2", [-1, 0, 1, 2, 3], [0, -6, 0, 0, 0]),
        ("Two Zeros", [1, 0, 3, 0], [0, 0, 0, 0]),
    ],
)
def test_product_except_self(name, nums, expected):
    """
    Tests the Product Except Self logic using parametrization
    to mimic the Go table-driven test style.
    """
    result = product_except_self(nums)

    # In Python, list comparison is deep by default
    assert result == expected, (
        f"Failed case '{name}': expected {expected}, got {result}"
    )

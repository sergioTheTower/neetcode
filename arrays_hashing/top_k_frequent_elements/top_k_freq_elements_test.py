import pytest
from top_k_freq_elements import top_k_freq_elements


@pytest.mark.parametrize(
    "nums, k, expected",
    [
        ([1, 1, 1, 2, 2, 3], 2, [1, 2]),
        ([7, 7], 1, [7]),
        ([1], 1, [1]),
        ([-1, -1], 1, [-1]),
        ([4, 1, -1, 2, -1, 2, 3], 2, [-1, 2]),
        ([1, 2], 2, [1, 2]),
    ],
)
def test_top_k_frequent(nums, k, expected):
    # Call the function directly
    actual = top_k_freq_elements(nums, k)

    # Using set comparison because the output order doesn't matter
    assert set(actual) == set(expected), (
        f"Failed on {nums}. Expected {expected}, got {actual}"
    )


def test_k_equals_length():
    nums = [1, 2, 3]
    k = 3
    actual = top_k_freq_elements(nums, k)
    assert set(actual) == {1, 2, 3}

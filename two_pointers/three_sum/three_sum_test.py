import pytest
from three_sum import three_sum


def normalize(triplets):
    """
    Helper to sort the results for comparison.
    1. Sorts each triplet internally: [-1, 2, -1] -> [-1, -1, 2]
    2. Sorts the list of triplets: [[0, 0, 0], [-1, -1, 2]] -> [[-1, -1, 2], [0, 0, 0]]
    """
    return sorted([sorted(t) for t in triplets])


@pytest.mark.parametrize(
    "nums, expected",
    [
        ([-1, 0, 1, 2, -1, -4], [[-1, -1, 2], [-1, 0, 1]]),
        ([0, 1, 1], []),
        ([0, 0, 0], [[0, 0, 0]]),
        ([], []),
    ],
)
def test_three_sum(nums, expected):
    result = three_sum(nums)
    assert normalize(result) == normalize(expected)

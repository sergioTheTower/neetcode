from longest_consecutive_sequence import longestConsecutive


def test_example_1():
    nums = [2, 20, 4, 10, 3, 4, 5]
    expected = 4
    assert longestConsecutive(nums) == expected


def test_example_2():
    nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
    expected = 9
    assert longestConsecutive(nums) == expected


def test_empty_array():
    nums = []
    expected = 0
    assert longestConsecutive(nums) == expected


def test_single_element():
    nums = [1]
    expected = 1
    assert longestConsecutive(nums) == expected


def test_no_consecutive():
    nums = [10, 30, 50]
    expected = 1
    assert longestConsecutive(nums) == expected


def test_negative_numbers():
    nums = [-1, -2, -3, 0, 1]
    expected = 5
    assert longestConsecutive(nums) == expected


def test_duplicates():
    nums = [1, 2, 0, 1]
    expected = 3
    assert longestConsecutive(nums) == expected


def test_large_gap():
    nums = [1, 2, 100, 101, 102]
    expected = 3
    assert longestConsecutive(nums) == expected

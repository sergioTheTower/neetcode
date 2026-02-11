import unittest
from two_sum import two_sum


# --- Test Suite ---
class TestTwoSum(unittest.TestCase):
    def test_example_1(self):
        # Input: numbers = [1,2,3,4], target = 3
        # Output: [1,2]
        self.assertEqual(two_sum([1, 2, 3, 4], 3), [1, 2])

    def test_example_2_non_adjacent(self):
        # Input: numbers = [2,3,4], target = 6
        # Output: [1,3] (2 + 4 = 6)
        self.assertEqual(two_sum([2, 3, 4], 6), [1, 3])

    def test_negative_numbers(self):
        # Input: numbers = [-1,0], target = -1
        # Output: [1,2] (-1 + 0 = -1)
        self.assertEqual(two_sum([-1, 0], -1), [1, 2])

    def test_mixed_signs(self):
        # Input: numbers = [-5, -2, 0, 3, 7], target = 2
        # Output: [1, 5] (-5 + 7 = 2)
        self.assertEqual(two_sum([-5, -2, 0, 3, 7], 2), [1, 5])

    def test_target_zero(self):
        # Input: numbers = [-3, 1, 2, 3], target = 0
        # Output: [1, 4] (-3 + 3 = 0)
        self.assertEqual(two_sum([-3, 1, 2, 3], 0), [1, 4])

    def test_large_gap(self):
        # Input: numbers = [1, 5, 10, 20, 100], target = 101
        # Output: [1, 5] (1 + 100 = 101)
        self.assertEqual(two_sum([1, 5, 10, 20, 100], 101), [1, 5])

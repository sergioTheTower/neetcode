from trap_rain_water import trap


class TestTrapRainWater:
    def test_example_1(self):
        """Standard example from problem description."""
        height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
        assert trap(height) == 6

    def test_example_2(self):
        """Another standard example."""
        height = [4, 2, 0, 3, 2, 5]
        assert trap(height) == 9

    def test_empty_list(self):
        """Test with an empty list."""
        assert trap([]) == 0

    def test_single_element(self):
        """Test with a single element (cannot trap water)."""
        assert trap([5]) == 0

    def test_two_elements(self):
        """Test with two elements (cannot trap water)."""
        assert trap([5, 10]) == 0
        assert trap([10, 5]) == 0

    def test_ascending(self):
        """Test with ascending heights (no water trapped)."""
        height = [1, 2, 3, 4, 5]
        assert trap(height) == 0

    def test_descending(self):
        """Test with descending heights (no water trapped)."""
        height = [5, 4, 3, 2, 1]
        assert trap(height) == 0

    def test_flat(self):
        """Test with flat surface (no water trapped)."""
        height = [3, 3, 3, 3]
        assert trap(height) == 0

    def test_bowl_shape(self):
        """Test a simple bowl shape."""
        height = [5, 0, 5]
        assert trap(height) == 5

    def test_w_shape(self):
        """Test a W shape which has two separate pockets of water."""
        height = [5, 1, 5, 1, 5]
        # Pocket 1: min(5,5) - 1 = 4
        # Pocket 2: min(5,5) - 1 = 4
        # Total = 8
        assert trap(height) == 8

    def test_pyramid(self):
        """Test a pyramid shape (no water trapped)."""
        height = [0, 1, 2, 3, 2, 1, 0]
        assert trap(height) == 0

    def test_large_numbers(self):
        """Test with larger integers."""
        height = [100, 0, 100]
        assert trap(height) == 100

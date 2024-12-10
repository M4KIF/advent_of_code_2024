from .puzzle_1 import Part1
import pytest

def test_runner_file_read_path_non_existent():
    with pytest.raises(Exception) as e_info:
        test_obj = Part1("data_non_existent.txt")
        assert test_obj.data == ""

def test_runner_file_read_path_existent_empty():
    with pytest.raises(Exception) as e_info:
        test_obj = Part1("test_data.txt")
        assert test_obj.data != ""

def test_runner_file_read_path_correct_data():
    test_obj = Part1("test_data.txt")
    assert len(test_obj.data) == 42

def test_solve_with_example_data():
    test = Part1("test_data.txt")
    assert test.solve() == 1928

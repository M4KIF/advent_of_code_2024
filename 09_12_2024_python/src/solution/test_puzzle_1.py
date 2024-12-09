from solution import puzzle_1
import pytest

def test_runner_file_read_path_non_existent():
    with pytest.raises(Exception) as e_info:
        test_obj = puzzle_1.Part1("../data_non_existent.txt")
        assert test_obj.data == ""

def test_runner_file_read_path_existent_empty():
    with pytest.raises(Exception) as e_info:
        test_obj = puzzle_1.Part1("../data_test.txt")
        assert test_obj.data == ""

def test_runner_file_read_path_existent():
    with pytest.raises(Exception) as e_info:
        test_obj = puzzle_1.Part1("../data.txt")
        assert test_obj.data == ""

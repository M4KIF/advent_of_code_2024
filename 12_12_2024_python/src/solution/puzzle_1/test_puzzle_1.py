from .puzzle_1 import Part1
import pytest

def test_puzzle_init_wrong_path():
    with pytest.raises(Exception) as e_info:
        test_obj = Part1("data_non_existent.txt")

def test_puzzle_init_correct_path_empty_file():
    test_obj = Part1("../solution/puzzle_1/test_data/data_1.txt")
    assert len(test_obj.area) == 0

def test_puzzle_init_correct_path_real_example_file():
    test_obj = Part1("../solution/puzzle_1/test_data/data_2.txt")
    assert len(test_obj.area) == 10
    assert len(test_obj.area[0]) == 10

def test_puzzle_flood_fill_attempt_one():
    # But only this one is working, 
    # despite spitting out the same result
    # for a small testcase. Harder edge cases ruin the
    # bottom one
    test_obj = Part1("../solution/puzzle_1/test_data/data_2.txt")

    visited = [len(test_obj.area[0]) * [0]] * len(test_obj.area)
    group = dict()
    stack = []
    stack.append((0,0))
    test_obj.find_group("R", group, stack, visited)
    assert len(group) == 12

def test_puzzle_flood_fill_attempt_two():

    test_obj = Part1("../solution/puzzle_1/test_data/data_2.txt")

    visited = [len(test_obj.area[0]) * [0]] * len(test_obj.area)
    group = dict()
    test_obj.flood_fill(0,0,"R", group, visited)
    assert len(group) == 12


def test_puzzle_1_solve_real_example_file():
    test_obj = Part1("../../data_example.txt")
    assert test_obj.solve() == 1930

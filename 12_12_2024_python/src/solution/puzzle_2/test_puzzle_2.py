from .puzzle_2 import Part2
import pytest

def test_puzzle_init_wrong_path():
    with pytest.raises(Exception) as e_info:
        test_obj = Part2("data_non_existent.txt")

def test_puzzle_init_correct_path_empty_file():
    test_obj_0 = Part2("test_data/data_1.txt")
    assert len(test_obj_0.area) == 0

def test_puzzle_group_normalise_points():
    # But only this one is working, 
    # despite spitting out the same result
    # for a small testcase. Harder edge cases ruin the
    # bottom one
    test_obj_1 = Part2("test_data/data_2.txt")

    visited = [len(test_obj_1.area[0]) * [0]] * len(test_obj_1.area)
    group = dict()
    stack = []
    stack.append((0,0))
    test_obj_1.find_group("R", group, stack, visited)
    group_obj = Part2.Group("R")
    
    for point in group:
        group_obj.add(point[0], point[1])

    grid = group_obj.normalize_points_to_array()

    assert len(grid) == 4
    assert len(grid[0]) == 5

def test_puzzle_boundary_walk():
    # But only this one is working, 
    # despite spitting out the same result
    # for a small testcase. Harder edge cases ruin the
    # bottom one
    test_obj_2 = Part2("test_data/data_2.txt")

    visited = [len(test_obj_2.area[0]) * [0]] * len(test_obj_2.area)
    group = dict()
    stack = []
    stack.append((0,0))
    test_obj_2.find_group("R", group, stack, visited)
    group_obj_1 = Part2.Group("R")
    
    for point in group:
        group_obj_1.add(point[0], point[1])

    sides = group_obj_1.boundary_walk()

    assert sides == 10

def test_puzzle_solve_with_boundary_scan():
    # But only this one is working, 
    # despite spitting out the same result
    # for a small testcase. Harder edge cases ruin the
    # bottom one
    test_obj_3 = Part2("test_data/data_3.txt")

    assert test_obj_3.solve() == 368

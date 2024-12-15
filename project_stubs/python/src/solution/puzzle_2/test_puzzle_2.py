from .puzzle_2 import Part2

import pytest

# def test_solve_with_test_data_1():
#     test = Part2("./test_data/test_data_1.txt")
#     print(test.data)
#     print(''.join(test.data))
#     assert test.solve() == 1267

# def test_solve_with_test_data_2():
#     test = Part2("./test_data/test_data_2.txt")
#     print(test.data)
#     print(''.join(test.data))
#     assert test.solve() == 1267

# def test_solve_with_test_data_3():
#     test = Part2("./test_data/test_data_3.txt")
#     print(test.data)
#     print(''.join(test.data))
#     assert test.solve() == 2858

def test_solve_with_test_data_4():
    test = Part2("./test_data/test_data_3.txt")
    print(test.data)
    print(''.join(test.data))
    assert test.alternative() == 2858

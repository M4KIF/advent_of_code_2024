import re

FILLER = "."

class Part2:

    # I've decided to keep the data as an array
    # Let's see how It will turn out
    reverse_string_data = ""
    string_data = ""
    data = []
    max_counter = 1
    encounters = dict()

    def __init__(self, path):
        with open(path, "r", encoding="utf-8") as f:
            # One line It is
            one_liner = f.readline()


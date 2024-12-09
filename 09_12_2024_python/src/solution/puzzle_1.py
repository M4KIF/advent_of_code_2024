
class Part1:

    data = str("")

    def __init__(self, path):
        with open(path, "r", encoding="utf-8") as f:
            print(f.readlines())
            # while f.read(1) is not None:
            #     print(f.read(1))

    def solve(self):
        print("None?")
        None
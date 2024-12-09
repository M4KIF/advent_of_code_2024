FILLER = "."

class Part1:

    # I've decided to keep the data as an array
    # Let's see how It will turn out
    data = []

    def __init__(self, path):
        with open(path, "r", encoding="utf-8") as f:
            # One line It is
            one_liner = f.readline()
            self.expand_data_list(one_liner)

    def expand_data_list(self, data_string):
        counter = 0
        data = []
        for i, char in enumerate(data_string):
            # On even, add data
            # On mean, increment counter and place "nils"
            if i % 2 == 0:
                for j in range(int(char)):
                    data.append(counter)
            else:
                counter+=1
                for j in range(int(char)):
                    data.append(FILLER)

        self.data = data
        
    def solve(self):

        reverse = self.data.copy()
        reverse.reverse()

        for i, char_reverse in enumerate(reverse):
            if char_reverse != FILLER:

                for j in range(0, len(self.data) - i):
                    if self.data[j] == FILLER:
                        self.data[j] = reverse[i]
                        self.data[len(self.data) - i - 1] = FILLER
                        break

        result = 0
        for i, char in enumerate(self.data):
            if char != FILLER:
                result += i * int(char)
            
        return result

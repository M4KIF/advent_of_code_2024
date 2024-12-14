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
            self.expand_data_list(one_liner)

    def expand_data_list(self, data_string):
        counter = 0
        data = []
        self.encounters.clear()
        for i, char in enumerate(data_string):
            # On even, add data
            # On mean, increment counter and place "nils"
            if i % 2 == 0:
                for j in range(int(char)):
                    data.append(str(counter))
                self.encounters[counter] = int(char)
                counter+=1
            else:
                for j in range(int(char)):
                    data.append(FILLER)

        self.max_counter = counter
        self.data = data

    def alternative(self):
        reverse = self.data.copy()
        reverse.reverse()

        counter = self.max_counter - 1 

        indexes_of_number = []
        indexes_of_dots = []
        for i, char_reverse in enumerate(reverse):
            if str(char_reverse) == str(counter):
                indexes_of_number.append(i)
                if len(indexes_of_number) == self.encounters[counter]:
                    for j in range(0, len(self.data) - i):
                        if self.data[j] == FILLER:
                            indexes_of_dots.append(j)
                            if len(indexes_of_dots) == len(indexes_of_number):
                                for k, num in enumerate(indexes_of_number):
                                    self.data[indexes_of_dots[k]] = reverse[num]
                                    self.data[len(self.data) - indexes_of_number[k] - 1] = "."

                                break
                        else:
                            indexes_of_dots.clear()
                    indexes_of_dots.clear()
                    indexes_of_number.clear()
                    counter-=1
            else:
                indexes_of_number.clear()

        print(self.data)

        result = 0
        for i, char in enumerate(self.data):
            if char != FILLER:
                result += i * int(char)
            
        return result
        
    # I want to fix this
    def mess(self):
        self.string_data = ''.join(self.data)

        self.index_to_exclude = len(self.string_data)
        for i in range(self.max_counter-1, -1, -1):
            temp_exclusion = 0
            # Search for the exclusion index
            for j in range(self.index_to_exclude - 1, 0, -1):
                #print(self.string_data[j:self.index_to_exclude])
                #print("Searching for: + " + f"{str(i) * self.encounters.get(i)}")
                index = self.string_data[j:self.index_to_exclude].find(f"{str(i) * self.encounters.get(i)}")
                if index < 0:
                    continue
                else:
                    temp_exclusion  = j
                    #print(self.string_data[j:temp_exclusion] + "\n")
                    break
            

            # self.string_data[j:self.index_to_exclude].replace(
            #         f"{str(i) * self.encounters.get(i)}",
            #         f"{"." * self.encounters.get(i)}"
            #         )
            # string_to_place = f"{str(i) * self.encounters.get(i)}"
            for j in range(0, temp_exclusion):
                index = self.string_data[:j].find(f"{"." * len(str(i)) * self.encounters.get(i)}")
                if index < 0:
                    continue
                else:
                    self.string_data =  self.string_data[:index] + self.string_data[index:j].replace(
                        f"{"." * len(str(i)) * self.encounters.get(i)}",
                        f"{str(i) * self.encounters.get(i)}",
                        1
                        ) + self.string_data[j:]
                    self.string_data = self.string_data[:temp_exclusion] + self.string_data[temp_exclusion:].replace(
                        f"{str(i) * self.encounters.get(i)}",
                        f"{"." * len(str(i)) * self.encounters.get(i)}",
                        1
                    )
                    break
            self.index_to_exclude = temp_exclusion
            #print(self.string_data)

            #print(f"{i} and {self.encounters.get(i)} index to which the freespace is sought {self.index_to_exclude}, czyli {self.string_data[self.index_to_exclude]}")
        #print(self.string_data)
        # reverse = self.data.copy()
        # reverse.reverse()

        # self.reverse_string_data = ''.join(reverse)
        # self.string_data = ''.join(self.data)

        # index = len(self.string_data)
        #print(self.encounters)

        # Descending order over file indexes
        # print(self.encounters)
        # print(self.string_data)
        # print("\n")
        # for i in range(self.max_counter-1, -1, -1):
        #     print(f"{i} and {self.encounters.get(i)} and index {index}")
            # print(f"{str(i) * self.encounters.get(i)}")
            # print(f"{"." * self.encounters.get(i)}")

             
            # if temp > 0:

            #     index = temp
            # else:
            #     continue
            # empty = (self.string_data[:index]).find(f"{"." * self.encounters.get(i)}")
            # print(f"{str(i) * self.encounters.get(i)}" + " " + f"{index} found num and {empty} dound empty")

            # if empty > 0:
            #     self.string_data = self.string_data.replace(f"{str(i) * self.encounters.get(i)}", f"{"." * self.encounters.get(i)}", 1)
            #     self.string_data = self.string_data.replace(f"{"." * self.encounters.get(i)}", f"{str(i) * self.encounters.get(i)}", 1)
            # print(self.string_data)
                # data_to_end = self.string_data
                # print(file[len(file)-1])
                # index = data_to_end.find(file[len(file)-1])

                # # Trying to place It from the start of the list
                # pattern = r"[.]{" + f"{len(file[0]) * len(file)}" +r"}"
                # dot_pattern = re.compile(pattern)
                # next_empty_space = dot_pattern.findall(self.string_data)

                # # It is needed to check whether the empty space is not above the file
                # if len(next_empty_space) > 0:
                #     string_data_copy = self.string_data
                #     string_data_copy = string_data_copy.replace(file[0], next_empty_space[0], 1)
                #     if (self.string_data.find(file[0]) - string_data_copy.find(next_empty_space[0])) >= i:
                #         self.string_data = self.string_data.replace(file[0], next_empty_space[0], 1)
                #         self.string_data = self.string_data.replace(next_empty_space[0], file[0], 1)
# def solve():
#     # regexes?
#     with open("../data.txt", "r") as f:

#         test = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'

#         # extract operations from data via regex
#         opeartionsPattern = re.compile(r'mul[(][0-9]{1,3},[0-9]{1,3}[)]')
#         operations = opeartionsPattern.findall(f.read())

#         # do operations via extracting numbers with a regex
#         num = re.compile(r'\d+')

#         to_multiply = [num.findall(o) for o in operations]
#         res = 0
#         for entry in to_multiply:
#             res += int(entry[0]) * int(entry[1])

#         print(res)

# solve()
            
        #     index = i
        #     length = 1
        #     for j, char in enumerate(reverse):
        #         if char != index:
        #             length = 1

        #         if char != FILLER:
        #             index = char

        #         if char == index:
        #             length+=1
        #     print("\n ")
                    

        # for i, char_reverse in enumerate(reverse):

        #     if char_reverse == FILLER:
        #         continue
            
        #     if index != char_reverse:
        #         # Do the work

        #         # for j in range(0, len(self.data) - i):
        #         #     if self.data[j] == FILLER:
        #         #         self.data[j] = reverse[i]
        #         #         self.data[len(self.data) - i - 1] = FILLER
        #         #         break
        #         length = 1
        #         index = char_reverse
        #     length+=1

        result = 0
        for i, char in enumerate(self.string_data):
            if char != FILLER:
                result += i * int(char)
            
        return result

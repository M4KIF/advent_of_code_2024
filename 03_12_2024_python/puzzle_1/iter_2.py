import re

def solve():
    # regexes?
    with open("../data.txt", "r") as f:
        test = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'
        to_multiply = [
            re.findall(r'\d+', o)
            for o in re.findall(r'mul[(][0-9]{1,3},[0-9]{1,3}[)]', f.read())
            ]
        print(sum([int(entry[0]) * int(entry[1]) for entry in to_multiply]))

solve()
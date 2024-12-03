import re

def solve():
    # regexes?
    with open("../data.txt", "r") as f:

        test = 'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'

        # extract operations from data via regex
        opeartionsPattern = re.compile(r'mul[(][0-9]{1,3},[0-9]{1,3}[)]')
        operations = opeartionsPattern.findall(f.read())

        # do operations via extracting numbers with a regex
        num = re.compile(r'\d+')

        to_multiply = [num.findall(o) for o in operations]
        res = 0
        for entry in to_multiply:
            print(f'%s multplied with %s\n', entry[0], entry[0])
            res += int(entry[0]) * int(entry[1])

        print(res)

solve()
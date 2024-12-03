import re

def solve():
    # regexes?
    with open("../data.txt", "r") as f:

        test = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

        # needed regexes
        mulRegex = "mul[(][0-9]{1,3},[0-9]{1,3}[)]"
        doRegex = "do[(][)]"
        dontRegex = "don't[(][)]"
        numRegex = "\d+"

        # patterns
        opeartionsPattern = re.compile(f"{mulRegex}|{doRegex}|{dontRegex}")
        mulPattern = re.compile(mulRegex)
        doPattern = re.compile(doRegex)
        dontPattern = re.compile(dontRegex)
        numPattern = re.compile(numRegex)

        # multiplication lock
        lock = False

        operations = opeartionsPattern.findall(f.read())

        res = 0
        for o in operations:
            if mulPattern.match(o) and not lock:
                nums = numPattern.findall(o)
                res += int(nums[0]) * int(nums[1])
            elif doPattern.match(o):
                lock = False
            elif dontPattern.match(o):
                lock = True

        print(res)

solve()
import re

UP = 0
RIGHT = 1
DOWN = 2
LEFT = 3

BLOCKING = "#"
CROSS = "X"

def find_starting_point(area: list[list[str]]):
    for y in range(0, len(area)):
        if area[y].count("^") != 1:
            continue
        else:
            for x in range(0, len(area[y])):
                if area[y][x] == "^":
                    return (y, x)

def is_in_bounds(x, y, m_x, m_y) -> bool:
    return (x >= 0 and x < m_x) and (y >= 0 and y < m_y)

def is_next_move_blocking(c_y, c_x, m_y, m_x, area) -> bool:
    return area[c_y+m_y][c_x+m_x] == BLOCKING

def main():
    # regexes?
    area = []

    with open("../data.txt", "r") as f:
        for line in f.readlines():
            newline_split = line.split("\n")
            area.append(re.findall("[.#^]{1}", newline_split[0]))

    # Possible movements of the guard
    dir_y = [-1, 0, 1, 0]
    dir_x = [0, 1, 0, -1]

    # Presence array
    visited = area.copy()

    y, x = find_starting_point(area)
    current_dir = UP

    while(is_in_bounds(x, y, len(area[0]), len(area))):
        visited[y][x] = CROSS

        if not is_in_bounds(x + dir_x[current_dir], y + dir_y[current_dir], len(area[0]), len(area)):
            break

        while is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], area):
            current_dir = (current_dir + 1) % 4

        x += dir_x[current_dir]
        y += dir_y[current_dir]

    count = 0
    for line in visited:
        count += line.count(CROSS)

    print(count)

    with open("result.txt", "w") as f:
        for line in visited:
            for char in line:
                f.write(char)
            f.write("\n")


if __name__ == "__main__":
    main()
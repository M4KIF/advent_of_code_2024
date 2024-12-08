import re
import math
import time

UP = 0
RIGHT = 1
DOWN = 2
LEFT = 3

BLOCKING = "#"
CROSS = "X"
NOT = "!"
NEW_OBSTRUCTION = "&"

def find_starting_point(area: list[list[str]]):
    for y in range(0, len(area)):
        if area[y].count("^") != 1:
            continue
        else:
            for x in range(0, len(area[y])):
                if area[y][x] == "^":
                    return (y, x)

def in_bounds(x, y, m_x, m_y) -> bool:
    return (x >= 0 and x < m_x) and (y >= 0 and y < m_y)

def is_next_move_blocking(c_y, c_x, m_y, m_x, area) -> bool:
    return area[c_y+m_y][c_x+m_x] == BLOCKING

def test_obstruction(x_t, y_t, c_t, dir_x, dir_y, visited, place_hodler_x, place_hodler_y, rany):

    if visited[place_hodler_y][place_hodler_x] == CROSS:
        visited[place_hodler_y][place_hodler_x] = BLOCKING
    else:
        return rany

   
    occurence = dict()
    ciekawe = 0

    while(in_bounds(x_t + dir_x[c_t], y_t + dir_y[c_t], len(visited[0]), len(visited))):
        while is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
            c_t = (c_t + 1) % 4
            continue

        x_t += dir_x[c_t]
        y_t += dir_y[c_t]

        value = occurence.get((x_t, y_t, c_t))
        if value is not None:
            occurence[(x_t, y_t, c_t)] +=1
            if occurence[(x_t, y_t, c_t)] >= 1:
                rany[(place_hodler_y, place_hodler_x, c_t)] = 1
                visited[place_hodler_y][place_hodler_x] = NEW_OBSTRUCTION
                return rany
        else:
            occurence[(x_t, y_t, c_t)] = 0


    visited[place_hodler_y][place_hodler_x] = CROSS
    return rany

def fill_occurence_table(x, y, dir_x, dir_y, area) -> list[list[str]]:
    # Meeting the first requirement to not alter the start position
    visited = area.copy()
    current_dir = UP
    visited[y][x] = NOT    

    while(in_bounds(x, y, len(area[0]), len(area))):
        if visited[y][x] != "!": 
            visited[y][x] = CROSS

        if not in_bounds(x + dir_x[current_dir], y + dir_y[current_dir], len(area[0]), len(area)):
            break

        if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
            current_dir = (current_dir + 1) % 4
            continue

        x += dir_x[current_dir]
        y += dir_y[current_dir]
    
    return visited

def find_loops(start_x, start_y, x, y, dir_x, dir_y, area) -> list[list[str]]:
    visited = area.copy()

    current_dir = UP

    rany = dict()
    counter = 0

    while(in_bounds(x + dir_x[current_dir], y + dir_y[current_dir], len(area[0]), len(area))):
        while is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
            current_dir = (current_dir + 1) % 4
            continue
        
        test_obstruction(start_x, start_y, UP, dir_x, dir_y, visited, x + dir_x[current_dir], y + dir_y[current_dir], rany)

        x += dir_x[current_dir]
        y += dir_y[current_dir]
        counter +=1
    
    return rany, visited

def main():
    # Part 1
    area = []

    with open("../data.txt", "r") as f:
        for line in f.readlines():
            newline_split = line.split("\n")
            area.append(re.findall("[.#^]{1}", newline_split[0]))

    # Possible movements of the guard
    dir_y = [-1, 0, 1, 0]
    dir_x = [0, 1, 0, -1]

    start = time.time()

    y, x = find_starting_point(area)
    visited_to_check = fill_occurence_table(x, y, dir_x, dir_y, area)
    res, visited = find_loops(x, y, x, y, dir_x, dir_y, visited_to_check)

    end = time.time()
    print(f"Part 1 completion time: {end - start} and result {len(res)}")

    with open("result1.txt", "w") as f:
        for line in visited:
            for char in line:
                f.write(char)
            f.write("\n")

    # Part 2
    area1 = []

    with open("../data_long.txt", "r") as f:
        for line in f.readlines():
            newline_split = line.split("\n")
            area1.append(re.findall("[.#^]{1}", newline_split[0]))

    start = time.time()

    y1, x1 = find_starting_point(area1)
    visited_to_check1 = fill_occurence_table(x1, y1, dir_x, dir_y, area1)

    res1, visited1 = find_loops(x1, y1, x1, y1, dir_x, dir_y, visited_to_check1)

    end = time.time()
    print(f"Part 2 completion time: {end - start} and result: {len(res1)}")

    with open("result1.txt", "w") as f:
        for line in visited1:
            for char in line:
                f.write(char)
            f.write("\n")
            

if __name__ == "__main__":
    main()
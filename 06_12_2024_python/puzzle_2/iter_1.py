import re
import math

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

def turn(y, x, last_dir, current_dir, dir_y, dir_x, visited) -> int:
    for i in range(0, len(dir_x)):
        current_dir = (current_dir + i) % 4
        if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
            continue
        else:
            # Dir switches that are not allowed
            # if (last_dir == UP and current_dir == DOWN or 
            #     last_dir == DOWN and current_dir == UP or 
            #     last_dir == RIGHT and current_dir == LEFT or 
            #     last_dir == LEFT and current_dir == RIGHT or
            #     last_dir == UP and current_dir == LEFT or 
            #     last_dir == RIGHT and current_dir == UP or
            #     last_dir == DOWN and current_dir == RIGHT or
            #     last_dir == LEFT and current_dir == DOWN):
            #     print("wow")
            #     return None
            
            return current_dir

    return None

def test_obstruction(x_t, y_t, x_start, y_start, c_t, dir_x, dir_y, area_p, visited, place_hodler_x, place_hodler_y, rany):

    if visited[place_hodler_y][place_hodler_x] == CROSS:
        visited[place_hodler_y][place_hodler_x] = BLOCKING
        # print(f"Try to block at {place_hodler_y} and {place_hodler_x}")
        # for line in visited:
        #     print(line)
        # print("\n") 
    else:
        return rany

    count = 0
    last_dir = c_t
    single_vertice_loop = 0
    counter = 0
    occurence = dict()

    while(in_bounds(x_t + dir_x[c_t], y_t + dir_y[c_t], len(visited[0]), len(visited))):
        if is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
            c_t = (c_t + 1) % 4
            continue
        # while is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
        #     c_t = (c_t + 1) % 4
        # for i in range(0, len(dir_x)):
        #     c_t = (c_t + i) % 4
        #     if is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
        #         continue
        #     else: 
        #         break

        # if is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
        #     visited[place_hodler_y][place_hodler_x] = CROSS
        #     return False
    
            

        # n_c_t = turn(y_t, x_t, c_t, c_t, dir_y, dir_x, visited)
        # if n_c_t is not None:
        #     c_t = n_c_t
        # else:
        #     visited[place_hodler_y][place_hodler_x] = CROSS
        #     return rany



        # Test change soon
        # if is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
        #     c_t = (c_t + 1) % 4

        # if last_dir == UP and c_t == DOWN or last_dir == DOWN and c_t == UP and last_dir == RIGHT and c_t == LEFT and last_dir == LEFT and c_t == RIGHT:
        #     single_vertice_loop +=1
        #     #print(f"Loop detected {last_dir}, {c_t}, {y_t}, {x_t}")
        # x_t += dir_x[c_t]
        # y_t += dir_y[c_t]

        # Went out of the paths
        if visited[y_t][x_t] == ".":
            #print("Banana")
            visited[place_hodler_y][place_hodler_x] = CROSS
            return rany

        # if x_t == x_start and y_t == y_start:
        #     count +=1

        # if count > 1 :
        #     visited[place_hodler_y][place_hodler_x] = NEW_OBSTRUCTION
        #     return True

        # if single_vertice_loop != 0:
        #     visited[place_hodler_y][place_hodler_x] = CROSS
        #     return False
        #print(f"Ohh, no! {x_t}, {y_t}, {y_start}, {x_start}")
        value = occurence.get((x_t, y_t))
        if value is not None:
            occurence[(x_t, y_t)] +=1
            if occurence.get((x_t, y_t)) > 4:
                rany[(place_hodler_y, place_hodler_x)] = 1
                visited[place_hodler_y][place_hodler_x] = NEW_OBSTRUCTION
                return rany
        else:
            occurence[(x_t, y_t)] = 1
        
        x_t += dir_x[c_t]
        y_t += dir_y[c_t]

        # if occurence.get((x_t, y_t)) > 4:
        #         # new_block = rany.get((place_hodler_y, place_hodler_x))
        #         # if new_block is None:
        #         rany[(place_hodler_y, place_hodler_x)] = 1
        #         #rany[(place_hodler_y, place_hodler_x)] = 1
        #         visited[place_hodler_y][place_hodler_x] = NEW_OBSTRUCTION
        #         return rany
        # for key in occurence:
        #     if occurence.get(key) >= 20:
        #         new_block = rany.get((y_t, x_t))
        #         if new_block is None:
        #             rany[(place_hodler_y, place_hodler_x)] = 1
        #         visited[place_hodler_y][place_hodler_x] = NEW_OBSTRUCTION
        #         return rany



    visited[place_hodler_y][place_hodler_x] = CROSS
    #print("Nope")
    return rany

def fill_occurence_table(x, y, dir_x, dir_y, area) -> list[list[str]]:
    # Meeting the first requirement to not alter the start position
    visited = area.copy()

    visited[y][x] = NOT
    #visited[y+dir_y[UP]][x + dir_x[UP]] = NOT

    current_dir = UP

    while(in_bounds(x, y, len(area[0]), len(area))):
        if visited[y][x] != "!": 
            visited[y][x] = CROSS

        if not in_bounds(x + dir_x[current_dir], y + dir_y[current_dir], len(area[0]), len(area)):
            break

        # while is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], area):
        #     current_dir = (current_dir + 1) % 4
        # for i in range(0, len(dir_x)):
        #     current_dir = (current_dir + i) % 4
        #     if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
        #         continue
        #     else: 
        #         break

        # if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
        #     return None

        if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
            current_dir = (current_dir + 1) % 4
            continue

        x += dir_x[current_dir]
        y += dir_y[current_dir]
    
    return visited

def find_loops(start_x, start_y, x, y, dir_x, dir_y, area) -> list[list[str]]:
    # Meeting the first requirement to not alter the start position
    visited = area.copy()

    current_dir = UP

    rany = dict()

    while(in_bounds(x, y, len(area[0]), len(area))):
        # if not in_bounds(x + dir_x[current_dir], y + dir_y[current_dir], len(area[0]), len(area)):
        #     break

        # # while is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], area):
        # #     current_dir = (current_dir + 1) % 4
        # #     last_corner_x = x
        # #     last_corner_y = y

        # # for i in range(0, len(dir_x)):
        # #     current_dir = (current_dir + i) % 4
        # #     if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
        # #         continue
        # #     else:
        # #         last_corner_x = x
        # #         last_corner_y = y
        # #         break

        # # if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
        # #     return None
        # n_c_t = turn(y, x, current_dir, current_dir, dir_y, dir_x, visited)
        # if n_c_t is not None:
        #     current_dir = n_c_t
        #     last_corner_x = x
        #     last_corner_y = y
        # else:
        #     return None

        # Next moves
        if not in_bounds(x + dir_x[current_dir], y + dir_y[current_dir], len(visited[0]), len(visited)):
            #print("Apple")
            break
        if is_next_move_blocking(y,x,dir_y[current_dir],dir_x[current_dir], visited):
            current_dir = (current_dir + 1) % 4
            continue
            

        # n_c_t = turn(y_t, x_t, c_t, c_t, dir_y, dir_x, visited)
        # if n_c_t is not None:
        #     c_t = n_c_t
        # else:
        #     visited[place_hodler_y][place_hodler_x] = CROSS
        #     return rany



        # Test change soon
        # if is_next_move_blocking(y_t,x_t,dir_y[c_t],dir_x[c_t], visited):
        #     c_t = (c_t + 1) % 4

        # if last_dir == UP and c_t == DOWN or last_dir == DOWN and c_t == UP and last_dir == RIGHT and c_t == LEFT and last_dir == LEFT and c_t == RIGHT:
        #     single_vertice_loop +=1
        #     #print(f"Loop detected {last_dir}, {c_t}, {y_t}, {x_t}")



        # # Went out of the paths
        # if visited[y_t][x_t] == ".":
        #     #print("Banana")
        #     visited[place_hodler_y][place_hodler_x] = CROSS
        #     return rany

        
        # try to place an obstruction
        # Test: Place in front of the current dir, turn right and hope for loop
        test_obstruction(x, y, start_x, start_y, current_dir, dir_x, dir_y, area, visited, x + dir_x[current_dir], y + dir_y[current_dir], rany)
        # counter+=1
        # print(counter)

        # x += dir_x[current_dir]
        # y += dir_y[current_dir]
        x += dir_x[current_dir]
        y += dir_y[current_dir]

    print(len(rany))
    
    return visited

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

    y, x = find_starting_point(area)
    visited_to_check = fill_occurence_table(x, y, dir_x, dir_y, area)


    looped = find_loops(x, y, x, y, dir_x, dir_y, visited_to_check)

    with open("result1.txt", "w") as f:
        for line in looped:
            for char in line:
                f.write(char)
            f.write("\n")


if __name__ == "__main__":
    main()
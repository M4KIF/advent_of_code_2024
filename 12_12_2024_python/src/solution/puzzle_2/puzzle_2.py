
import math
import sys
import numpy

RIGHT = 0
DOWN = 1
LEFT = 2
UP = 3

MOVEMENT = {(0, 1), (1, 0), (0, -1), (-1, 0)}
#MOVEMENT = list({(0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1,-1), (-1, 0), (-1, 1)})

class Part2:

    # I've decided to keep the data as an array
    # Let's see how It will turn out
    area = []
    plants = dict()

    class Group:

        plant = ""
        points = []
        grid = []
        def __init__(self, plant):
            #print(f"Creating a new group of {plant}")
            self.plant = plant
            self.points = []


        def __adjacent(self, y1, x1, y2, x2) -> bool:
            delta_x = int(math.fabs(x1 - x2))
            delta_y = int(math.fabs(y1 - y2))

            debug = ((delta_y == 1 and delta_x == 0) or
                (delta_x == 1 and delta_y == 0))

            #print(f"old point ({y1}, {x1}), new ({y2}, {x2}), result {debug}")

            # Only  
            if debug:
                    return True
            return False


        def can_be_added(self, new_point) -> bool:
            # print("\n")
            if new_point[0] != self.plant:
                return False
            
            for point in self.points:
                if self.__adjacent(point[0], point[1],
                new_point[1], new_point[2]):
                    return True
            return False


        def add(self, y, x) -> bool:
            adjancencies = 0

            #print(f"Addding {y}, {x}")

            for i, point in enumerate(self.points):
                if self.__adjacent(point[0], point[1],
                y, x):
                    self.points[i] = (point[0], point[1], point[2] - 1)
                    adjancencies +=1
            self.points.append((y, x, 4 - adjancencies))


        def merge(self, points_to_merge) -> bool: 
            possible = False
            # Checking if the merge is possible
            for new_point in points_to_merge:
                for i, point in enumerate(self.points):
                    if self.__adjacent(point[0], point[1],
                    new_point[1], new_point[2]):
                        possible = True
                        break

            if not possible:
                return False

            for new_point in points_to_merge:
                self.add(new_point[0], new_point[1])

            points_to_merge.clear()
            return True


        def get_group_size(self) -> int:
            return len(self.points)


        def get_perimeter(self) -> int:
            perimeter = 0
            for point in self.points:
                perimeter += point[2]
            return perimeter

        def normalize_points_to_array(self) -> []:

            m_y = 0
            m_x = 0
            s_y = sys.maxsize
            s_x = sys.maxsize
            for point in self.points:
                if point[0] > m_y:
                    m_y = point[0]
                if point[1] > m_x:
                    m_x = point[1]
                if point[0] < s_y:
                    s_y = point[0]
                if point[1] < s_x:
                    s_x = point[1]
            
            #print(f" min ({s_y}, {s_x}) and max ({m_y}, {m_x})")

            #print(self.points)
            #self.grid = [(m_x - s_x + 2) * [False]] * (m_y - s_y + 2)

            self.grid = None
            self.grid = numpy.zeros(shape=((m_y - s_y + 1), (m_x - s_x + 1)))

            #print(self.points)
            #print(self.grid)
            for point in self.points:
                #print(f"Addyng ({point[0]}, {point[1]})")
                self.grid[point[0] - s_y][point[1] - s_x] = 1
                #print(self.grid)

            return (self.grid, s_y, s_x)

        def boundary_walk(self) -> int:
            grid, s_y, s_x = self.normalize_points_to_array()

            # # Iteration then ...
            # sides = 0

            # # Top - down scan
            # for y, row in enumerate(grid):
            #     top = 0
            #     bottom = 0
            #     for x, field in enumerate(row):

            #         if field == 1:
            #             # Different behaviour for top/bottom margins
            #             if y == 0 or y == len(grid) - 1:
            #                top += 1
            #             elif len(grid) > 1:
            #                 print(f"{y}")
            #                 if grid[y-1][x] == 0:
            #                     top += 1
            #                 else:
            #                     if top > 0:
            #                         top = 0
            #                         sides += 1
            #                 if grid[y+1][x] == 0:
            #                     bottom += 1
            #                 else:
            #                     if bottom > 0:
            #                         bottom = 0
            #                         sides += 1
            #         else:
            #              # Different behaviour for top/bottom margins
            #             if y == 0 or y == len(grid) - 1:
            #                if top > 0:
            #                     top = 0
            #                     sides += 1
                    
            #         if bottom > 0 and x == len(grid[0]) - 1:
            #             sides+=1
            #         if top > 0 and x == len(grid[0]) - 1:
            #             sides +=1

            # # Left - Right scan
            # for x in range(len(grid[0])):
            #     left = 0
            #     right = 0
            #     for y in range(len(grid)):
            #         if grid[y][x] == 1:
            #             # Different behaviour for top/bottom margins
            #             if x == 0 or x == len(grid[0]) - 1:
            #                left += 1
            #             elif len(grid[0]) > 1:
            #                 print(f"{x}")
            #                 if grid[y][x - 1] == 0:
            #                     left += 1
            #                 else:
            #                     if left > 0:
            #                         left = 0
            #                         sides += 1
            #                 if grid[y][x + 1] == 0:
            #                     right += 1
            #                 else:
            #                     if right> 0:
            #                         right = 0
            #                         sides += 1
            #         else:
            #              # Different behaviour for top/bottom margins
            #             if x == 0 or x == len(grid[0]) - 1:
            #                if left > 0:
            #                     left = 0
            #                     sides += 1
                    
            #         if left > 0 and y == len(grid) - 1:
            #             sides+=1
            #         if right > 0 and y == len(grid) - 1:
            #             sides +=1

            sides = 0

            # Top bounds
            for y, row in enumerate(grid):
                boundary = 0
                for x, point in enumerate(row):
                    if point == 1:
                        # Top
                        if (y - 1 < 0):
                            boundary += 1
                        else:                            
                            if grid[y - 1][x] == 0:
                                boundary += 1
                            else:
                                if boundary > 0:
                                    boundary = 0
                                    sides += 1
                                            
                    else: 
                        if boundary > 0:
                            sides += 1
                            boundary = 0

                    if x == len(grid[0]) - 1:
                        if boundary > 0:
                            sides += 1

                    print(grid) 
                    print(f"Top, point ({y}, {x}), check boundary {boundary} sides {sides}")

            # Bottom bounds
            for y, row in enumerate(grid):
                boundary = 0
                for x, point in enumerate(row):
                    # Check the corners when 1
                    if point == 1:
                        # Top
                        if (y + 1 > len(grid) - 1):
                            boundary += 1
                        else:                            
                            if grid[y + 1][x] == 0:
                                boundary += 1
                            else:
                                if boundary > 0:
                                    boundary = 0
                                    sides += 1
                                            
                    else: 
                        if boundary > 0:
                            sides += 1
                            boundary = 0

                    if x == len(grid[0]) - 1:
                        if boundary > 0:
                            sides += 1
                    
                    print(grid) 
                    print(f"Bottom, point ({y}, {x}), check boundary {boundary} sides {sides}")

                        # Bottom bounds
            for x in range(len(grid[0])):
                boundary = 0
                for y in range(len(grid)):
                    # Check the corners when 1
                    if grid[y][x] == 1:
                        # Top
                        if (x - 1 < 0):
                            boundary += 1
                        else:                 
                            if grid[y][x - 1] == 0:
                                boundary += 1 
                            else:
                                if boundary > 0:
                                    boundary = 0
                                    sides += 1
                                            
                    else: 
                        if boundary > 0:
                            sides += 1
                            boundary = 0

                    if y == len(grid) - 1:
                        if boundary > 0:
                            sides += 1

                    print(grid) 
                    print(f"Left, point ({y}, {x}), check boundary {boundary} sides {sides}")

            for x in range(len(grid[0])):
                boundary = 0
                for y in range(len(grid)):
                    # Check the corners when 1
                    if grid[y][x] == 1:
                        # Top
                        if (x + 1 > len(grid[0]) - 1):
                            boundary += 1
                        else:                 
                            if grid[y][x + 1] == 0:
                                boundary += 1
                            else:
                                if boundary > 0:
                                    boundary = 0
                                    sides += 1
                                            
                    else: 
                        if boundary > 0:
                            sides += 1
                            boundary = 0

                    if y == len(grid) - 1:
                        if boundary > 0:
                            sides += 1
                    
                    print(grid) 
                    print(f"Right, point ({y}, {x}), check boundary {boundary} sides {sides}")
                            

                        # Different for the main body of the array
                        

                    # if field == 1:
                    #     edge+=1
                    # else:
                    #     if edge > 0:
                    #         edge = 0
                    #         sides += 1
                    #     # if y == 0:
                    #     #     if edge > 0:
                    #     #         edge = 0
                    #     #         sides +=1
                    #     #     continue
                    #     # if y > 1:
                    #     #     if grid[y-1][x] == 0:
                    #     #         i f edge > 
                    #     # if y < len(grid) - 1:
                    #     #     if grid[y+1][x] == 0:
                    #     #         None
                    #     # if y == len(grid) - 1:
                    #     #     None

                    #     # if edge > 0:
                    #     #     sides+=1
                    #     #     edge = 0
                    # if edge > 0:
                    #     sides +=1
                        

                          
                    # if y == len(grid) - 1: 
                    #     None
                    # if field == 1: 
                    #     if grid[y-1][x] == 0:
                    #         # For the top check
                    #         edge_top += 1
                    #     elif edge_top > 0:
                    #         edge_top = 0
                    #         sides+=1

                    #     if grid[y+1][x] == 0 and field == 1:
                    #         # For the top check
                    #         edge_bottom += 1
                    #     elif edge_bottom > 0:
                    #         edge_bottom = 0
                    #         sides+=1


                    
                # if edge_top == len(grid[0]) - 1 or edge_bottom == len(grid[0]) - 1:
                #     sides += 1

            # # Right - Left scan
            # for x in range(grid[0]):
            #     edge = 0
            #     for y in range(grid):

            #         # For the boundary vertice
            #         if x == 0 or x == len(grid) - 1:
            #             if grid[y][x] == 1:
            #                 # For the top check
            #                 edge += 1
            #             else:
            #                 edge = 0
            #                 sides+=1
            #             continue
                    
            #         # Right edge
            #         if grid[y][x] == 1 and grid[y][x-1] == 0:
            #             sides+=1

            #         # Left edge
            #         if grid[y][x] == 1 and grid[y][x+1] == 0:
            #             sides+=1

                    
            #     if edge == len(grid) - 1:
            #         sides += 1

            return sides

            # grid, s_y, s_x = self.normalize_points_to_array()

            # boundary = dict()
            # boundary[(s_y, s_x)] = 1

            # # Trying to turn left\
            # move = (s_y, s_x)
            # direction = UP
            # next_move = (s_y, s_x)
            # h = 0
            # count = 0

            # def invalid(grid, point, boundary):
            #     return point[0] >= len(grid) or point[0] < 0 or point[1] >= len(grid[0]) or point[1] < 0 or grid[point[0]][point[1]] == 0 or boundary.get(point) == 1

            # while next_move[0] is not s_y and next_move[1] is not s_x:
            #     print(boundary)
            #     print("!")
            #     if invalid(grid, next_move, boundary):
            #         # make it valid
            #         for i, m in enumerate(MOVEMENT):
            #             next_move = (move[0] + m[0], move[1] + m[1])
            #             if not invalid(grid, next_move, boundary):
            #                 direction = i
            #                 print(f"NIKAKDA, direzion {direction}")
            #                 break
            #     else:
            #         print(f"No i jak to tak, dir {direction}")
            #         h+=1
            #         boundary[next_move] = 1
            #         move = next_move
            #         next_move = (next_move[0] - 1, next_move[1])
            #         direction = UP
            #     count+=1
            
            # return h


                
                # # Every direction has it's own characteristics
                # # But at everyone I should try to go left
                # if direction == RIGHT:
                #     None
                # elif direction == DOWN:
                #     None
                # elif direction == LEFT:
                #     None
                # elif direction == UP:
                #     None


                # print("Kup")
                # # If can't go left, steer to the right until You can move
                # if invalid(grid, next_move, boundary):
                #     i = 0
                #     for m in MOVEMENT:
                #         next_move = (move[0] + m[0], move[1] + m[1])
                #         if not invalid(grid, next_move, boundary):
                #             boundary[(next_move[0], next_move[1])] = 1
                #             # Attempt left once again
                #             next_move = (next_move[0] - 1, next_move[1])
                #         else:
                #             break

                # else:
                #     boundary[(next_move[0], next_move[1])] = 1
                #     next_move = (next_move[0] - 1, next_move[1])
                    
            # Not working             
            # # Normalize the points to local indexes + get the start point
            # grid, s_y, s_x = self.normalize_points_to_array()

            # def go_right(p):
            #     return (-p[0], p[1])

            # def go_left(p):
            #     return (p[0],-p[1]) 

            # boundary = dict()
            # boundary[(s_y, s_x)] = 1

            # next_step = (1, 0)
            # point = (s_y + next_step[0], s_x + next_step[1])

            # count = 0
            # while count < 100:
            #     print(f"Welkin ({point[0]}, {point[1]})")


            #     if point[0] > len(grid) or point[0] < 0 or point[1] > len(grid[0]) or point[1] < 0 or grid[point[0]][point[1]] == 0:
            #         point = (point[0] - next_step[0], point[1] - next_step[1]) 
            #         next_step = go_right(next_step)
            #         point = (point[0] + next_step[0], point[1] + next_step[1])
            #     else:
            #         boundary[(point[0], point[1])] = 1
            #         next_step = go_left(next_step)
            #         point = (point[0] + next_step[0], point[1] + next_step[1])
            #     count +=1

                    
            # return len(boundary)
        

    def __init__(self, path):
        self.area.clear()
        self.plants.clear()

        with open(path, "r", encoding="utf-8") as f:
            lines = f.readlines()
            
            for y, line in enumerate(lines):
                plants_line = []
                for x, char in enumerate(line.strip()):
                    if self.plants.get(char) is None:
                        list_of_plants = []
                        list_of_plants.append((y,x))
                        self.plants[char] = list_of_plants
                    else:
                        temp = self.plants.get(char)
                        temp.append((y,x))
                        self.plants[char] = temp
                    plants_line.append(char)
                self.area.append(plants_line)
            
                #self.area.append(list(line.strip()))
                

    # NO
    def dfs(self, plant, y, x, visited, group):

        print(f"({y}, {x})")
        # Ze base kejs
        if self.area[y][x] == plant:
            group[(y,x)] = 1

        for move in MOVEMENT:

            n_y = move[0]
            n_x = move[1]
            if n_y >= 0 and n_y < len(self.area) and n_x >= 0 and n_x < len(self.area[0]) and not visited[n_y][n_x]:
                visited[y][x] = True
                self.dfs(plant, n_y, n_x, visited, group)
                visited[y][x] = False
    
    # Flood fill (with pattern checking(?))
    def find_group(self, plant, points: dict, stack: list, visited):
        if len(stack) > 0:
            vertice = stack.pop()

            if self.area[vertice[0]][vertice[1]] == plant:
                points[(vertice[0], vertice[1])] = 1
                visited[vertice[0]][vertice[1]]+=1
            else:
                return

            for move in MOVEMENT:

                n_y = vertice[0] + move[0]
                n_x = vertice[1] + move[1]
                if (n_y >=0 and n_y < len(self.area) and n_x >= 0 and n_x < len(self.area[0])):
                    if self.area[n_y][n_x] == plant and points.get((n_y, n_x)) is None:
                        stack.append((n_y, n_x))
            self.find_group(plant, points, stack, visited)
        
    # This one was close
    def flood_fill(self, x, y, plant, points, visited):
        #print(f"({y}, {x})")

        if not (y >= 0 and y < len(self.area) and x >=0 and x < len(self.area[0])):
            return

        if points.get((y,x)) is not None:
            return

        if visited[y][x] > 6:
            return

        if self.area[y][x] == plant:
            #print("HASD")
            points[(y,x)] = 1
            visited[y][x]+=1
            self.flood_fill(x, y + 1, plant, points, visited)
            self.flood_fill(x + 1, y, plant, points, visited)
            self.flood_fill(x , y -1, plant, points, visited)
            self.flood_fill(x - 1, y, plant, points, visited)

        return

    def solve(self) -> int:

        #Search for a group
        groups = []
        found = dict()

        visited = [len(self.area[0]) * [0]] * len(self.area)
        group = dict()
        stack = []

        for y, row in enumerate(self.area):
            for x, plant in enumerate(row):
                if found.get((y,x)) is not None:
                    continue

                # Searching for a group via flooding the area
                visited = [len(self.area[0]) * [0]] * len(self.area)
                group.clear()
                stack.clear()
                stack.append((y,x))
                self.find_group(plant, group, stack, visited)

                #print(f"plant: {plant} with group {group}")

                # Appending to found and to the group
                new_group = self.Group(plant)

                for found_plant_point in group:
                    found[found_plant_point] = 1
                    new_group.add(found_plant_point[0], found_plant_point[1])
                
                # Adding a new group
                groups.append(new_group)

        res = 0


        
        for group_item in groups:
            #res += group_item.boundary_walk()
            res += group_item.boundary_walk() * group_item.get_group_size()
            #es += group_item.get_group_size() * group_item.get_perimeter()

        #print(len(groups))
 
        return res
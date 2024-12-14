
import math
import collections

class Part1:

    # I've decided to keep the data as an array
    # Let's see how It will turn out
    area = []
    plants = dict()
    movement = {(0, 1), (-1, 0), (0, -1), (1, 0)}

    class Group:

        plant = ""
        points = []
        def __init__(self, plant):
            print(f"Creating a new group of {plant}")
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

            print(f"Addding {y}, {x}")

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
            
            print(self.plants)
                #self.area.append(list(line.strip()))

    # NO
    def dfs(self, plant, y, x, visited, group):

        print(f"({y}, {x})")
        # Ze base kejs
        if self.area[y][x] == plant:
            group[(y,x)] = 1

        for move in self.movement:

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

            for move in self.movement:

                n_y = vertice[0] + move[0]
                n_x = vertice[1] + move[1]
                if (n_y >=0 and n_y < len(self.area) and n_x >= 0 and n_x < len(self.area[0])):
                    if self.area[n_y][n_x] == plant and points.get((n_y, n_x)) is None:
                        stack.append((n_y, n_x))
            self.find_group(plant, points, stack, visited)
        
    # This one was close
    def flood_fill(self, x, y, plant, points, visited):
        print(f"({y}, {x})")

        if not (y >= 0 and y < len(self.area) and x >=0 and x < len(self.area[0])):
            return

        if points.get((y,x)) is not None:
            return

        if visited[y][x] > 6:
            return

        if self.area[y][x] == plant:
            print("HASD")
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

                print(f"plant: {plant} with group {group}")

                # Appending to found and to the group
                new_group = self.Group(plant)

                for found_plant_point in group:
                    found[found_plant_point] = 1
                    new_group.add(found_plant_point[0], found_plant_point[1])
                
                # Adding a new group
                groups.append(new_group)

        res = 0


        
        for group_item in groups:
            res += group_item.get_group_size() * group_item.get_perimeter()

        print(len(groups))
 
        return res
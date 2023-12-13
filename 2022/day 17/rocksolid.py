from dataclasses import dataclass
@dataclass
class Stone:

    def __init__(self, shape):
        self.shape = shape

    def get_shape(self, pos: tuple[int,int]):
        return [(sx + pos[0], sy + pos[1]) for sx,sy in self.shape]
    
    def is_coliding(self, pos, hist):
        for shape in self.shape:
            x,y = pos[0] + shape[0], pos[1] + shape[1]
            if x < 0 or x >= len(hist[0]):
                return True
            if y < len(hist):
                if hist[y][x]:
                    return True

        return False 


hor_line = Stone(
    shape= [(0,0),(1,0),(2,0),(3,0)],
)

plus = Stone(
    shape=[(0,1),(1,1),(2,1),(1,0),(1,2)],
)

L = Stone(
    shape= [(0,0),(1,0),(2,0),(2,1),(2,2)],
)

ver_line = Stone(
    shape= [(0,0),(0,1),(0,2),(0,3)],
)

square = Stone(
    shape=[(0,0),(1,0),(1,1),(0,1)],
)

hoff = {
    "<": -1,
    ">": 1
}

stones = [hor_line,plus,L,ver_line,square]

def get_rock():
    while True:
        for s in stones:
            yield s

def create_direction_iter(directions: str):
    def dir_iter():
        while True:
            for d in directions:
                yield hoff[d]

    print("returning dir iter")
    return dir_iter

def to_str(hist, idx):

    char = {
        True: "#",
        False: "."
    }

    with open(f"out/out_{idx}","w") as f:
        for line in reversed(hist):
            f.write(f"|{''.join([char[v] for v in line])}| \n")
        f.write("+-------+\n")


def part1(directions):
    get_dir = create_direction_iter(directions)()

    i = 0
    
    hist = [[True]*7]
    for stone in get_rock():
        if i >= 2022:
            break

        x, y= 2, len(hist) + 3

        d = next(get_dir)
        if not stone.is_coliding((x + d, y),hist):
                x, y = x + d, y

        while not stone.is_coliding((x,y-1),hist):
            x, y = x, y - 1
            d = next(get_dir)
            if not stone.is_coliding((x + d, y),hist):
                x, y = x + d, y


        to_add = stone.get_shape((x,y))      
        for pos in to_add:
            while pos[1] >= len(hist):
                hist.append([False] * 7)
            hist[pos[1]][pos[0]] = True

        to_str(hist,i)
        i += 1
    return len(hist)



def main():
    with open("inp") as f:
        directions = f.read()

    res = part1(directions.strip())
    print(res)

# main()
n  = 1000000000000
acc = 0
for i in range(n):
    acc += i

    
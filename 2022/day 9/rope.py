import math
changes = {
    "U": (1, 0),
    "D": (-1, 0),
    "L": (0, -1),
    "R": (0, 1),
}
def dist(a,b):
    return max(abs(a[0] - b[0]), abs(a[1] - b[1]))

def main():
    with open("inp") as f:
        positions = set()
        rope = [(0,0) for _ in range(10)]
        positions.add(rope[9])
        for line in f.readlines():
            tmp = line.strip().split()
            dr, moves =  changes[tmp[0]], int(tmp[1])
            for _ in range(moves):
                rope[0] = rope[0][0] + dr[0], rope[0][1] + dr[1] 
                for i in range(1,len(rope)):
                    if dist(rope[i-1], rope[i]) > 1:
                        dx = rope[i-1][0] - rope[i][0]
                        dx = abs(dx)/dx if dx != 0 else 0                   
                        dy = rope[i-1][1] - rope[i][1]
                        dy = abs(dy)/dy if dy != 0 else 0
                        rope[i] = rope[i][0] + dx, rope[i][1] + dy
                positions.add(rope[9])
    return len(positions)

print(main())
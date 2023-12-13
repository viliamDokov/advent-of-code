from dataclasses import dataclass, field
from typing import Self

@dataclass
class Node:
    h: int = field(repr=False)
    x: int
    y: int
    neighbours: list[Self] = field(default_factory=list, repr=False) 
    prev: Self | None = field(default=None,repr=False)

    def __hash__(self): return hash((self.x, self.y))
    def __eq__(self, other):
        if type(self) == type(other):
            return self.x == other.x and self.y == other.y
        return False 

def add_n(a: Node, b:Node):
    if a.h + 1 >= b.h:
        a.neighbours.append(b)


def main():
    with open("inp") as f:
        el_map: list[list[Node]] = []
        x = 0
        starts = []
        for line in f.readlines():
            tmp = []
            y = 0
            for ch in line.strip():
                if ch == "S":
                    node = Node(h=ord("a"), x=x, y=y)
                    tmp.append(node)
                    starts.append(node)
                elif ch == "a":
                    node = Node(h=ord("a"), x=x, y=y)
                    tmp.append(node)
                    starts.append(node)
                elif ch == "E":
                    end = Node(h=ord("z"), x=x, y=y)
                    tmp.append(end)
                else:
                    tmp.append(Node(h=ord(ch), x=x, y=y))
                y += 1
            el_map.append(tmp)
            x += 1

        nx = len(el_map)
        ny = len(el_map[0])
        for i in range(nx):
            for j in range(ny):
                curr = el_map[i][j]
                if i - 1 >= 0: add_n(curr, el_map[i-1][j])
                if i + 1 < nx: add_n(curr, el_map[i+1][j])
                if j - 1 >= 0: add_n(curr, el_map[i][j - 1])
                if j + 1 < ny: add_n(curr, el_map[i][j + 1])
                # print(curr,curr.h, curr.neighbours)

        best = 1000000000
        for start in starts:
            count = True 
            seen = set()
            seen.add(start)
            q = [start]
            while curr != end:
                if len(q) == 0:
                    count = False
                    break
                curr = q.pop(0)
                for nb in curr.neighbours:
                    if nb not in seen:
                        nb.prev = curr
                        q.append(nb)
                        seen.add(nb)
            if count:
                acc = 0
                while curr != start:
                    curr = curr.prev
                    acc += 1
                    
                if acc < best:
                    best = acc

        print(best)
                

main()

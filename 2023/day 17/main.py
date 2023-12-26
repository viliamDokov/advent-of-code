from queue import PriorityQueue
from collections import defaultdict
from enum import IntEnum
from math import inf

dist = defaultdict(lambda: defaultdict(lambda: defaultdict(lambda: [inf for _ in range(11)])))
queue = PriorityQueue()


class Direction(IntEnum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3


def main():
    with open('inp.main') as f:
        lines = f.read().strip().split('\n')
    lines = [[int(char) for char in line] for line in lines]
    queue.put((0, 0, 0, 0, Direction.RIGHT, []))
    for d in Direction:
        dist[0][0][d] = [0 for _ in range(len(dist[0][0][d]))]
    ans = inf
    while not queue.empty():
        curr_dist, straight, row, col, d, lis = queue.get()
        if row == len(lines) - 1 and col == len(lines[0]) - 1:
            if curr_dist < ans:
                ans = curr_dist
                pog_list = lis 
        positions = get_next_pos(curr_dist, straight, row, col, d, lines)
        for position in positions:
            next_dist, next_straight, next_row, next_col, next_d = position
            if next_dist >= dist[next_row][next_col][next_d][next_straight]:
                continue
            new_lis = lis.copy()
            new_lis.append((row, col, curr_dist, straight))
            dist[next_row][next_col][next_d][next_straight] = next_dist
            queue.put((*position, new_lis))
    print(pog_list)
    print(ans)


def get_next_pos(dist, straight, row, col, d, lines):
    next_pos = []
    for direction in Direction:
        if direction == opposite(d):
            # can't go back
            continue
        if direction == d:
            if straight == 10:
                continue
            try:
                next_row, next_col, next_dir, add = get_next_loc(row, col, direction, lines, 1)
            except IndexError:
                continue
            next_pos.append((dist + add, straight + 1, next_row, next_col, direction))
        else:
            try:
                next_row, next_col, next_dir, add = get_next_loc(row, col, direction, lines, 4)
            except IndexError:
                continue
            next_pos.append((dist + add, 4, next_row, next_col, direction))
    return next_pos


def opposite(dir):
    if dir == Direction.UP:
        return Direction.DOWN
    if dir == Direction.DOWN:
        return Direction.UP
    if dir == Direction.LEFT:
        return Direction.RIGHT
    if dir == Direction.RIGHT:
        return Direction.LEFT


def get_next_loc(row, col, direction, lines, step):
    add = 0
    if direction == Direction.DOWN:
        for i in range(row + 1, row + step + 1):
            add += lines[i][col]
        return (row + step, col, direction, add)
    elif direction == Direction.UP:
        if row - step < 0:
            raise IndexError
        for i in range(row - 1, row - step - 1, -1):
            add += lines[i][col]
        return (row - step, col, direction, add)
    elif direction == Direction.LEFT:
        if col - step < 0:
            raise IndexError
        for j in range(col - 1, col - step - 1, -1):
            add += lines[row][j]
        return (row, col - step, direction, add)
    elif direction == Direction.RIGHT:
        for j in range(col + 1, col + step + 1):
            add += lines[row][j]
        return (row, col + step, direction, add)


if __name__ == "__main__":
    main()
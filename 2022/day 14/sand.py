def get_next(pos, stones, flr):
    nxt = (pos[0], pos[1] + 1)
    if nxt[1] == flr:
        return pos
    if nxt not in stones:
        return nxt
    nxt = (pos[0] - 1, pos[1] + 1)
    if nxt not in stones:
        return nxt
    nxt = (pos[0] + 1, pos[1] + 1)
    if nxt not in stones:
        return nxt
    return pos

def spawn(stones: set[tuple[int,int]],flr):
    start = 500,0
    pos = start
    nxt = get_next(pos, stones,flr)
    while nxt != pos:
        pos = nxt
        nxt = get_next(pos, stones, flr)
    
    stones.add(pos)
    return pos, stones

def main():
    with open("inp") as f:
        stones: set[tuple[int,int]] = set()
        for line in f.readlines():
            segments = line.strip().split(" -> ")
            for i in range(1, len(segments)):
                prv = [int(t) for t in segments[i - 1].split(",")]
                cr = [int(t) for t in segments[i].split(",")]

                if prv[0] == cr[0]:
                    smaller, larger = (prv[1], cr[1]) if prv[1] < cr[1] else (cr[1], prv[1])
                    # print(smaller,larger)
                    for y in range(smaller, larger + 1):
                        tp = (prv[0],y)
                        stones.add(tp)

                elif prv[1] == cr[1]:
                    smaller, larger = (prv[0], cr[0]) if prv[0] < cr[0] else (cr[0], prv[0])
                    # print(smaller,larger)
                    for x in range(smaller, larger + 1):
                        tp = (x,prv[1])
                        stones.add(tp)
        oldl = len(stones)
        acc = 0
        flr = max(stones,key=lambda x: x[1])[1] + 2
        new, stones = spawn(stones,flr)
        while new != (500,0):
            new, stones = spawn(stones,flr)
            acc += 1
        print(len(stones) - oldl)



main()
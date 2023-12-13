from dataclasses import dataclass, field
from typing import Self
from time import perf_counter

@dataclass()
class Valve:
    idx: int
    name: str
    nxts: list[str] = field(hash=False)
    flow: int = field(hash=False)
    nxt: list[Self] = field(default_factory=list, hash=False, repr=False)

    def __eq__(self,other):
        return type(self) == type(other) and self.name == other.name

    def __hash__(self) -> int:
        return hash(self.name)
    def __repr__(self) -> str:
        return self.name

def set_dists(curr):
    q: list[Valve] = [curr]
    seen = set()
    seen.add(curr)
    while len(q) > 0:
        x = q.pop(0)
        for n in x.nxt:
            if n not in seen:
                n.dist = x.dist + 1
                q.append(n)
                seen.add(n)
def get_distance_matrix(vavles:list[Valve]):
    matrx = {}
    for valve in vavles:
        q = [valve]
        matrx[(valve,valve)] = 0 
        seen = set()
        seen.add(valve)
        while len(q) > 0:
            x = q.pop(0)
            for n in x.nxt:
                if n not in seen:
                    matrx[(valve,n)] = matrx[(valve,x)] + 1 
                    q.append(n)
                    seen.add(n)
    return matrx

def max_flow(valves: list[Valve], dists, time, start, seen):
    best = 0
    stack = [(time,start,start,time,time,0,seen)] 
    rf = sum([v.flow for v in valves])

    while len(stack) > 0:
        time, start_ME, start_ELE, time_ME, time_ELE, flow, seen = stack.pop()
        # print(time, start_ME, start_ELE, time_ME, time_ELE, flow, seen)

        if time_ME == time:
            # print("ME")
            for valve in valves:
                bit_val_m = 1 << valve.idx
                if (bit_val_m & seen) == 0:
                    new_time_m = time - 1 - dists[(start_ME,valve)] 
                    new_flow = flow + valve.flow * new_time_m
                    if new_flow > best:
                        best = new_flow

                    if flow + new_time_m * rf < best or time < 0:
                        # print("TRIM!")
                        continue
                    
                        
                    new_t = max(time_ELE,new_time_m)
                    if new_t > 0:
                        stack.append((
                            new_t,
                            valve,
                            start_ELE,
                            new_time_m,
                            time_ELE,
                            new_flow,
                            seen | bit_val_m
                        ))

        elif time_ELE == time:
            # print("ELE")
            for valve in valves:
                bit_val_e = 1 << valve.idx
                if (bit_val_e & seen) == 0:
                    new_time_e = time - 1 - dists[(start_ELE,valve)] 
                    new_flow = flow + valve.flow * new_time_e
                    
                    if new_flow > best:
                        best = new_flow

                    if flow + new_time_e * rf < best or time < 0:
                        # print("TRIM!")
                        continue

                    new_t = max(new_time_e,time_ME)
                    if new_t > 0:
                        stack.append((
                            new_t,
                            start_ME,
                            valve,
                            time_ME,
                            new_time_e,
                            new_flow,
                            seen | bit_val_e
                        ))

    return best

def main():
    valves: dict[str,Valve] = {}
    tf = 0
    with open("inp") as f:
        for i, line in enumerate(f.readlines()):
            tmp = line.strip().split(";")
            name = tmp[0][6:].split(" has flow rate=")[0]
            flow= int(tmp[0][6:].split(" has flow rate=")[1])
            tf += flow
            tmp2 = tmp[1][23:].split(", ")
            nbs = []
            for n in tmp2:
                nbs.append(n)
            v = Valve(i,name,nbs,flow)
            valves[name] = v


    aa = valves['AA']
    for n,v in valves.items():
        for ns in v.nxts:
            v.nxt.append(valves[ns])
        
    valves: list[Valve] = list(valves.values())
    dists = get_distance_matrix(valves)
    
    seen = 0
    for v in valves:
        if v.flow <= 0:
            seen |= 1 << v.idx

    valves = [v for v in valves if v.flow > 0]
    print(dists)
    t1 = perf_counter()
    r = max_flow(valves,dists,26,aa,seen)
    t2 = perf_counter()

    print(r, t2 -t1)

main()
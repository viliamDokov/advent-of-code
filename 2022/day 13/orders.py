from dataclasses import dataclass, field
from itertools import batched
from typing import Self
from functools import total_ordering
WRONG = -1
TIE = 0
RIGHT = 1

@total_ordering
class MyList:
    data: list[Self,int] 

    def __init__(self, x=None):
        if x == None:
            self.data = []
        else:
            self.data = [x]

    def __len__(self) -> int:
        return len(self.data)

    def __eq__(self, other):
        return validate(self, other) == TIE
    
    def __lt__(self, other):
        return validate(self, other) == RIGHT

    def __getitem__(self,key):
        return self.data[key]

    def append(self, x):
        self.data.append(x)
        
    def __repr__(self) -> str:
        return self.data.__repr__()

def parse_line(line:str):
    if line.strip().startswith("["):
        line = line[1:]
        obj = MyList()
        while True:
            if line.startswith("]"):
                break
            thing, line = parse_line(line)
            obj.append(thing)
            if line.startswith(","):
                line = line[1:]
        line = line[1:]
        return obj, line 
    else:
        acc = ""
        while not line.startswith(",") and not line.startswith("]"):
            acc += line[0]
            line = line[1:]
        return int(acc), line
    
def validate(left, right):
    if type(left) == int and type(right) == int:
        if left < right:
            return RIGHT
        if left == right:
            return TIE
        else:
             return WRONG
    elif type(left) == MyList and type(right) == MyList: 
        smallest = min(len(left), len(right))
        acc = 1
        for i in range(smallest):
            corr = validate(left[i], right[i])
            if corr != TIE:
                return corr
        if len(left) < len(right):
            return RIGHT
        if len(left) == len(right):
            return TIE
        else:
             return WRONG
    elif type(left) == int and type(right) == MyList: 
        return validate(MyList(left), right)   
    elif type(left) == MyList and type(right) == int: 
        return validate(left, MyList(right))   
        
def main():
    with open("inp") as f:
        
        a1 = MyList(MyList(2))
        a2 = MyList(MyList(6))
        packets = [a1, a2]
        for pair in batched(f.readlines(),3):
            p1, _ = parse_line(pair[0].strip())
            p2, _ = parse_line(pair[1].strip())
            # packets.append((p1,p2))
            packets.append(p1)
            packets.append(p2)

    # acc = 0
    # for i, pair in enumerate(packets, 1):
    #     p1, p2 = pair
    #     print(i, p1 < p2)
    #     if p1 < p2:
    #         acc += i  
    # print(acc)

    packets = sorted(packets)
    for packet in packets:
        print(packet)
    i1 = packets.index(a1) + 1
    i2 = packets.index(a2) + 1
    print(i1,i2, i1 * i2)



main()



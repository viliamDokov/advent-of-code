def prio(a):
    v = ord(a) - 96
    if v <= 0:
        v = 26 + ord(a) - 64
    return v

def main():
    with open("inp") as f:
        s = 0
        line = f.readline().strip()
        while line!= "":
            r1 = set(list(line))
            line = f.readline().strip()
            r2 = set(list(line))
            line = f.readline().strip()
            r3 = set(list(line))
            common = set(r1).intersection(set(r2)).intersection(r3)
            s += prio(common.pop())
            line = f.readline().strip()

    return s

print(main())
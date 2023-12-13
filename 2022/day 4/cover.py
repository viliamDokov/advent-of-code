def cover(a, b):
    return int(a[0] <= b[0] <= a[1] or a[0] <= b[1] <= a[1])

def main():
    with open("inp") as f:
        s = 0
        for line in f.readlines():
            p1, p2 = line.strip().split(",")
            p1 = int(p1.split("-")[0]), int(p1.split("-")[1])
            p2 = int(p2.split("-")[0]), int(p2.split("-")[1])
            print(p1, p2)
            if cover(p1,p2):
                s += 1
            else:
                s += cover(p2,p1)
    return s

print(main())
                


            
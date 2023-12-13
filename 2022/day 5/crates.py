def main():
    with open("inp") as f:
        line = f.readline()
        n = 9
        stacks = [[] for _ in range(n)]
        while line.startswith("["):
            for i in range(n):
                a = line[i * 4 + 1]
                print(line, i * 4 + 1, a)
                if a != " ":
                    stacks[i].insert(0,a)
            line = f.readline()
        f.readline()

        for line in f.readlines():
            vs = line.strip()[5:].replace(" from ", ",").replace(" to ", ",").split(",")
            c, s, t = int(vs[0]), int(vs[1]) - 1, int(vs[2]) - 1
            print(c,s,t)
            tower = []
            for _ in range(c):
                tower.append(stacks[s].pop())
            tower = tower[::-1]
            stacks[t] += tower
        r = ""
        for stack in stacks:
            r += stack[-1]
        return r
                
                

print(main())
                
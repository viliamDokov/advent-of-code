def main():
    with open("inp") as f:
        acc = 0
        reg = 1

        signals = [20,60,100,140,180,220]
        cycle = 1

        drawing  = [["."] * 40 for i in range(6)]

        def tick():
            nonlocal cycle
            nonlocal acc

            cycle += 1

            nonlocal drawing
            x, y = (cycle - 1) % 40, (cycle - 1) // 40
            print(x,y)
            if reg - 1 <= x and x <=  reg + 1:
                drawing[y][x] = "#"
                
            if len(signals) > 0 and signals[0] == cycle:
                acc += cycle * reg
                signals.pop(0)
        for line in f.readlines():
            if line.strip().startswith("addx"):
                a = int(line.strip().split()[1])
                tick()
                reg += a
                tick()
            else:
                tick()
        return drawing
            

dr = main()
with open("out", "w") as f:
    for row in dr:
        for char in row:
            f.write(char)
        f.write("\n")
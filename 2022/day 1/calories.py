def main():
    cals = []
    acc = 0
    with open("inp") as f:
        for line in f.readlines():
            v = line.strip() 
            print(v, v == "")
            if v == "":
                cals.append(acc)
                acc = 0
            else:
                acc += int(v)
    top3 = []
    for i in range(3):
        top3.append(max(cals))
        cals.remove(max(cals))

    print(cals)
    return sum(top3)

print(main())

        
with open("out") as f:
    data = set()
    for line in f.readlines():
        if line in data:
            print(line)
        data.add(line)
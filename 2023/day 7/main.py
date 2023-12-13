with open("inp") as f:
    data = f.read()

hands = []
for line in data.split("\n"):
    h, b = line.split(" ")
    hands.append((h, b))


def value(h: str):
    h = h.replace("A", "Z")
    h = h.replace("K", "Y")
    h = h.replace("Q", "X")
    h = h.replace("J", "!")
    h = h.replace("T", "V")

    counts = {}
    for c in h:
        counts[c] = counts.get(c, 0) + 1

    j = counts.get("!", 0)
    if "!" in counts:
        del counts["!"]
    freqs = [0] + list(sorted(counts.values()))

    if freqs[-1] + j >= 5:
        return 10, h
    if freqs[-1] + j >= 4:
        return 9, h
    if freqs[-1] + freqs[-2] + j >= 5:
        return 8, h
    if freqs[-1] + j >= 3:
        return 7, h
    if freqs[-1] + freqs[-2] + j >= 4:
        return 6, h
    if freqs[-1] + j >= 2:
        return 5, h
    if freqs[-1] + j >= 1:
        return 4, h

    return 69, h


s = sorted(hands, key=lambda h: value(h[0]))
acc = 0
for i, h in enumerate(s, 1):
    acc += i * int(h[1])
print(acc)

ROCK = 1
PAPER = 2
SCISSORS = 3

LOSS = 0
DRAW = 3
WIN = 6




def beats(a, b):
    return (
        (a == ROCK and b == SCISSORS) or 
        (a == SCISSORS and b == PAPER) or 
        (a == PAPER and b == ROCK)
    ) 

op = {
    "A": ROCK,
    "B": PAPER,
    "C": SCISSORS,
}

me = {
    "X": 0,
    "Y": 3,
    "Z": 6,
}

def main():
    with open("inp") as f:
        score = 0
        for line in f.readlines():
            inp = line.strip().split(" ")
            a, out =  op[inp[0]], me[inp[1]]
            score += out

            if out == DRAW:
                predicted = a
            elif out == LOSS:
                predicted = ((a + 1) % 3) + 1
            elif out == WIN:
                predicted =  (a % 3) + 1

            print(a, out, predicted)
            score += predicted
    return score

            

r = main()
print(r)
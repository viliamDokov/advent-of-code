from collections.abc import Callable

mod = 1

class Monkey():

    def __init__(self, items, operation, test, success, fail):
        self.items: list[int] = items
        self.operation: Callable[[int], int] = operation
        self.test:  Callable[[int], bool]  = test
        self.success: int = success
        self.fail: int = fail
        self.inspected = 0


    def inspect(self):
        global mod
        self.inspected += 1
        it = self.items.pop(0)
        it = self.operation(it)
        it = it % mod
        if self.test(it):
            return self.success, it
        else: 
            return self.fail, it
def make_operation(op, num, f):
    if f:
        if op == "*":
            return lambda x: x * x
        if op == "+":
            return lambda x: x + x
    else:
        if op == "*":
            return lambda x: x * num
        if op == "+":
            return lambda x: x + num
        
def make_test(num):
    return lambda x : x % num == 0

def main():
    global mod
    with open("inp") as f:
        monkeys: list[Monkey] = []
        while True:
            lines = [f.readline().strip() for i in range(7)]
            if lines[0] == "":
                break 
            items = [int(it) for it in lines[1][16:].split(", ")]
            if "*" in lines[2]:
                sign = "*" 
            if "+" in lines[2]:
                sign = "+" 

            second_term = lines[2].split(sign)[-1].strip()
            if second_term == "old":
                print("op", 1)
                op = make_operation(sign, 69, True)
            else:
                print("op", 2, sign)
                num = int(second_term)
                op = make_operation(sign, num, False)


            num = int(lines[3].split("by")[-1].strip())
            mod = mod * num
            test = make_test(num)
            success = int(lines[4].split("monkey")[-1].strip()) 
            fail = int(lines[5].split("monkey")[-1].strip())

            print("op", op)
            monkeys.append(Monkey(items, op, test, success, fail))

        print("action")
        for i in range(10000):
            for monkey in monkeys:
                while len(monkey.items) > 0:
                    nxt, it = monkey.inspect()
                    monkeys[nxt].items.append(it)

        print(max(monkeys, key = lambda x : x.inspected).inspected)
        most1 = [m for m in monkeys if m.inspected == max(monkeys, key = lambda x : x.inspected).inspected][0]

        monkeys.remove(most1)
        most2 = [m for m in monkeys if m.inspected == max(monkeys, key = lambda x : x.inspected).inspected][0]

        return most1.inspected * most2.inspected


print(main())

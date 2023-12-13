def get_ub(time, svalves):
    n = len(svalves)
    acc = 0
    for i, valve in enumerate(svalves):
        x = time - 2*i
        if x > 0:
            acc += valve * x 
    return acc

print(get_ub(8,[10,5,1]))
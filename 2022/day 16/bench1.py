from random import randint, choices
from time import perf_counter
n = 60

# lst = [ ] 
# dct = {}
# for i in range(n):
#     tmp = []
#     for j in range(n):
#         num = randint(0,10)
#         tmp.append(num)
#         dct[(str(i),str(j))] = num
#     lst.append(tmp)

# m = 100_000
# t1 = perf_counter()
# for i in range(m):
#     x,y = randint(0,n-1),randint(0,n-1)
#     tmp = lst[x][y]
# t2 = perf_counter()

# print("list", t2 - t1)

# t1 = perf_counter()
# for i in range(m):
#     x,y = randint(0,n-1),randint(0,n-1)
#     tmp = dct[(str(x),str(y))]
# t2 = perf_counter()

# print("dict", t2 - t1)

nums = list(range(n))
chosen = choices(nums, k=20)
mask = 0
for num in chosen:
    mask |= 1 << num

full = set(nums)
seen = set(chosen)

print(mask) 
m = 1000_0000
acc = 0

t1 = perf_counter()
for i in range(m):
    mask = ~mask
    # acc += (1 << t & mask) != 0 
t2 = perf_counter()
print("maks", t2-t1)

t1 = perf_counter()
for i in range(m):
    seen = full.difference(seen)
    # t = randint(0,n-1)
    # acc += t in seen
t2 = perf_counter()
print("set", t2-t1)


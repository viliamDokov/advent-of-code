from copy import deepcopy

def main():
    with open("out.txt", "w") as w:
        with open("inp") as f:
            n = 99
            trees = [
                [0 for _ in range(n)] 
                for i in range(n)]
            for i, line in enumerate(f.readlines()):
                s = line.strip()
                for j in range(n):
                    trees[i][j] = int(s[j])
            
            b = 0
            for i in range(n):
                for j in range(n):
                    t = 0
                    a = 0
                    print(i,j)
                    for k in range(1,j+1):
                        a += 1
                        if trees[i][j] <= trees[i][j-k]:
                            break
                    # print(1, a)
                    t = a
                    a = 0
                    for k in range(j+1,n):
                        a += 1
                        if trees[i][j] <= trees[i][k]:
                            break
                    # print(2, a)
                    t = t * a
                    a = 0

                    for k in range(1,i+1):
                        a += 1
                        if trees[i][j] <= trees[i - k][j]:
                            break
                    
                    # print(3, a)

                    t = t * a
                    a = 0
                    for k in range(i+1,n):
                        a += 1
                        if trees[i][j] <= trees[k][j]:
                            break
                    # print(4, a)
                    t = t* a
                    if t > b:
                        b = t
                    w.write(f"t {t} {i} {j}\n")
            return b
                
print(main())





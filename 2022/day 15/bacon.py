from bisect import bisect_left
from dataclasses import dataclass, field



def get_blocked(sensor, row, rad, blocked):


    return blocked
    

def main():
    sensors = []
    with open("inp") as f:
        for line in f.readlines():
            tmp = line[10:].strip().split(": closest beacon is at ")
            s = tmp[0].split(", ")
            xs = int(s[0][2:])
            ys = int(s[1][2:]) 
            b = tmp[1].split(", ")
            xb = int(b[0][2:])
            yb = int(b[1][2:])
            dist = abs(xs - xb) + abs(ys - yb)

            sensors.append((xs,ys,dist))


        n = 4000000
        for i in range(n + 1):
            blocked = []
            for sensor in sensors:
                rad = sensor[2] 
                x = sensor[0]
                diff = abs(sensor[1] - i)
                rad = rad - diff
                if rad >= 0:
                    start = x - rad
                    end = x + rad + 1
                    blocked.append((start,end))

            blocked = sorted(blocked, key=lambda x: x[0])


            mn = 0
            for j, _ in enumerate(blocked):
                if blocked[j][0] < 0: blocked[j] = (0,  blocked[j][1])
                else: break

            mx = n
            start = None
            end =  None
            acc = 0

            for it in blocked:
                if start == None:
                    start = it[0]
                    end = it[1]
                    if end > mx:
                        end = mx
                        break
                elif end >= it[0]:
                    if it[1] > end:
                        end = it[1]
                        if end > mx:
                            end = mx
                            break
                else:
                    print("AAAAAAAAAsx",end, i, end*4000000 + i)
                    break
                    acc += end - start + 1
                    start = None
                    end =  None

            if start != None:
                acc += end - start + 1
        
            # print("thing",i,acc)

main()
from z3 import Real, solve, Real

t1 = Real("t1")
t2 = Real("t2")
t3 = Real("t3")

p1 = Real("p1")
p2 = Real("p2")
p3 = Real("p3")

v1 = Real("v1")
v2 = Real("v2")
v3 = Real("v3")

pA1 = 152594199160345
pA2 = 147562599184759
pA3 = 291883234654893

vA1 = 229
vA2 = 220
vA3 = -31

pB1 = 181402578613976
pB2 = 206158696386036
pB3 = 294595238970734

vB1 = 179
vB2 = 99
vB3 = -32

pC1 = 306345582484815
pC2 = 290719456201785
pC3 = 306246299945991

vC1 = -19
vC2 = -64
vC3 = -43


solve(
    pA1 + (t1 * vA1) == p1 + (t1 * v1),
    pA2 + (t1 * vA2) == p2 + (t1 * v2),
    pA3 + (t1 * vA3) == p3 + (t1 * v3),
    pB1 + (t2 * vB1) == p1 + (t2 * v1),
    pB2 + (t2 * vB2) == p2 + (t2 * v2),
    pB3 + (t2 * vB3) == p3 + (t2 * v3),
    pC1 + (t3 * vC1) == p1 + (t3 * v1),
    pC2 + (t3 * vC2) == p2 + (t3 * v2),
    pC3 + (t3 * vC3) == p3 + (t3 * v3),
)

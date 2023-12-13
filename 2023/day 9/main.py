# dining philosophers

from threading import Thread, Lock, Semaphore
from time import sleep
from random import uniform

N = 10000
s = Semaphore(N - 1)

forks = [Lock() for _ in range(N)]
eatings = 0
l = Lock()


def think():
    sleep(uniform(0.01, 0.05))


def get_forks(i):
    global eatings
    s.acquire()
    forks[(i - 1) % N].acquire()
    forks[i].acquire()

    with l:
        print("eats", eatings)
        eatings += 1


def leave_forks(i):
    global eatings
    s.release()
    forks[(i - 1) % N].release()
    forks[i].release()
    with l:
        eatings -= 1


def eat():
    ...


def philo(i):
    print(f"Philo {i}")
    while True:
        think()
        get_forks(i)
        eat()
        leave_forks(i)


for i in range(N):
    t = Thread(target=philo, args=[i])
    t.daemon = True
    t.start()
while True:
    pass

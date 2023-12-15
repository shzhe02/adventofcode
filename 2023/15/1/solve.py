from functools import reduce
line = input()

total = 0
for x in line.split(","):
    res = reduce(lambda a, b: ((a + b) * 17) % 256, map(ord, x), 0)
    total += res

print(total)

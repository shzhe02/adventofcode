from functools import reduce
line = input()

total = 0
for x in line.split(","):
    curr = 0
    for char in x:
        curr += ord(char)
        curr *= 17
        curr %= 256
    total += curr

print(total)


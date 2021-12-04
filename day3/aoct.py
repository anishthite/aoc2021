from collections import Counter
d = Counter()
f = open('input.txt').readlines()
for b in f:
    for i, c in enumerate(b.rstrip()):
        if c == '1':
            d[i] +=1
        else: 
            d[i] -=1
e = ''.join(['1'  if d[i] > 0 else '0' for i in range(12)])
g = ''.join(['0' if i == '1' else '1' for i in e ])
print(int(e, 2) * int(g,2))

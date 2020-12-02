f = open("input.txt", 'r')

dictionary = {}

for line in f:
    dictionary[int(line.strip())] = True

f.close()

for key in dictionary.keys():
   if (2020-key) in dictionary.keys():
    print(key*(2020-key))

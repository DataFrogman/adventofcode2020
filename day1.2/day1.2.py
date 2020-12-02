f = open("input.txt", 'r')

dictionary = {}

for line in f:
    dictionary[int(line.strip())] = True

f.close()

for key in dictionary.keys():
    for key2 in dictionary.keys():
        if (2020-key-key2) in dictionary.keys():
            print(key*key2*(2020-key-key2))
            exit()

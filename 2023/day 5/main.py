import re

with open("inp") as f:
    allInput = f.read()

splitInput = allInput.split(":")
seeds = re.findall("\d+", splitInput[1])

seedToSoil = re.findall("\d+ \d+ \d+", splitInput[2])
# print (seedToSoil)
soilToFertilizer = re.findall("\d+ \d+ \d+", splitInput[3])
fertilizerToWater = re.findall("\d+ \d+ \d+", splitInput[4])
waterToLight = re.findall("\d+ \d+ \d+", splitInput[5])
lightToTemperature = re.findall("\d+ \d+ \d+", splitInput[6])
temperatureToHumidity = re.findall("\d+ \d+ \d+", splitInput[7])
humidityToLocation = re.findall("\d+ \d+ \d+", splitInput[8])

checkedIds = []
for idx, seed in enumerate(seeds):
    # for thing in splitInput[2:]:
    #     seedToSoil = re.findall("\d+ \d+ \d+", thing)
    # print(seedToSoil)
    for soils in seedToSoil:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break

print(seeds)
checkedIds = []
for idx, seed in enumerate(seeds):
    for soils in soilToFertilizer:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break


print(seeds)
checkedIds = []
for idx, seed in enumerate(seeds):
    for soils in fertilizerToWater:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break
checkedIds = []
print(seeds)
for idx, seed in enumerate(seeds):
    for soils in waterToLight:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break

print(seeds)
checkedIds = []
for idx, seed in enumerate(seeds):
    for soils in lightToTemperature:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break

print(seeds)
checkedIds = []
for idx, seed in enumerate(seeds):
    for soils in temperatureToHumidity:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break
print(seeds)
checkedIds = []
for idx, seed in enumerate(seeds):
    for soils in humidityToLocation:
        soil = soils.split(" ")
        if int(seed) >= int(soil[1]) and int(seed) <= int(soil[1]) + int(soil[2]) - 1:
            seeds[idx] = str(int(seed) - int(soil[1]) + int(soil[0]))
            break

print(seeds)
# print(seed)
minSeed = int(seeds[0])
for seed in seeds:
    if minSeed > int(seed):
        minSeed = int(seed)

print(minSeed)

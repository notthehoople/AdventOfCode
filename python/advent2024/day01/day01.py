from argparse import ArgumentParser

def countItem(itemList, itemToCount):
    count = 0

    for pos, item in enumerate(itemList):
        if itemToCount == item:
            count += 1

    return count

def main():
    parser = ArgumentParser()
    parser.add_argument('--file', dest="filename", type=str, required=False, default="test-input", help='File with puzzle input')
    parser.add_argument('--part', dest="part", type=str, required=False, default="a", help='Puzzle Part to complete')
    args = parser.parse_args()

    filename = args.filename
    part = args.part

    firstList = []
    secondList = []

    file = open(filename, "r")

    for line in file:
        x=line.split("   ")
        firstList.append(int(x[0]))
        secondList.append(int(x[1].replace('\n', '')))

    file.close()

    if part == "a":
        firstList.sort()
        secondList.sort()

        distanceTotal = 0

        for pos, item in enumerate(firstList):
            distanceTotal += abs(firstList[pos] - secondList[pos])

        print ("Part a:", distanceTotal)
        return distanceTotal
    
    similarityTotal = 0
    for pos, item in enumerate(firstList):
        similarityTotal += item * countItem(secondList, item)

    print("Part b:", similarityTotal)    

if __name__ == "__main__":
    main()
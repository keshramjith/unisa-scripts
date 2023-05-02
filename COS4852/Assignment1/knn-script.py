import random
from matplotlib import pyplot as plt
import numpy as np

def generateCatDataset(sizeOfDataset):
    catDataset = []
    for i in range(sizeOfDataset):
        catDataset.append(random.randint(48, 65))
    return catDataset

def generateDogDataset(sizeOfDataset):
    dogDataset = []
    for i in range(sizeOfDataset):
        dogDataset.append(random.randint(15, 38))
    return dogDataset

def generateTestcase():
    testDataset = []
    testDataset.append(random.randint(15, 65))
    return testDataset

def main():
    datasetSize = 5
    dogDataset = generateDogDataset(datasetSize)
    catDataset = generateCatDataset(datasetSize)
    testcaseDataset= generateTestcase()
    print("dogDataset: ", dogDataset)
    print("catDataset: ", catDataset)
    print("testDataset: ", testcaseDataset)
    val = 0.
    plt.scatter(dogDataset, np.zeros_like(dogDataset) + val, color='red')
    plt.scatter(catDataset, np.zeros_like(catDataset) + val, color='blue')
    plt.scatter(testcaseDataset, np.zeros_like(testcaseDataset) + val, color='grey')
    plt.show()

if __name__ == "__main__":
    main()

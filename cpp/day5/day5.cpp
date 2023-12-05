#include <iostream>
#include <unistd.h>
#include <limits>
#include <chrono>
#include <sstream>
#include "../utils/AdventReader.h"

#include <algorithm>

#include <map>
using namespace std;

long part1(const vector<string> lines) {
    return 0;
}

long part2(const vector<string> lines) {
    return 0;
}

struct Entry
{
    long keyMin;
    long keyMax;
    long valMin;
    long valMax;


    bool checkInValRange(long val) {
        return val >= valMin && val <= valMax;
    }

    bool checkInKeyRange(long val) {
        return val >= keyMin && val <= keyMax;
    }
};
struct SeedNums {
    long min;
    long max;
};

void print(const Entry &s) {
    cout << "Keys: { " << s.keyMin << ", " << s.keyMax << "} Vals : { " << s.valMin << ", " << s.valMax <<"}" << endl;
}

typedef vector<vector<Entry>> MapList;


void readToNextSection(ifstream *inFile, string sectionHead) {
    string line;
    while (getline(*inFile, line)) {

        if (line.compare(sectionHead) == 0) {
            return;
        }
        
    }
}

vector<long> readSeedList(ifstream *inFile) {

    vector<long> seedList;
    string seedLine;
    getline(*inFile, seedLine);

    seedLine.replace(0, 7, "");

    seedList = avreader::splitToLong(seedLine, " ");

    readToNextSection(inFile, "seed-to-soil map:");

    return seedList;
}

map<long, long> readAndPopulateMap(ifstream *inFile) {
    map<long, long> myMap;

    string line;
    vector<long> row;
    while (getline(*inFile, line)) {
        row = avreader::splitToLong(line, " ");

        if (row.empty()) {
            getline(*inFile, line);//Read one more to the next section header
            break;
        }
        

        //row[0] is val start
        //row[1] is key start
        //row[2] is length

        Entry entry;
        entry.keyMin = row[1];
        entry.keyMax = row[1] + row[2] - 1;
        entry.valMin = row[0];
        entry.valMax = row[0] + row[2] - 1;
    }

    return myMap;

}

vector<Entry> readAndPopulateVec(ifstream *inFile) {
    vector<Entry> vec;

    string line;
    vector<long> row;
    while (getline(*inFile, line)) {
        row = avreader::splitToLong(line, " ");

        if (row.empty()) {
            getline(*inFile, line);//Read one more to the next section header
            break;
        }

        //row[0] is val start
        //row[1] is key start
        //row[2] is length
        Entry entry;
        entry.keyMin = row[1];
        entry.keyMax = row[1] + row[2] - 1;
        entry.valMin = row[0];
        entry.valMax = row[0] + row[2] - 1;

        vec.push_back(entry);
    }

    return vec;

}

int solveSeedList(vector<vector<Entry>> mappings, vector<long> seedList) {
    map<long, long> seedToLocationMap;

    for (auto seed : seedList) {
        long nextKey = seed;
        //At each mapping, we are trying to find the match for this 'nextKey'
        for (auto entries : mappings) {

            //For each entry, if  entry.MinVal <= nextKey <= entry.MaxVal calculate nextKey and break
            for (auto entry : entries) {
                
                if (entry.keyMin <= nextKey && entry.keyMax >= nextKey) {
                    //The next key is the found value plus the key offset
                    nextKey = entry.valMin + (nextKey - entry.keyMin);
                    break;
                }
            }
        }
        seedToLocationMap[seed] = nextKey;
    }

    long min = INT64_MAX;
    long minSeed = 0;
    for (auto r : seedToLocationMap  ) {

        if (r.second < min)
        {
            min = r.second;
            minSeed = r.first;
        }
        
    }
    return min;
}

int getSeedMapping(MapList mappings, int startVal) {
    int seedVal = startVal;//The value that we are trying to find key for
    bool foundSeed;
    //Loop through mappings in reverse order
    for (int mapping = mappings.size()-1; mapping > -1; mapping--){

        // cout << "Looking at Mapping = " << mapping << " For Key = " << seedVal << endl;
        //For each mapping, try to find match
        for (auto match : mappings[mapping]) {

            print(match);
            if (match.checkInValRange(seedVal)) {
                //If match found, then we have found a mapped entry to this key
                //Need to get the mapped key
                // cout << "Found Match for Seed = " << seedVal << endl;
                seedVal = match.keyMin + (seedVal - match.valMin);
                // cout << "Key for Seed = " << seedVal << endl;
                if (mapping == 0) foundSeed = true;
                break;
            }
        }
    }

    return foundSeed ? seedVal : -1;
}
int search(MapList mappings, int maxLocation, int bound, int offset) {
    // int bestLocation;
    // int bestSeed;

    for (int i = 0 + offset; i <= maxLocation; i+= bound)
    {
        int foundSeed = getSeedMapping(mappings, i);

        if (foundSeed != -1 )
        {
            cout << "Found New Max at " << i << endl;
            return i;//Once we find a mapping, return the new max
            // maxLocation = i;//Return
            // bestSeed = foundSeed;
            // cout << "Found Best Location = " << maxLocation << endl;
            // return bestSeed;
        }
    }
    
    return -1;

}

int solveP2(MapList mappings, int maxLocation) {

    //Max Location is the highest known mapping to a location
    //Thus any additional default mapping must be higher

    //We need to search 0-maxLocation - 1 in large bounds and slowly decrease as our possiilities decrease

    //1) Search in bounds of ? from 0 to maxLocation
    //      a) Iterate up through the maps and find a valid seed
    //          if seed valid and < max, set max to search Location and store bestSeed

    int bound = 1024;
    int offset = 0;
    while (bound != 1) {
        int result = search(mappings, maxLocation, bound, offset);

        if (result == -1) {
            //No match was found, shift
            offset = (offset + 1) % bound;
            if (offset == 0) {
                cout << "Wrapped all the way around on this bound, error" << endl;
                exit(1);
            }
        } else {
            //match was found
            maxLocation = result;
            offset = 0;
            bound = bound / 2;
        }
    }
    return maxLocation;
}



int main() {

    // auto start = timeSinceEpochMillisec();

    ifstream inFile("/home/benn/CODE/adventCode/sample");

    vector<long> seedList = readSeedList(&inFile);
    long minSeed = INT64_MAX;

    for (auto s : seedList) {
        if (s < minSeed)
        {
            minSeed = s;
        }
    }
    cout << "Min Seed = " << minSeed << endl;
    vector<Entry> seedToSoil = readAndPopulateVec(&inFile);
    vector<Entry> soilFertilizer = readAndPopulateVec(&inFile);
    vector<Entry> fertilizerWater = readAndPopulateVec(&inFile);
    vector<Entry> waterLight = readAndPopulateVec(&inFile);
    vector<Entry> lightTemperature = readAndPopulateVec(&inFile);
    vector<Entry> tempHumidity = readAndPopulateVec(&inFile);
    vector<Entry> humidityLocation = readAndPopulateVec(&inFile);

    vector<vector<Entry>> mappings;
    mappings.push_back(seedToSoil);
    mappings.push_back(soilFertilizer);
    mappings.push_back(fertilizerWater);
    mappings.push_back(waterLight);
    mappings.push_back(lightTemperature);
    mappings.push_back(tempHumidity);
    mappings.push_back(humidityLocation);


    cout << "Part 1 = " << solveSeedList(mappings, seedList) << endl;//173706076 PART1

    long minTemp = INT64_MAX;
    long minHumidity = 0;

    vector<int> locations;
    int maxLocation = INT32_MIN;
    for (auto p : humidityLocation) {

        if (p.valMax > maxLocation) {
            maxLocation = p.valMax;
        }
        // for (int i = p.valMin; i <= p.valMax; i++) {
        //     locations.push_back(i);
        // }
    }
    vector<Entry> seedRanges;
    for (int i = 0; i < seedList.size(); i += 2) {
        Entry e;
        e.valMin = seedList[i];
        e.valMax = seedList[i] + seedList[i+1];
        seedRanges.push_back(e);
    }
    mappings.insert(mappings.begin(), seedRanges);


    cout << "Max Location = " << maxLocation << endl;

    cout << "Part 2 = " << solveP2(mappings, maxLocation) << endl;
    // cout << "Part 2 = " << solveSeedList(mappings, seedList2) << endl;//173706076 PART1
    
    // cout << "Part 2 = " << solveP2(mappings, nextSeedList) << endl;//

    // auto end = std::chrono::system_clock::now().time_since_epoch().count();

    // cout << "TIME ELAPSED = " << (timeSinceEpochMillisec() - start) << endl;//1154, 1566, 1900

    return 0;
}

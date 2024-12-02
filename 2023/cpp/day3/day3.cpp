#include <iostream>
#include <chrono>
#include "../utils/AdventReader.h"

#include <algorithm>

#include <map>

using namespace std;

int part1(const vector<string> lines) {
    return 0;
}

int part2(const vector<string> lines) {
    return 0;
}

uint64_t timeSinceEpochMillisec() {
  using namespace std::chrono;
  return duration_cast<microseconds>(system_clock::now().time_since_epoch()).count();
}

int main() {

    auto start = timeSinceEpochMillisec();

    vector<string> lines = avreader::parseFile("day1.txt");

    vector<vector<string>> grid;

    for (auto line : lines) {

    }



    cout << "Part 1 = " << part1(lines) << endl;//55130
    cout << "Part 2 = " << part2(lines) << endl;//54985

    auto end = std::chrono::system_clock::now().time_since_epoch().count();

    cout << "TIME ELAPSED = " << (timeSinceEpochMillisec() - start) << endl;//1154, 1566, 1900

    return 0;
}


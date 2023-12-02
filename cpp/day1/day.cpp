#include <iostream>
#include <chrono>
#include "../utils/AdventReader.h"

#include <algorithm>

#include <map>

using namespace std;

int getCombinedVal(const string line) {
    string first, last;
    size_t foundIndex;

    foundIndex = line.find_first_of("0123456789");
    first = line.substr(foundIndex, 1);

    foundIndex = line.find_last_of("0123456789");
    last = line.substr(foundIndex, 1);

    return stoi(first.append(last));

}

int part1(const vector<string> lines) {

    int sum = 0;

    for (auto line : lines) {
        sum += getCombinedVal(line);
    }

    return sum;
}

int part2(const vector<string> lines) {

    map<string, string> replacements = {
        {"one", "o1e"},
        {"two", "t2o"},
        {"three", "t3e"},
        {"four", "f4r"},
        {"five", "f5e"},
        {"six", "s6x"},
        {"seven", "s7n"},
        {"eight", "e8t"},
        {"nine", "n9e"},

    };

    int sum = 0;

    for (auto line : lines) {

        for (const auto& [key, value] : replacements) {
            size_t startIndex = 0;

            while((startIndex = line.find(key)) != string::npos) {
                line.replace(startIndex, value.length(), value);
                startIndex += value.length();
            }
        }
        sum += getCombinedVal(line);
    }


    return sum;
}
uint64_t timeSinceEpochMillisec() {
  using namespace std::chrono;
  return duration_cast<microseconds>(system_clock::now().time_since_epoch()).count();
}

int main() {

    auto start = timeSinceEpochMillisec();

    vector<string> lines = avreader::parseFile("day1.txt");



    cout << "Part 1 = " << part1(lines) << endl;//55130
    cout << "Part 2 = " << part2(lines) << endl;//54985

    auto end = std::chrono::system_clock::now().time_since_epoch().count();

    cout << "TIME ELAPSED = " << (timeSinceEpochMillisec() - start) << endl;

    return 0;
}

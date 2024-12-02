#include <algorithm>
#include <iostream>
#include <fstream>
#include <vector>
#include <memory>
#include <stdexcept>

#include <sstream>

#include <string>
#include "../../../../../../usr/include/c++/11/bits/stream_iterator.h"
#include <chrono>

using namespace std;

namespace avreader
{
    const string FILE_PATH = "/home/benn/CODE/adventCode/";

    vector<string> parseFile(const string fileName) {
        vector<string> lines;
    
        ifstream inFile(FILE_PATH + fileName);

        string line;

        while (std::getline(inFile, line)) {
            lines.push_back(line);
        }

        return lines;
    }

    ifstream openFile(const string fileName) {

        ifstream inFile(FILE_PATH + fileName);
        return inFile;
    }

    uint64_t timeSinceEpochMillisec() {
        using namespace std::chrono;
        return duration_cast<milliseconds>(system_clock::now().time_since_epoch()).count();
    }

    uint64_t timeSinceEpochMicros() {
        using namespace std::chrono;
        return duration_cast<microseconds>(system_clock::now().time_since_epoch()).count();
    }

    template <typename F>
    void runAndPrintMicros(F&& runner) {
        auto start = timeSinceEpochMicros();
        runner();
        cout << timeSinceEpochMicros() - start << " Microseconds" << endl;

    }

    vector<vector<string>> parseFileToGrid(const string fileName) {
        vector<vector<string>> lines;
    
        ifstream inFile(FILE_PATH + fileName);

        string line;

        while (std::getline(inFile, line)) {
            vector<string> row;
            lines.push_back(row);
        }

        return lines;
    }

    vector<vector<string>> parseFileToGroups(const string fileName, const string delimeter) {
        vector<vector<string>> groups;
    
        ifstream inFile(FILE_PATH + fileName);

        string line;

        vector<string> *group = new vector<string>();
        while (std::getline(inFile, line)) {
            if (line == delimeter) {
                groups.push_back(*group);
                group = new vector<string>();
                continue;
            }
            group->push_back(line);
        }
        groups.push_back(*group);

        return groups;
    }
 
    template <typename T>
    vector<T> slice(vector<T> toSlice, int from = 0, int to = -1) {
        if (to == -1)
        {
            to = toSlice.size();
        }
        
        return vector<T>(toSlice.begin() + from, toSlice.end());
    }
 
} // namespace avreader

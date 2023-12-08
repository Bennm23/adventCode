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

    vector<vector<string>> parseFileToGrid(const string fileName) {
        vector<vector<string>> lines;
    
        ifstream inFile(FILE_PATH + fileName);

        string line;

        while (std::getline(inFile, line)) {
            vector<string> row;
            

            // std::stringstream ss(line);

            // std::istream_iterator<std::string> begin(ss);
            // std::istream_iterator<std::string> end;
            // std::vector<std::string> vstrings(begin, end);
            // std::copy(vstrings.begin(), vstrings.end(), std::ostream_iterator<std::string>(std::cout, "\n")); 
            
            lines.push_back(row);
        }

        return lines;
    }

 
} // namespace avreader

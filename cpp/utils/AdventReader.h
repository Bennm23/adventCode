
#include <iostream>
#include <fstream>
#include <vector>
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

    vector<string> split(string startString, string delimiter) {
        vector<string> res;
        size_t pos = 0;
        std::string token;
        while ((pos = startString.find(delimiter)) != std::string::npos) {
            token = startString.substr(0, pos);
            res.push_back(token);
            std::cout << token << std::endl;
            startString.erase(0, pos + delimiter.length());
        }
        res.push_back(startString);

        return res;
    }

    vector<long> splitToLong(string startString, string delimiter) {
        vector<long> res;
        size_t pos = 0;
        std::string token;
        while ((pos = startString.find(delimiter)) != std::string::npos) {
            token = startString.substr(0, pos);
            res.push_back(stol(token));
            startString.erase(0, pos + delimiter.length());
        }
        try {
            res.push_back(stol(startString));
        } catch (...){}

        return res;
    }



    
} // namespace avreader

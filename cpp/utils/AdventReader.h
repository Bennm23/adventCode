
#include <iostream>
#include <fstream>
#include <vector>
#include <string>

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
    
} // namespace avreader

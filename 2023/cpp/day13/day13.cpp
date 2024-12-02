#include <iostream>
#include <chrono>
#include "../utils/AdventReader.h"
#include "../utils/Maths.h"

#include <algorithm>

#include <map>

using namespace std;
using namespace avreader;


vector<vector<char>> buildGrid(const vector<string> &group) {
    vector<vector<char>> grid;

    for (auto line : group) {
        vector<char> tmp;
        for (auto c : line) {
            tmp.push_back(c);
        }
        grid.push_back(tmp);
    }

    return grid;
}

void checkReflections(const vector<vector<char>> &grid, int &p1, int &p2) {

    for (int row = 0; row < grid.size(); row++) {
        int rowRange = std::min(row, int(grid.size() - row));
        int upRow = row - 1;
        int downRow = row;

        int matchCount = 0;

        for (int dRow = 0; dRow < rowRange; dRow++) {
            auto upLine = grid[upRow - dRow];
            auto downLine = grid[downRow + dRow];

            for (int index = 0; index < upLine.size(); index++) {
                if (upLine[index] == downLine[index])
                {
                    matchCount++;
                }
            }
        }

        auto perfectMatch = rowRange * grid[row].size();
        if (matchCount == perfectMatch)
        {
            p1 = row;
        }
        else if (matchCount == perfectMatch - 1)
        {
            p2 = row;
        }
        
    }
}

int main() {


    runAndPrintMicros([]() {
        vector<vector<string>> groups = parseFileToGroups("day13.txt", "");

        vector<vector<string>> grid;

        long p1,p2;
        for (auto group : groups) {

            auto grid = buildGrid(group);
            
            int ver1, ver2 = 0;
            checkReflections(grid, ver1, ver2);

            auto transposed = maths::transpose(grid);
            int hor1, hor2 = 0;
            checkReflections(transposed, hor1, hor2);

            p1 += ver1 * 100 + hor1;
            p2 += ver2 * 100 + hor2;
        }


        cout << "Part 1 = " << p1 << endl;//37975
        cout << "Part 2 = " << p2 << endl;//32497

    });//4831, 7131, 4785

    return 0;
}
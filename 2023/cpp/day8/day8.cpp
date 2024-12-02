#include <iostream>
#include <math.h>
#include <unistd.h>
#include <limits>
#include <chrono>
#include <sstream>
#include <thread>
#include <future>
#include "../utils/AdventReader.h"
#include "../utils/Maths.h"
#include "../utils/StringUtils.h"

#include <algorithm>

#include <map>

using namespace std;

struct Node {

    string name;
    string leftNode;
    string rightNode;
    Node *left;
    Node *right;


    string asString() {
        return name;
    }
};


int main() {
    string instructionSet;
    ifstream in = ifstream(avreader::FILE_PATH + "day8.txt");

    getline(in, instructionSet);

    cout << "Instructions" << endl;
    cout << instructionSet << endl;

    string hold;
    getline(in, hold);//Clear blank

    map<string, Node *> nodeMap;
    vector<Node *> startNodes;

    while (getline(in, hold)) {
        string key = hold.substr(0, 3);


        vector<string> lr = avstrings::split(avstrings::getTextWithinParens(hold), ", ");
  
        nodeMap[key] = new Node { key, lr[0], lr[1] };
        
        if (key.back() == 'A')
        {
            startNodes.push_back(nodeMap[key]);
        }
        
    }

    bool solved = false;

    long steps = 0;

    Node *curr;
    curr = nodeMap["AAA"];

    while (!solved){

        for (string::iterator it = instructionSet.begin(); it != instructionSet.end(); ++it) {
            steps++;
            if (*it == 'R')
            {
                curr = nodeMap[curr->rightNode];
            }
            else
            {
                curr = nodeMap[curr->leftNode];
            }
            if (curr->name.compare("ZZZ") == 0)
            {
                solved = true;
                break;
            }
        }

    }


    cout << "P1 = " << steps << endl;

    steps = 0;
    solved = false;
    vector<long> solutions;

    for (auto starter : startNodes) {
        long stepCount = 0;        
        bool solved = false;
        //For each start Node, find the terminating point
        while (!solved) {

            for (string::iterator it = instructionSet.begin(); it != instructionSet.end(); ++it) {
                stepCount++;
                if (*it == 'R')
                {
                    starter = nodeMap[starter->rightNode];
                }
                else
                {
                    starter = nodeMap[starter->leftNode];
                }
                if (starter->name.back() == 'Z')
                {
                    solved = true;
                    break;
                }
            }

        }
        solutions.push_back(stepCount);
    }

    long multiple = maths::lcm(solutions[0], solutions[1]);
    for (int i = 2; i < solutions.size(); i++) {
        multiple = maths::lcm(multiple, solutions[i]);
    }

    cout << "P2 = " << multiple << endl;//21165830176709

    return 0;
}
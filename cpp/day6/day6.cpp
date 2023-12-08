
#include "../utils/AdventReader.h"

using namespace std;

struct Race {
    long time, recordDistance;

    string asString() {
        return avreader::string_format("Time = %d, Distance = %d", time, recordDistance);
    }

    long computeWinners() {
        long total = 0;
        for (long i = 0; i < time; i++) {
            long score = i * (time - i);
            if (score > recordDistance) total++;
        }
        return total;
    }
};

int main() {

    vector<string> lines = avreader::parseFile("day6.txt");

    for (auto line : lines) {
        cout << line << endl;
    }

    vector<long> times = avreader::parseLineToLongs(lines[0]);
    vector<long> distances = avreader::parseLineToLongs(lines[1]);
    vector<Race> races;
    
    for (int i = 0; i < times.size(); i++) {

        races.push_back(Race{times[i], distances[i]});

    }
    int total = 1;
    for (auto s : races) {
        total *= s.computeWinners();
    }

    cout << "Part 1 = " << total << endl;//771628

    long time = avreader::parseLineToOneLong(lines[0]);
    long distance = avreader::parseLineToOneLong(lines[1]);

    cout << "Time = " << time << endl;
    cout << "Distance = " << distance << endl;

    Race p2 = {time, distance};

    cout << "Part 2 = " << p2.computeWinners() << endl;//27363861

    return 0;
}
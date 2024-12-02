#include <iostream>
#include <chrono>
#include "../utils/AdventReader.h"
#include "../utils/StringUtils.h"

#include <algorithm>
#include <functional>

#include <map>

using namespace std;


typedef map<size_t, long> ValMap;

template <typename T>
std::size_t hash_combine(std::size_t seed, const T& v) {
    std::hash<T> hasher;
    seed ^= hasher(v) + 0x9e3779b9 + (seed << 6) + (seed >> 2);
    return seed;
}
size_t my_hash(string cfg, vector<int> nums) {

    std::hash<string> hash1;
    std::hash<int> hash2;

    size_t seed = hash_combine(hash1(cfg), 0);

    for (const auto& el : nums) {
        seed = hash_combine(seed, hash2(el));
    }

    return seed;
}


long solve(string cfg, vector<int> nums, ValMap &vals) {
    if (cfg.empty())
    {
        return nums.empty() ? 1 : 0;
    }
    if (nums.empty())
    {
        return cfg.find("#") != string::npos ? 0 : 1;
    }

    auto rec = my_hash(cfg, nums);
    if (vals.count(rec)) return vals[rec];
    long sum = 0;

    if (cfg[0] == '.' || cfg[0] == '?') {

        sum += solve(cfg.substr(1), nums, vals);
    }
    
    if (cfg[0] == '#' || cfg[0] == '?') {

        if (nums[0] <= cfg.size() && !avstrings::contains(cfg.substr(0, nums[0]), ".") &&
        
            (nums[0] == cfg.size() || cfg.substr(nums[0], 1) != "#")
        
        ){
            if (nums[0] + 1 > cfg.size())
            {
                sum += solve("", avreader::slice(nums, 1), vals);
            }
            else {
                sum += solve(cfg.substr(nums[0] + 1), avreader::slice(nums, 1), vals);
            }
            
        }
    }

    

    vals[rec] = sum;
    return sum;
}

void solve() {


    vector<string> lines = avreader::parseFile("day12.txt");

    ValMap vals;

    long p1 = 0;
    long p2 = 0;
    for (auto line : lines) {

        auto split = avstrings::split(line, " ");

        auto cfg = split[0];
        auto nums = avstrings::stringsToInt(avstrings::split(split[1], ","));

        p1 += solve(cfg, nums, vals);

        string cfg2 = "";
        for (int i = 0; i < 5; i++) {
            cfg2.append(cfg);
            if (i != 4)
            {
                cfg2.append("?");
            }
        }

        vector<int> nums2;
        for (int i = 0; i < 5; i++) {
            for (auto num : nums) {
                nums2.push_back(num);
            }
        }

        p2 += solve(cfg2, nums2, vals);
    }

    cout << "Part 1 = " << p1 << endl;
    cout << "Part 2 = " << p2 << endl;

}

int main() {

    avreader::runAndPrintMicros([]() {
        solve();
    });//490413, 494757

    return 0;
}

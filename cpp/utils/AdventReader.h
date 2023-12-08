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
    const string FILE_PATH = "/home/bennmellinger/CODE/adventCode/";

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

    struct not_digit {
        bool operator()(const char c) {
            return c != ' ' && !std::isdigit(c);
        }
    };

    vector<int> parseLineToInts(string str) {

        not_digit not_a_digit;
        std::string::iterator end = std::remove_if(str.begin(), str.end(), not_a_digit);
        std::string all_numbers(str.begin(), end);
        std::stringstream ss(all_numbers);
        std::vector<int> numbers;

        for(int i = 0; ss >> i; ) {
            numbers.push_back(i);
        }
        return numbers;
    }

    vector<long> parseLineToLongs(string str) {

        not_digit not_a_digit;
        std::string::iterator end = std::remove_if(str.begin(), str.end(), not_a_digit);
        std::string all_numbers(str.begin(), end);
        std::stringstream ss(all_numbers);
        std::vector<long> numbers;

        for(long i = 0; ss >> i; ) {
            numbers.push_back(i);
        }
        return numbers;
    }

    long parseLineToOneLong(string str) {
        not_digit not_a_digit;
        std::string::iterator end = std::remove_if(str.begin(), str.end(), not_a_digit);
        std::string all_numbers(str.begin(), end);
        std::stringstream ss(all_numbers);
        cout << "ALL = " << all_numbers << endl;

        string numbers;


        for (string s = ""; ss >> s;) {
            numbers.append(s);
        }
        return stol(numbers);
    }


    template<typename ... Args>
    std::string string_format( const std::string& format, Args ... args )
    {
        int size_s = std::snprintf( nullptr, 0, format.c_str(), args ... ) + 1; // Extra space for '\0'
        if( size_s <= 0 ){ throw std::runtime_error( "Error during formatting." ); }
        auto size = static_cast<size_t>( size_s );
        std::unique_ptr<char[]> buf( new char[ size ] );
        std::snprintf( buf.get(), size, format.c_str(), args ... );
        return std::string( buf.get(), buf.get() + size - 1 ); // We don't want the '\0' inside
    }

    
} // namespace avreader

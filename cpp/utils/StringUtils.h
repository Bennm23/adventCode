#include <algorithm>
#include <iostream>
#include <vector>

#include <sstream>
#include <string>

using namespace std;

namespace avstrings
{

   vector<string> split(string startString, string delimiter) {
        vector<string> res;
        size_t pos = 0;
        std::string token;
        while ((pos = startString.find(delimiter)) != std::string::npos) {
            token = startString.substr(0, pos);
            res.push_back(token);
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

    

    string getTextWithinParens(string input) {
        stringstream res;


        bool mark;
        for(string::iterator it = input.begin(); it != input.end(); ++it) {
            if (*it == ')')
            {
                break;
            }

            if (mark) {
                res << *it;
            }

            if (*it == '(')
            {
                mark = true;
            }
        }

        res.str();

        return res.str();
    }

    vector<int> stringsToInt(vector<string> strings) {
        vector<int> res;
        for (auto s : strings) {
            res.push_back(stoi(s));
        }
        return res;
    }

    bool contains(string str, string val) {
        return str.find(val) != string::npos;
    }
}
#include <vector>
using namespace std;

namespace maths
{
    template<class num>
    num gcd(num a, num b) {
        return (b == 0) ? a : gcd(b, a%b);
    }

    template<class num>
    num lcm(num a, num b) {
        return (a * b) / gcd(a, b);
    }

    template<class T>
    vector<vector<T>> transpose(const vector<vector<T>> &matrix) {
        vector<vector<T>> transposed(matrix[0].size(), vector<T>());

        for (size_t i = 0; i < matrix.size(); i++)
        {
            vector<T> tmp;
        
            for (size_t j = 0; j < matrix[0].size(); j++) {

                transposed[j].push_back(matrix[i][j]);
            }
        }
         
        return transposed;
    }
}
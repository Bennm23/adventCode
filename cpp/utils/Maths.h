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
}
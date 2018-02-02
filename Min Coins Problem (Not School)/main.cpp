#include <iostream>
#include <vector>
#include <sstream>
#include <algorithm>
#include <iterator>
#include <numeric>

using std::vector; using std::cin; using std::cout; using std::string;

vector<vector<int> > subsets(vector<int> input) {
  if(0 == input.size()) {
    return vector<vector<int> >(1, vector<int>());
  }
  const int last = input.back();
  input.pop_back();
  vector<vector<int> > result = subsets(input);
  result.reserve(result.size() * 2);
  const size_t resultSize = result.size();
  for(size_t i = 0; i < resultSize; ++i) {
    vector<int> tmp = result[i];
    tmp.push_back(last);
    result.push_back(tmp);
  }
  return result;
}

int solve(string input){
    std::istringstream iss(input);
    auto parsed_input = vector<int>(std::istream_iterator<int>( iss ), std::istream_iterator<int>() );
    auto coins = vector<int>(parsed_input.begin()+1, parsed_input.end());
    auto all_subsets = subsets(coins);
    auto min = coins.size()+1; 
    for(auto vec : all_subsets){
        if(std::accumulate(vec.rbegin(), vec.rend(), 0) == parsed_input[0]){
            if(vec.size() < min)
                min = vec.size();
        }
    }
    if(min > coins.size())
    	return -1;
    else return min;
}

int main(){
    string input = "150 100 50 50 50 50";
    string input1 = "130 100 20 18 12 5 5";
    string input2 = "200 50 50 20 20 10";
    auto solutions = { solve(input), solve(input1), solve(input2) };
    for(auto s : solutions)
    	cout << (s == -1 ? string("no solution") : std::to_string(s)) << std::endl;
}







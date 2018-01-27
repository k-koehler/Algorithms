#include <iostream>
#include <vector>
#include <cmath>
#include <cstring>

struct node {
	node(int n){ value = n; }
	int value;
	std::vector<node* > neighbours;
	bool isEmpty(){ return neighbours.size() == 0; }
};

typedef std::vector<node*> graph;

bool is_square(int n){
	int root(round(sqrt(n)));
	return n == root * root;
}

bool contains(graph g, int v){
	for(auto i : g) 
		if(i->value == v)
			return true;
	return false;
}

bool can_reach(node* start_node, graph& g, graph& visited, int cur_size=0){
	if(start_node->isEmpty()){
		if(cur_size < g.size())
			return false;
		else {
			std::cout<<"true";
			return true;
		}
	}
	for(auto neighbour : start_node->neighbours){
		if(contains(visited, neighbour->value))
			continue;
		else 
			visited.push_back(start_node);
		if(can_reach(neighbour, g, visited, cur_size+1))
			return true;
	}
	return false;
}

bool hamiltonian_path(graph g, graph& path){
	for(auto node : g)
		if(can_reach(node, g, path))
			return true;
	return false;
}

int main(int argc, char** argv){
	if(argc != 1)
		return -1;
	int num = atoi(argv[0]);
	
	//create the graph
	graph g;
	for(int i=0; i<num; ++i)
		g.push_back(new node(i));
	
	//connect edges
	for(auto i : g){
		for(auto j : g){
			if( is_square(i->value + j->value) ){
				i->neighbours.push_back(j);
				j->neighbours.push_back(i);
			}
		}
	}
	
	//find a path that visits each node once
	graph path;
	auto success = hamiltonian_path(g, path);
	if(!success){
		std::cout<<"failure";
		return -1;
	}
	for(auto s : g){
		std::cout << s->value + " ";
	}
	std::cout<<::std::endl;
				
}

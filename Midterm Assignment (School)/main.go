/*
Kevin Koehler
1163209
kak750
CMPT 360 - Programming part of Midterm 1
*/

package main

import (
	"fmt"
	"os"
	"bufio"
)

//checks error
func check(e error) {
    if e != nil {
        panic(e)
    }
}

//in this instance, a vertex has a weight (num phones)
type Vertex struct {
	weight int
	id     int
	neighbours []*Vertex
}

//adjacency list implementation
type Graph []*Vertex


//tests that the graph constructed is properly written
func testGraph(str []string, g Graph){
	var V, w1 int
	_, err := fmt.Sscanf(str[0], "%d", &V)
	check(err)
	_, err = fmt.Sscanf(str[1], "%d", &w1)
	check(err)
	if len(g)!=V {
		fmt.Println("test failed, graph wrong size")
	} else if g[0].weight != w1 {
		fmt.Println("test failed, vertex wrong weight")
	} else if len(g[0].neighbours) < 1 {
		fmt.Println("test failed, edges not working")
	}
}

//creates a graph from a string read from stdin
//returns a list of the leaves as well
func readInput(str []string, ) (Graph, []*Vertex) {
	var numVertices int
	var sum int
	var g Graph
	
	_, err := fmt.Sscanf(str[0], "%d", &numVertices)
	check(err)
	for i:=0; i<numVertices-1; i++ {
		var w int
	  _, err = fmt.Sscanf(str[i+1], "%d", &w)
	  check(err)
	  sum+=w
	  v := new(Vertex)
	  v.id = i
	  v.weight = w
		g = append(g, v)
	}
	v := new(Vertex)
	v.id = numVertices-1
	v.weight = numVertices - sum
	g = append(g, v)
	
	for i:=0; i<numVertices-1; i++{
		var u, v int
		_, err = fmt.Sscanf(str[i+numVertices], "%d%d", &u, &v)
		check(err)
		g[u].neighbours = append(g[u].neighbours , g[v]) 
		g[v].neighbours = append(g[v].neighbours , g[u]) 
	}
	var leaves []*Vertex
	for i:= 0; i<len(g); i++{
		if len(g[i].neighbours) == 1 {
			leaves = append(leaves, g[i])
		}
	}
	return g, leaves
}

var prevID = -1

//finds leaf in the tree/leaf tracker
func findLeaf(lt []*Vertex) *Vertex {
	return lt[0]
}

//returns the first neighbour of a leaf
func findNeighbour(v *Vertex) *Vertex {
	return v.neighbours[0]
}

//removes an index from a slice
//fast method O(1)
func RemoveIndex(s []*Vertex, i int) []*Vertex {
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}

//deletes a vertex l from another vertex n's adjacency list
func deleteFromNeighbour(n *Vertex, l *Vertex) *Vertex{
	for i:=0; i<len(n.neighbours); i++ {
		if n.neighbours[i].id == l.id {
			n.neighbours = RemoveIndex(n.neighbours, i)
		}
	}
	return n
}

//function safely removes a leaf
func removeLeaf(v *Vertex, lt []*Vertex) []*Vertex {
	lt = RemoveIndex(lt, 0)
	leafNeighbour := findNeighbour(v)
	deleteFromNeighbour(leafNeighbour, v)
	if len(leafNeighbour.neighbours) == 1 {
		lt = append(lt, leafNeighbour)
	}
	return lt
}

//abs val of int
func intAbs(i int) int{
	if i < 0 { return (-1)*i } else { return i }
}

//function containing the actual algorithm
func calculateNumShipments(lt []*Vertex) int {
	counter := 0
	for len(lt) > 1 { //nontrivial
		l := findLeaf(lt)
		m := findNeighbour(l)
		xfer := l.weight - 1
		counter = counter + intAbs(xfer)
		m.weight = m.weight + xfer
		lt = removeLeaf(l, lt)
	}
	return counter
}

//main function, entry point
//program example usage: progName < test > output
func main(){
	scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  strs := []string{}
  for scanner.Scan(){
  	strs = append(strs, scanner.Text())
  }
  g, lt := readInput(strs)
  testGraph(strs, g)
  
  v := findLeaf(lt)
  if len(v.neighbours) != 1 {
  	fmt.Println("error, leaf found with neighbours != 1, id: ", v.id, " num neighbours: ", len(v.neighbours))
  }
  m := findNeighbour(v)
  if(v.neighbours[0].id != m.id){
  	fmt.Println("error, leaf found wrong neighbour")
  }
  
  if len(lt[0].neighbours) != 1 {
  	fmt.Println("error, no leaf in non-trivial tree")
  }
  
  numShipments := calculateNumShipments(lt)
  fmt.Println(numShipments)
}

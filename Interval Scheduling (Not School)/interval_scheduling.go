package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
) 

type interval struct {
    t1 int
    t2 int
}

func check(e error){
    if(e != nil){
        panic(e)
    }
}

func readInput() []string{
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    strs := []string{}
    for scanner.Scan() {
        strs = append(strs, scanner.Text())
    }
    return strs
}

func processIntervals(strs []string) []interval {
    var intervals []interval
    var x, y int
    for i:=0; i<len(strs); i++{
        _, err :=fmt.Sscanf(strs[i], "%d%d", &x, &y)
        check(err)
        intervals = append(intervals, interval{t1: x, t2: y})
    }
    return intervals
}

func intersecting(i1 interval, i2 interval) bool {
    return (i2.t2 > i1.t1 && i2.t1 < i1.t2) || (i2.t1 < i1.t2 &&  i2.t2 > i1.t1)
}

func pruneIntersecting(ivl interval, intervals []interval) []interval {
    for i:=0; i<len(intervals); i++ {
        if intersecting(ivl, intervals[i]){
            intervals = append(intervals[:i], intervals[i+1:]...)
            i--
        }
    }
    return intervals
}

func scheduleIntervals(intervals []interval) []interval {
    var optimalList []interval   
    sort.Slice(intervals, func(i, j int) bool { return intervals[i].t2 < intervals[j].t2 })
    for len(intervals) > 0 {
        x := intervals[0]
        optimalList = append(optimalList, x)
        intervals = intervals[1:]
        intervals = pruneIntersecting(x, intervals)     
    }
    return optimalList
}

func displayIntervals(ivls []interval){
    for i:=0; i<len(ivls); i++ {
        fmt.Println("(", ivls[i].t1, ",", ivls[i].t2, ")")
    }
}

func main(){
    input := readInput()
    intervals := processIntervals(input)
    optimalList := scheduleIntervals(intervals)
    displayIntervals(optimalList)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func check(e error){
  if e != nil {
    panic(e)
  }
}

func diff(a, b int) int {
  if a < b {
    return b - a
  } else {
    return a - b
  }
}

func main(){

  dat, err := os.Open("input.txt")

  check(err)

  //fmt.Println(dat)
  //fmt.Print("Starting")

  Scanner := bufio.NewScanner(dat)
  Scanner.Split(bufio.ScanWords)
  //fmt.Println(Scanner)
  
  var leftList []int
  var rightList []int
    var totalDistance int

  check := 0
  for Scanner.Scan(){
    if (check == 0){
      number, err := strconv.Atoi(Scanner.Text())
      if (err != nil){
      fmt.Println("Error:", err)
      }
      leftList = append(leftList,number)
      check+=1
    } else {
      number, err := strconv.Atoi(Scanner.Text())
      if (err != nil){
      fmt.Println("Error:", err)
      }
      rightList = append(rightList,number)
      check-=1
    }
    //fmt.Println(Scanner.Text())
  }

  slices.Sort(leftList)
  slices.Sort(rightList)
 
  fmt.Printf("Left %v, Right: %v\n",len(leftList), len(rightList))

  for i := range leftList{
    totalDistance += diff(leftList[i],rightList[i])
  }

  fmt.Println("Total Distance:",totalDistance)
  
  // DAY #2
  // SIMILARITY SCORE


  var dictSimilarity map[int]int
  var totalSimilarity int = 0

  dictSimilarity = make(map[int]int)


  for i,v := range leftList{
      founds := 0
    for j := range rightList{
      if (rightList[j] == leftList[i]){
        founds+=1
      }
    }
    dictSimilarity[v] = founds
    totalSimilarity+=v * founds
  }

  //fmt.Println(dictSimilarity)


  fmt.Println("Total Similarity:", totalSimilarity)





}

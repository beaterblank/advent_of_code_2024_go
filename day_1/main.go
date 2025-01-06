package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func check(e error){
  if e!=nil{
    panic(e)
  }
}


func sorted_insert_int(arr []int64,item int64) []int64{
  l := len(arr)
  arr = append(arr,item)
  for l-1 >= 0 && arr[l-1] > item {
    arr[l] = arr[l-1]
    l -= 1
  }
  arr[l] = item
  return arr
}

func day_1_part_1(input_dir string)(diff uint32){
  input_bytes, err := os.ReadFile(input_dir)
  check(err)
  
  input := string(input_bytes)
  lines := strings.Split(input,"\n")
  length := len(lines)
  // ignore last line because its a blank
  lines = lines[:length-1]
  length -= 1

  left_arr := make([]int64,0,length)
  right_arr := make([]int64,0,length)

  for i:=0; i<length; i++{
    items := strings.SplitN(lines[i],"   ",2)
    left,err_left := strconv.ParseInt(items[0],10,64)
    check(err_left)
    right,err_right := strconv.ParseInt(items[1],10,64)
    check(err_right)
    left_arr = sorted_insert_int(left_arr,left)
    right_arr = sorted_insert_int(right_arr,right)
  }

  for i:=0; i<length; i++{
    if (left_arr[i]>right_arr[i]){
      diff += uint32(left_arr[i]-right_arr[i])
    } else {
      diff += uint32(right_arr[i]-left_arr[i])
    }
  }

  return
}


func day_1_part_2(input_dir string)(diff uint32){
  input_bytes, err := os.ReadFile(input_dir)
  check(err)
  
  input := string(input_bytes)
  lines := strings.Split(input,"\n")
  length := len(lines)
  // ignore last line because its a blank
  lines = lines[:length-1]
  length -= 1

  left_arr := make([]int64,0,length)
  right_map := make(map[int64]int64)

  for i:=0; i<length; i++{
    items := strings.SplitN(lines[i],"   ",2)
    left,err_left := strconv.ParseInt(items[0],10,64)
    check(err_left)
    right,err_right := strconv.ParseInt(items[1],10,64)
    check(err_right)
    left_arr =  append(left_arr,left)
    right_map[right] += 1
  }

  for i:=0; i<length; i++{
    diff += uint32(left_arr[i]*right_map[left_arr[i]])
  }

  return
}


func main(){
  diff := day_1_part_1("./input.txt")
  fmt.Println("sorted difference is : ",diff)
  diff = day_1_part_2("./input.txt")
  fmt.Println("similarity  score is : ",diff)
}

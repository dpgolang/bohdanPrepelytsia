package main

import (
  "os/exec"
  "fmt"
  "log"
  "os"  
  "strings"
  "strconv"
)

func main() {

  Start:
  fmt.Print("Enter size: ")
  inputString :=""
  fmt.Scanln(&inputString)
  if count, err := strconv.Atoi(inputString); err != nil {
    fmt.Printf("Wrong input info!\n")
    goto Start
} 
  height,width := GetSize()
  OutPut(height,width,count)
  
}

func GetSize() (int, int){
  cmd := exec.Command("stty", "size")
  cmd.Stdin = os.Stdin
  out, err := cmd.Output()  
  if err != nil {
    log.Fatal(err)
  }  
  return ConvertSizeToInt(string(out))
}

func ConvertSizeToInt(sizeOfString string) (int,int){
  tmp := sizeOfString[:len(sizeOfString)-1]
  fmt.Println(tmp)
  str := strings.Split(tmp, " ")
  height, err := strconv.Atoi(string(str[0]))
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
  width, err := strconv.Atoi(string(str[1]))
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
  return height,width
}

func OutPut(height,width,count int){
  //var outPutString string
if count > height || count + 1 > width {
    fmt.Println("Sorry, but size of this window cannot allow to write the desk")
  } else {
    for i:= 0; i < count; i++ {
      if i%2 != 0{
      fmt.Printf(" ")
    }
      for j:= 0;j < count; j++ {
          fmt.Printf("* ")
      }
      fmt.Printf("\n")
    }
  }
}
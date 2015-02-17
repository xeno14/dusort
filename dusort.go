package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

type Directory struct {
  Name string
  Size string
}

func SizeToFloat64(sizeStr string) float64 {
  var sub string = sizeStr[:len(sizeStr)-1]
  size, err := strconv.ParseFloat(sub, 64)
  if err != nil {
    fmt.Println(err)
  }
  const (
    K = 1000.0
  )
  switch {
    case strings.HasSuffix(sizeStr, "K"):
      size *= K
    case strings.HasSuffix(sizeStr, "M"):
      size *= K*K
    case strings.HasSuffix(sizeStr, "M"):
      size *= K*K*K
    case strings.HasSuffix(sizeStr, "G"):
      size *= K*K*K*K
    case strings.HasSuffix(sizeStr, "T"):
      size *= K*K*K*K*K
  }
  return size
}

func NewDirectory(name string, sizeStr string) Directory {
  return Directory{name, sizeStr}
}

type Directories []Directory

func (d Directories) Len() int {
  return len(d)
}

func (d Directories) Swap(i, j int) {
  d[i], d[j] = d[j], d[i]
}

func (d Directories) Less(i, j int) bool {
  return SizeToFloat64(d[i].Size) < SizeToFloat64(d[j].Size)
}

func ReadLines() Directories {
  directories := make(Directories, 0, 256)
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    line := scanner.Text()
    splited := strings.Split(line, "\t")
    directories = append(directories,
                         NewDirectory(strings.TrimSpace(splited[1]),
                                      strings.TrimSpace(splited[0])))
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Error at reading stdin:", err)
  }
  return directories
}

func main() {
  directories := ReadLines()

  sort.Sort(sort.Reverse(directories))

  for _, d := range directories {
    fmt.Println(d.Size + "\t" + d.Name)
  }
}

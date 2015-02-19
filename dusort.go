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
    fmt.Fprintln(os.Stderr, "Error at parse float:", err)
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

func ReadStdin() <-chan string {
  ch := make(chan string)
  scanner := bufio.NewScanner(os.Stdin)
  go func() {
    for scanner.Scan() {
      ch <- scanner.Text()
    }
    if err := scanner.Err(); err != nil {
      fmt.Fprintln(os.Stderr, "Error at reading stdin:", err)
    }
    close(ch)
  }()
  return ch
}

func main() {
  dirs := make(Directories, 0, 1024)
  for line := range ReadStdin() {
    splited := strings.Split(line, "\t")
    dirs = append(dirs, NewDirectory(strings.TrimSpace(splited[1]),
                                     strings.TrimSpace(splited[0])))
  }
  sort.Sort(sort.Reverse(dirs))

  for _, d := range dirs {
    fmt.Println(d.Size + "\t" + d.Name)
  }
}

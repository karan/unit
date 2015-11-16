package main

import (
  "fmt"
  "io/ioutil"
  "go/parser"
  "go/token"
  "os"
  "strings"
)


var UNITS_DIR = "units/"
var QUALIFIER_DOC = "- unit-dep"
var MARTINI_PATH = "github.com/go-martini/martini"


type DepSet struct {
  set map[string]bool
}

func NewDepSet() *DepSet {
  return &DepSet{make(map[string]bool)}
}

func (set *DepSet) Add(s string) bool {
  _, found := set.set[s]
  set.set[s] = true
  return !found
}

func (set *DepSet) Len() int {
  return len(set.set)
}

func (set *DepSet) Keys() []string {
  keys := make([]string, 0, set.Len())
  for k := range set.set {
    keys = append(keys, k)
  }
  return keys
}

func getUnitsFiles() []os.FileInfo {
  files, _ := ioutil.ReadDir(UNITS_DIR)
  return files
}

func GetDeps() []string {
  units_files := getUnitsFiles()
  set := NewDepSet()

  for _, unit := range units_files {
    fset := token.NewFileSet()

    f, err := parser.ParseFile(fset, UNITS_DIR + unit.Name(), nil, parser.ParseComments)
    if err != nil {
      fmt.Println(err)
      return nil
    }

    for _, s := range f.Imports {
      comment := strings.Trim(s.Comment.Text(), "\n")
      path := strings.Trim(s.Path.Value, "\"")

      if comment == QUALIFIER_DOC && path != MARTINI_PATH {
        set.Add(path)
      }
    }

  }

  return set.Keys()
}

func main() {
  deps := GetDeps()
  fmt.Println(strings.Join(deps, " "))
}

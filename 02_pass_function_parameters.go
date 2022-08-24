package main

import (
  "fmt"
  "sort"
)

type spare_part struct {
  part_name string
  part_number int
  drawing_number int
  active bool
}

func main() {
spare_parts := []spare_part {
  {"bolt", 12902, 3374, true},
  {"washer", 76432, 1263, true},
  {"nut", 38819, 4873, true},
  {"wrench", 13451, 5397, false},
  }
  
  fmt.Println("part name;","part number;","drawing number")
  fmt.Println(spare_parts)
  fmt.Println("")
  
  // sort by drawing number
  sort.Slice(spare_parts, func(a int, b int)bool {
    return spare_parts[a].drawing_number < spare_parts[b].drawing_number
  })
  fmt.Println("Sorted by drawing number:")
  fmt.Println(spare_parts)
  fmt.Println("")
  
  // sort by part name
  sort.Slice(spare_parts, func(a int, b int)bool {
    return spare_parts[a].part_name < spare_parts[b].part_name
  })
  fmt.Println("Sorted by part name:")
  fmt.Println(spare_parts)
  fmt.Println("")
  
  // sort by active product
  sort.Slice(spare_parts, func(a int, b int)bool {
    return spare_parts[a].active != spare_parts[b].active
  })
  fmt.Println("Sorted by inactive/active product:")
  fmt.Println(spare_parts) 
}

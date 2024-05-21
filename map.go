package piscine

func Map(f func(int) bool, a []int) []bool {
  var bools []bool = make([]bool, len(a))

  for i, v := range a {
    bools[i] = f(v)
  }
  return bools
}

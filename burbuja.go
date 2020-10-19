package main

func Burbuja(s []int64)  {
  len_s := len(s)
  for i := 1; i < len_s; i++ {
    for j := 0; j < len_s - i; j++ {
      if s[j+1] < s[j] {
        aux := s[j]
        s[j] = s[j+1]
        s[j+1] = aux
      }
    }
  }
}

func main()  {

}

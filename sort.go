package main

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[j] > s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func SelectSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		max := s[i]
		k := i
		for j := i; j < len(s); j++ {
			if s[j] > max {
				max = s[j]
				k = j
			}
		}
		s[i], s[k] = s[k], s[i]
	}
}

func QuickSort(s []int) {

}

func Add() {
}

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			go Add()
		}(i)
	}
}
	//s := []int{4, 2, 1, 5, 5, 7, 3}
	//
	////BubbleSort(s)
	//SelectSort(s)
	//fmt.Println(s)

//}

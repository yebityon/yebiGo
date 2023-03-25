package main

func Create2DVec(N int, M int, e int ) [][]int {
	ret := make([][]int, N)
	for i := range ret {
		ret[i] = make([]int, M)
		for j  := range ret[i] {
			ret[i][j] = e
		}		
	}
	return ret
}

func DeepCopy( a  *[]int , b *[]int )  {
	for _, val := range *a {
		*b = append(*b, val)
	}
}



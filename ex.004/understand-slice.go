package main

import "fmt"

func main() {
	fmt.Printf("sliceAppend\n")
	sliceAppend()
	fmt.Printf("\nsliceAssign\n")
	sliceAssign()
}

func sliceAssign() {
	var (
		aSlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		bSlice []int
		cSlice []int
	)
	fmt.Printf("aSlice: %p, len = %d, cap = %d\n", aSlice, len(aSlice), cap(aSlice))
	fmt.Printf("aSlice[0]: %p %d\n", &aSlice[0], aSlice[0])
	fmt.Printf("aSlice[1]: %p %d\n", &aSlice[1], aSlice[1])
	fmt.Printf("aSlice[2]: %p %d\n", &aSlice[2], aSlice[2])
	fmt.Printf("aSlice[3]: %p %d\n", &aSlice[3], aSlice[3])
	fmt.Printf("aSlice[4]: %p %d\n", &aSlice[4], aSlice[4])
	fmt.Printf("aSlice[5]: %p %d\n", &aSlice[5], aSlice[5])
	fmt.Printf("aSlice[6]: %p %d\n", &aSlice[6], aSlice[6])
	fmt.Printf("aSlice[7]: %p %d\n", &aSlice[7], aSlice[7])
	fmt.Printf("aSlice[8]: %p %d\n", &aSlice[8], aSlice[8])
	fmt.Printf("aSlice[9]: %p %d\n", &aSlice[9], aSlice[9])

	bSlice = aSlice[3:5]
	fmt.Printf("bSlice: %p, len = %d, cap = %d\n", bSlice, len(bSlice), cap(bSlice))
	fmt.Printf("bSlice[0]: %p %d\n", &bSlice[0], bSlice[0])
	fmt.Printf("bSlice[1]: %p %d\n", &bSlice[1], bSlice[1])

	cSlice = aSlice[2:6]
	fmt.Printf("cSlice: %p, len = %d, cap = %d\n", cSlice, len(cSlice), cap(cSlice))
	fmt.Printf("cSlice[0]: %p %d\n", &cSlice[0], cSlice[0])
	fmt.Printf("cSlice[1]: %p %d\n", &cSlice[1], cSlice[1])
	fmt.Printf("cSlice[2]: %p %d\n", &cSlice[2], cSlice[2])
	fmt.Printf("cSlice[3]: %p %d\n", &cSlice[3], cSlice[3])

	bSlice[0] = 100
	fmt.Printf("bSlice[0] = 100\n")

	fmt.Printf("aSlice: %p, len = %d, cap = %d\n", aSlice, len(aSlice), cap(aSlice))
	fmt.Printf("aSlice[0]: %p %d\n", &aSlice[0], aSlice[0])
	fmt.Printf("aSlice[1]: %p %d\n", &aSlice[1], aSlice[1])
	fmt.Printf("aSlice[2]: %p %d\n", &aSlice[2], aSlice[2])
	fmt.Printf("aSlice[3]: %p %d\n", &aSlice[3], aSlice[3])
	fmt.Printf("aSlice[4]: %p %d\n", &aSlice[4], aSlice[4])
	fmt.Printf("aSlice[5]: %p %d\n", &aSlice[5], aSlice[5])
	fmt.Printf("aSlice[6]: %p %d\n", &aSlice[6], aSlice[6])
	fmt.Printf("aSlice[7]: %p %d\n", &aSlice[7], aSlice[7])
	fmt.Printf("aSlice[8]: %p %d\n", &aSlice[8], aSlice[8])
	fmt.Printf("aSlice[9]: %p %d\n", &aSlice[9], aSlice[9])

	bSlice = aSlice[3:5]
	fmt.Printf("bSlice: %p, len = %d, cap = %d\n", bSlice, len(bSlice), cap(bSlice))
	fmt.Printf("bSlice[0]: %p %d\n", &bSlice[0], bSlice[0])
	fmt.Printf("bSlice[1]: %p %d\n", &bSlice[1], bSlice[1])

	cSlice = aSlice[2:6]
	fmt.Printf("cSlice: %p, len = %d, cap = %d\n", cSlice, len(cSlice), cap(cSlice))
	fmt.Printf("cSlice[0]: %p %d\n", &cSlice[0], cSlice[0])
	fmt.Printf("cSlice[1]: %p %d\n", &cSlice[1], cSlice[1])
	fmt.Printf("cSlice[2]: %p %d\n", &cSlice[2], cSlice[2])
	fmt.Printf("cSlice[3]: %p %d\n", &cSlice[3], cSlice[3])

}

func sliceAppend() {
	var (
		aSlice []int = []int{1, 2, 3, 4}
		bSlice []int = make([]int, 3)
		cSlice []int = make([]int, 3, 5)
	)

	fmt.Printf("aSlice: %p, len = %d, cap = %d\n", aSlice, len(aSlice), cap(aSlice))
	fmt.Printf("aSlice[0]: %p\n", &aSlice[0])
	fmt.Printf("aSlice[1]: %p\n", &aSlice[1])
	fmt.Printf("aSlice[2]: %p\n", &aSlice[2])
	fmt.Printf("aSlice[3]: %p\n", &aSlice[3])
	aSlice = append(aSlice, 5, 6)
	fmt.Printf("After Append 5, 6. aSlice: %p, len = %d, cap = %d\n", aSlice, len(aSlice), cap(aSlice))
	fmt.Printf("aSlice[0]: %p\n", &aSlice[0])
	fmt.Printf("aSlice[1]: %p\n", &aSlice[1])
	fmt.Printf("aSlice[2]: %p\n", &aSlice[2])
	fmt.Printf("aSlice[3]: %p\n", &aSlice[3])
	fmt.Printf("aSlice[4]: %p\n", &aSlice[4])
	fmt.Printf("aSlice[5]: %p\n", &aSlice[5])

	fmt.Printf("\nbSlice: %p, len = %d, cap = %d\n", bSlice, len(bSlice), cap(bSlice))
	fmt.Printf("bSlice[0]: %p\n", &bSlice[0])
	fmt.Printf("bSlice[1]: %p\n", &bSlice[1])
	fmt.Printf("bSlice[2]: %p\n", &bSlice[2])
	bSlice = append(bSlice, 4, 5)
	fmt.Printf("After Append 4, 5. bSlice: %p, len = %d, cap = %d\n", bSlice, len(bSlice), cap(bSlice))
	fmt.Printf("bSlice[0]: %p\n", &bSlice[0])
	fmt.Printf("bSlice[1]: %p\n", &bSlice[1])
	fmt.Printf("bSlice[2]: %p\n", &bSlice[2])
	fmt.Printf("bSlice[3]: %p\n", &bSlice[3])
	fmt.Printf("bSlice[4]: %p\n", &bSlice[4])

	fmt.Printf("\ncSlice: %p, len = %d, cap = %d\n", cSlice, len(cSlice), cap(cSlice))
	fmt.Printf("cSlice[0]: %p\n", &cSlice[0])
	fmt.Printf("cSlice[1]: %p\n", &cSlice[1])
	fmt.Printf("cSlice[2]: %p\n", &cSlice[2])
	cSlice = append(cSlice, 4, 5)
	fmt.Printf("After Append 4, 5. cSlice: %p, len = %d, cap = %d\n", cSlice, len(cSlice), cap(cSlice))
	fmt.Printf("cSlice[0]: %p\n", &cSlice[0])
	fmt.Printf("cSlice[1]: %p\n", &cSlice[1])
	fmt.Printf("cSlice[2]: %p\n", &cSlice[2])
	fmt.Printf("cSlice[3]: %p\n", &cSlice[3])
	fmt.Printf("cSlice[4]: %p\n", &cSlice[4])
	cSlice = append(cSlice, 6, 7)
	fmt.Printf("After Append 6, 7. cSlice: %p, len = %d, cap = %d\n", cSlice, len(cSlice), cap(cSlice))
	fmt.Printf("cSlice[0]: %p\n", &cSlice[0])
	fmt.Printf("cSlice[1]: %p\n", &cSlice[1])
	fmt.Printf("cSlice[2]: %p\n", &cSlice[2])
	fmt.Printf("cSlice[3]: %p\n", &cSlice[3])
	fmt.Printf("cSlice[4]: %p\n", &cSlice[4])
	fmt.Printf("cSlice[5]: %p\n", &cSlice[5])
	fmt.Printf("cSlice[6]: %p\n", &cSlice[6])

}
func hammingWeight(num uint32) (ans int) {
	for num != 0 {
		num -= num & -num
		ans++
	}
	return
}
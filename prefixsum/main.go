package prefixsum

func prefixSum(input []int) []int {
	output := make([]int, 0)

	prev := input[0]
	output = append(output, prev)

	for i := 1; i < len(input); i++ {
		prev += input[i]
		output = append(output, prev)
	}

	return output
}

func rangePrefixSum(prefixSum []int, l, r int) int {
	if l == 0 {
		return prefixSum[r]
	}

	return prefixSum[r] - prefixSum[l-1]
}

func main() {

}

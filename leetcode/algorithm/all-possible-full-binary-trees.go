package main

func FBT(n int, _memorize map[int][]*TreeNode) []*TreeNode {

	memorize := make(map[int][]*TreeNode)
	if memorize[n] != nil {
		return memorize[n]
	}
	if n%2 == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{{Val: 0}}
	}
	result := make([]*TreeNode, 0)
	for i := 1; i < n; i += 2 {
		if _memorize[i] == nil {
			_memorize[i] = FBT(i, _memorize)
		}
		left := _memorize[i]
		if _memorize[n-i-1] == nil {
			_memorize[n-i-1] = FBT(n-i-1, _memorize)
		}
		right := _memorize[n-i-1]
		for _, l := range left {
			for _, r := range right {
				result = append(result, &TreeNode{Val: 0, Left: l, Right: r})
			}
		}
	}
	return result
}

func allPossibleFBT(n int) []*TreeNode {
	_memorize := make(map[int][]*TreeNode)
	_memorize[1] = []*TreeNode{{Val: 0}}
	return FBT(n, _memorize)
}

func main() {
	allPossibleFBT(7)
}

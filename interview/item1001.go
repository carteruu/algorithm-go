package interview

func merge(A []int, m int, B []int, n int) {
	idxA := m - 1
	idxB := n - 1
	for i := m + n - 1; i >= 0; i-- {
		if idxB < 0 || (idxA >= 0 && A[idxA] > B[idxB]) {
			A[i] = A[idxA]
			idxA--
		} else {
			A[i] = B[idxB]
			idxB--
		}
	}
}

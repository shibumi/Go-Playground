package T9_permutator

var T9 = map[int]string{
	1: "",
	2: "ABC",
	3: "DEF",
	4: "GHI",
	5: "JKL",
	6: "MNO",
	7: "PQRS",
	8: "TUV",
	9: "WXYZ",
	0: "+"}

func T9Permutator(numberblock []int) (result []string) {
	return result
}

func main() {
	T9Permutator([]int{2, 2})
}

//TODO Heap's Algorithm
////i acts similarly to the stack pointer
//i := 0;
//while i < n do
//if  c[i] < i then
//if i is even then
//swap(A[0], A[i])
//else
//swap(A[c[i]], A[i])
//end if
//output(A)
////Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
//c[i] += 1
////Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
//i := 0
//else
////Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
//c[i] := 0
//i += 1
//end if
//end while

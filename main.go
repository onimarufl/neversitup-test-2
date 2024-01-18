package main

import "fmt"

func main() {
	slice_of_slices := make([]string, 4)
	slice_of_slices[0] = "a"
	slice_of_slices[1] = "ab"
	slice_of_slices[2] = "abc"
	slice_of_slices[3] = "aabb"
	for _, str := range slice_of_slices {
		resp := test2(str)

		fmt.Println(fmt.Sprintf("With input %s : return %s", str, resp))
	}
}

func test2(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	strSlice := []string{}
	for i, c := range str {

		subStr := str[:i] + str[i+1:]
		permutation := test2(subStr)

		for _, p := range permutation {
			strSlice = append(strSlice, string(c)+string(p))
		}
	}

	resp := removeDuplicateStr(strSlice)
	return resp
}

func removeDuplicateStr(strSlice []string) []string {
	keys := make(map[string]bool)
	slice := []string{}
	for _, item := range strSlice {
		if _, value := keys[item]; !value {
			keys[item] = true
			slice = append(slice, item)
		}
	}
	return slice
}

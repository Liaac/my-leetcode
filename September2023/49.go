package september2023

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// 思路：若字符串包含字母种类和每个字母个数均相同 则为题意的字母异位词
// 考虑使用map将字符串转换为字符到个数的映射 相同的放入同一集合中

// 以下通过111/119 时间超时
// func groupAnagrams(strs []string) [][]string {
// 	res := make([][]string, 0)
// 	mappings := map[string]map[int]int{}
// 	for _, str := range strs {
// 		mapping := transToMap(str)
// 		mappings[str] = mapping
// 		if len(res) == 0 {
// 			temp := []string{}
// 			temp = append(temp, str)
// 			res = append(res, temp)
// 		} else {
// 			flag := false
// 			for i, strs := range res {
// 				if isEqual(mappings[strs[0]], mapping) {
// 					res[i] = append(res[i], str)
// 					flag = true
// 				}
// 			}
// 			if !flag {
// 				temp := []string{}
// 				temp = append(temp, str)
// 				res = append(res, temp)
// 			}
// 		}
// 	}
// 	return res
// }

// func transToMap(str string) map[int]int {
// 	res := make(map[int]int, 0)
// 	for _, s := range str {
// 		sInt := int(s)
// 		_, ok := res[sInt]
// 		if !ok {
// 			res[sInt] = 0
// 		}
// 		res[sInt]++
// 	}
// 	return res
// }

// func isEqual(a, b map[int]int) bool {
// 	if !(len(a) == len(b)) {
// 		return false
// 	}
// 	for ka, va := range a {
// 		vb, ok := b[ka]
// 		if !ok || vb != va {
// 			return false
// 		}
// 	}
// 	return true
// }

// 官方题解使用计数法 Go可以将数组作为map的key
// 每个str对应一个[26]int{}的数组，作为代表这个str的计数
func groupAnagrams(strs []string) [][]string {
	mappings := map[[26]int][]string{} // 每个计数相同的数组对应的所有str []string
	for _, str := range strs {
		cnt := [26]int{}
		for _, v := range str {
			cnt[v-'a']++
		}
		mappings[cnt] = append(mappings[cnt], str)
	}
	res := make([][]string, 0)
	for _, curStrs := range mappings {
		res = append(res, curStrs)
	}
	return res
}

package Week_02

// 暴力思路1：
//     map[string] ----> [string, string]
// 遍历数组，将每一个字符串排序后的串作为key，然后添加到key 对应数组
// 最后将map的值加入到数组中就是结果
//
//
// ["eat","tea","tan","ate","nat","bat"]
//
// map["aet"] = ["eat", "tea", "ate"]
// map["ant"] = ["tan", "nat"]
// map["abt"] = ["bat"]
//
// out = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
//
// 排序 log(n), 那么时间复杂度就是 nlog(k), k 表示字母的长度

// 网友思路：
//
// 同上面思路大致一致，只不过key的算法是 字符串字符的和 + 字符串字符的乘机
//
func groupAnagrams(strs []string) [][]string {
	out := make([][]string, 0, 3)
	m := make(map[uint64][]string)

	for _, str := range strs {
		var sum uint64 = 0
		var multi uint64 = 1

		for _, v := range str {
			sum += uint64(v)
			multi *= uint64(v)
		}

		key := sum + multi

		m[key] = append(m[key], str)
	}

	for _, v := range m {
		out = append(out, v)
	}

	return out
}

package Week_02

//
//
// 思路：第一个数组的每一个元素存储到map中，并计数
//       遍历第二个map，若指定字符不存在则返回false
//       若存在计数-1, 当某个字符的计数<=0是删除key
//       最后判定map长度是否为0，为0，则返回true，否则
//       false
//
func isAnagram(s string, t string) bool {
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; !ok {
			return false
		} else {
			m[t[i]] -= 1
			if m[t[i]] <= 0 {
				delete(m, t[i])
			}
		}
	}

	if len(m) == 0 {
		return true
	}

	return false
}

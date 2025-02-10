package main

func findAnagrams(s string, p string) (ans []int) {
	m, n := len(s), len(p)
	if m < n {
		return ans
	}
	var cntsP, cnts [26]int
	for i := 0; i < n; i++ {
		cntsP[p[i]-byte('a')]++
	}
	for i := 0; i < m; i++ {
		cnts[s[i]-byte('a')]++
		if i >= n-1 {
			if cnts == cntsP {
				ans = append(ans, i-n+1)
			}
			cnts[s[i-n+1]-byte('a')]--
		}
	}
	return ans

}
func findAnagrams2(s string, p string) []int {
	// 特殊情况：当字符串s长度小于字符串p，则肯定不满足字母异位词，返回空
	if len(s) < len(p) {
		return []int{}
	}

	// 初始化变量
	// 初始化结果集
	res := make([]int, 0)
	// 初始化一个count数组
	// 数组长度为26，记录滑动窗口中字符与字符串p中字符的差异
	// 每个索引分别代表一个字母（如0代表'a',1代表'b'，依次类推）
	// 某个索引处的数字符号分别表示：
	// 	正数：表示滑动窗口内该字母的个数比字符串p的多
	// 	0：表示滑动窗口内该字母的个数与字符串的一样
	// 	负数：表示滑动窗口内该字母的个数比字符串p的少
	count := [26]int{}

	// 初始化形成滑动窗口
	// 在s中设置一个起点在索引0处，长度为m的滑动窗口，并初始化统计两个字符串中有差异的字母种类和数量
	for i := 0; i < len(p); i++ {
		count[s[i]-'a']++ // 在s中的字母就+1
		count[p[i]-'a']-- // 在p中的字母就-1
	}

	// 初始化differ，记录当前窗口中字符串与字符串p中有差异的字母种类数
	differ := 0
	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	// 当两个字符串有差异的字母种类为0时，说明两个字符串为字母异位词，就将索引0加入结果集
	if differ == 0 {
		res = append(res, 0)
	}

	// 持续向右滑动并比较
	left := 1
	right := len(p)
	for right < len(s) {
		// 左指针右移
		if count[s[left-1]-'a'] == 1 {
			// 如果移动前，左指针处字母的count值为1，左指针右移（移除左指针元素）会导致该字母处的count变成0，说明该字母不再存在差异，differ-1
			differ--
		} else if count[s[left-1]-'a'] == 0 {
			// 如果移动前，左指针处字母的count值为0，左指针右移（移除左指针元素）会导致该字母处的count变成-1，说明滑动窗口中该字母比字符串p中少了一个，此时该字母产生差异，differ+1
			differ++
		}
		// 将滑动窗口移动前的左指针元素从count中移除
		count[s[left-1]-'a']--

		// 右指针右移
		if count[s[right]-'a'] == -1 {
			// 如果移动后，右指针处字母的count值为-1，右指针右移（新增右指针元素）会导致该字母处的count变成0，说明该字母不再存在差异，differ-1
			differ--
		} else if count[s[right]-'a'] == 0 {
			// 如果移动后，右指针处字母的count值为0，右指针右移（新增右指针元素）会导致该字母处的count变成1，说明滑动窗口中该字母比字符串p中多了一个，此时该字母产生差异，differ+1
			differ++
		}
		// 将滑动窗口移动后的右指针元素添加到count中
		count[s[right]-'a']++

		// 当两个字符串有差异的字母种类为0时，说明两个字符串为字母异位词
		if differ == 0 {
			res = append(res, left)
		}

		// 滑动窗口右移，寻找下一个结果
		left++
		right++
	}

	// 返回结果集
	return res
}

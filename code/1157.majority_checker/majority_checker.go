package code

/*
1157. 子数组中占绝大多数的元素
设计一个数据结构，有效地找到给定子数组的 多数元素 。

子数组的 多数元素 是在子数组中出现 threshold 次数或次数以上的元素。

实现 MajorityChecker 类:

MajorityChecker(int[] arr) 会用给定的数组 arr 对 MajorityChecker 初始化。
int query(int left, int right, int threshold) 返回子数组中的元素  arr[left...right] 至少出现 threshold 次数，如果不存在这样的元素则返回 -1。


示例 1：

输入:
["MajorityChecker", "query", "query", "query"]
[[[1, 1, 2, 2, 1, 1]], [0, 5, 4], [0, 3, 3], [2, 3, 2]]
输出：
[null, 1, -1, 2]

解释：
MajorityChecker majorityChecker = new MajorityChecker([1,1,2,2,1,1]);
majorityChecker.query(0,5,4); // 返回 1
majorityChecker.query(0,3,3); // 返回 -1
majorityChecker.query(2,3,2); // 返回 2
*/

type MajorityChecker struct {
	QueryMap map[int]map[int]int
}

func Constructor(arr []int) MajorityChecker {
	var mc MajorityChecker
	mc.QueryMap = make(map[int]map[int]int)
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			mc.QueryMap[i] = make(map[int]int)
		} else {
			mc.QueryMap[i] = cloneTags(mc.QueryMap[i-1])
		}
		mc.QueryMap[i][arr[i]]++
	}
	return mc
}

func cloneTags(tags map[int]int) map[int]int {
	cloneTags := make(map[int]int)
	for k, v := range tags {
		cloneTags[k] = v
	}
	return cloneTags
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	for key, value := range this.QueryMap[right] {
		if value-this.QueryMap[left-1][key] >= threshold {
			return key
		}
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */

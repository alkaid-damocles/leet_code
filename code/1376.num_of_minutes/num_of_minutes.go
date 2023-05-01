package code

/*
1376. 通知所有员工所需的时间
公司里有 n 名员工，每个员工的 ID 都是独一无二的，编号从 0 到 n - 1。公司的总负责人通过 headID 进行标识。

在 manager 数组中，每个员工都有一个直属负责人，其中 manager[i] 是第 i 名员工的直属负责人。对于总负责人，manager[headID] = -1。题目保证从属关系可以用树结构显示。

公司总负责人想要向公司所有员工通告一条紧急消息。他将会首先通知他的直属下属们，然后由这些下属通知他们的下属，直到所有的员工都得知这条紧急消息。

第 i 名员工需要 informTime[i] 分钟来通知它的所有直属下属（也就是说在 informTime[i] 分钟后，他的所有直属下属都可以开始传播这一消息）。

返回通知所有员工这一紧急消息所需要的 分钟数 。
*/
func numOfMinutes(n, headID int, manager, informTime []int) (ans int) {
	g := make([][]int, n)
	for i, m := range manager {
		if m >= 0 {
			g[m] = append(g[m], i) // 建树
		}
	}
	var dfs func(int) int
	dfs = func(x int) (maxPathSum int) {
		for _, y := range g[x] { // 遍历 x 的儿子 y（如果没有儿子就不会进入循环）
			maxPathSum = max(maxPathSum, dfs(y))
		}
		// 这和 104 题代码中的 return max(lDepth, rDepth) + 1 是一个意思
		return maxPathSum + informTime[x]
	}
	return dfs(headID) // 从根节点 headID 开始递归
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

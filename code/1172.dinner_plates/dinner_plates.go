package code

/*
1172. 餐盘栈
我们把无限数量 ∞ 的栈排成一行，按从左到右的次序从 0 开始编号。每个栈的的最大容量 capacity 都相同。

实现一个叫「餐盘」的类 DinnerPlates：

DinnerPlates(int capacity) - 给出栈的最大容量 capacity。
void push(int val) - 将给出的正整数 val 推入 从左往右第一个 没有满的栈。
int pop() - 返回 从右往左第一个 非空栈顶部的值，并将其从栈中删除；如果所有的栈都是空的，请返回 -1。
int popAtStack(int index) - 返回编号 index 的栈顶部的值，并将其从栈中删除；如果编号 index 的栈是空的，请返回 -1。
  - Your DinnerPlates object will be instantiated and called as such:
  - obj := Constructor(capacity);
  - obj.Push(val);
  - param_2 := obj.Pop();
  - param_3 := obj.PopAtStack(index);
*/
type DinnerPlates struct {
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{}
}

func (this *DinnerPlates) Push(val int) {

}

func (this *DinnerPlates) Pop() int {
	return 0
}

func (this *DinnerPlates) PopAtStack(index int) int {
	return 0
}

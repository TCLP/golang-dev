package bit

/*
	题目：uint32数的二进制位的反转
	方法：取出uint32的每一个二进制位，然后左移对应的位数之后于结果相加
*/
func BitReverse(nums uint32) uint32 {
	var ans uint32
	for i := 0; nums != 0; i++ {
		ans |= (nums % 2) << uint32(31-i)
		nums /= 2
	}
	return ans
}

/*
	题目：非空整数数组，除了某个元素只出现一次之外，其他元素均出现两次
	方法：异或运算，两个相同的数进行异或运算得到的结果是0
*/
func SingleNumber(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		nums[i+1] = nums[i] ^ nums[i+1]
	}
	return nums[len(nums)-1]
}

/*
	题目：输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数，也被称为汉明重量
	方法：依次取出其二进制位，统计为1的个数
*/
func HangmingWeight(nums uint32) uint32 {
	var count uint32
	for ; nums != 0; nums /= 2 {
		if nums%2 == 1 {
			count++
		}

	}
	return count
}

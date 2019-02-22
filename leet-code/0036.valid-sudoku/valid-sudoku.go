package problem0036

/*
Topic:
有效的数独
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
1.数字 1-9 在每一行只能出现一次。
2.数字 1-9 在每一列只能出现一次。
3.数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

数独部分空格内已填入了数字，空白格用 '.' 表示

Examlple:
实例1：
输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true

实例2：
输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。

Thought:
依次检查这三个方面有没有重复
 */

func isValidSudoku(board [][]byte) bool {
	// 记录某行，某位数字是否已经被摆放
	var row [9][10]bool
	// 记录某列，某位数字是否已经被摆放
	var col [9][10]bool
	// 记录某 3x3 宫格内，某位数字是否已经被摆放
	var block [9][10]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				num := int(board[i][j])
				if row[i][num] || col[j][num] || block[i/3*3+j/3][num] {
					return false
				} else {
					row[i][num] = true
					col[j][num] = true
					block[i/3*3+j/3][num] = true
				}
			}
		}
	}
	return true
}

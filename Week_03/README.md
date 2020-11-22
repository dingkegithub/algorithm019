学习笔记

> 1、递归、分治、回溯本质

解决重复问题


> 2、重复性

- 最近重复：分治，回溯解决

- 最优重复：动态规划


> 3、问题解决思路

- 大问题分解成子问题

- 处理子问题

- 组合子问题结果，汇总得到最终结果


> 4、递归

切记杜绝人肉递归，正确做法找最近重复子问题，对子问题进行归纳，对子问题进行简单编码，然后找到边界条件，利用递归解决

> 递归代码模板

- 终止条件
- 找出重复问题并归纳，写成子问题处理模板
- 改变输入条件，下探一层
- 递归返回后清理

```
func recursion(level int, maxLevel int, p1, p2, int) {
    // 终止边界条件
    if level > maxLevel {
        process_result
        return
    }

    // 归纳出的重复的核心逻辑处理
    process(level, data)

    // 下探一层处理
    recursion(level+1, maxLevel, p1, p2)

    // 返回后进一步处理
    cleanCurStatus()
}
```

> 5、分治解题思路

```
func divideConquer(problem *P, param1, param2) {
     // 子问题结束条件
    if p == nil {
        结果处理
        return 
    } 
  
    // 处理当前问题， 
    subproblems = splitProblem(problem)

    // 下探一层处理子问题
    res1 := divideConquer(subproblems[0], p1)
    res2 := divideConquer(subproblems[2], p1)

    // 合并子问题
    result := mergeSubResult(res1, res1)

    reverseStates()
}
```

> 回溯

回溯一个认识：对所有可能的问题去试探，然后判断是否满足条件，然后退回来，继续试探,

然后去掉不满足的得到最终结果, 因此从做题中归纳出基本一个思路

- 分出小的问题，然后画出树结构,然后寻找规律

- 对树中不满足的情况总结，用作边界条件


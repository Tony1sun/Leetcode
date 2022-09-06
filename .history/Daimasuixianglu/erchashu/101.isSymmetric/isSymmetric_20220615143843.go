package main

import "github.com/halfrost/LeetCode-Go/structures"

type TreeNode = structures.TreeNode

func isSymmetric(root *TreeNode) bool {
    var queue []*TreeNode;
    if root != nil {
        queue = append(queue, root.Left, root.Right);
    }
    for len(queue) > 0 {
        left := queue[0];
        right := queue[1];
        queue = queue[2:];
        if left == nil && right == nil {
            continue;
        }
        if left == nil || right == nil || left.Val != right.Val {
            return false;
        };
        queue = append(queue, left.Left, right.Right, right.Left, left.Right);
    }
    return true;
}

作者：carlsun-2
链接：https://leetcode.cn/problems/symmetric-tree/solution/101-dui-cheng-er-cha-shu-di-gui-fa-die-d-vvf4/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
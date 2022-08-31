package Leetcode

import "github.com/halfrost/LeetCode-Go/structures"

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    result:=&ListNode{}
    result.Next=head
    var pre *ListNode
    cur:=result

    i:=1
    for head!=nil{
        if i>=n{
            pre=cur
            cur=cur.Next
        }
        head=head.Next
        i++
    }
    pre.Next=pre.Next.Next
    return result.Next

}

作者：carlsun-2
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/dai-ma-sui-xiang-lu-19-shan-chu-lian-bia-2hxt/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

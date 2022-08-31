package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    //声明快慢指针
    slow,fast := head,head
    //链表为nil，直接返回
    if fast == nil {
        return head
    }
    //fast指针先跑n步
    for n>0 {
        fast=fast.Next
        n--
    }
    //如果这时fast为nil，则head节点是要删除的那个(容易忽略)
    if fast == nil {
        return head.Next
    }
    //快慢指针一起移动
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }
    //slow.next为需要删除的节点
    slow.Next = slow.Next.Next
    return head
}

作者：musing-fermin3z
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/19-shan-chu-lian-biao-de-dao-shu-di-n-ge-376g/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {

}

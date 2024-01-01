# List

## 链表结构体
```
// go
type ListNode struct {
    Val     int
    Next    *ListNode
} 
```

## 反转链表(reverse)
```
func reverseList(head *ListNode) *ListNode {
	var pre, cur, nxt *ListNode = nil, head, nil
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
```

## 反转链表II
```
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}

	var pre, cur, nxt *ListNode = nil, p0.Next, nil
	for i := 0; i < right-left+1; i++ {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	
	p0.Next.Next = cur
	p0.Next = pre
	
	return dummy.Next
}
```
## K个一组反转链表
```
func reverseKGroup(head *ListNode, k int) *ListNode {
	cnt := 0
	cntP := head
	for cntP != nil {
		cnt++
		cntP = cntP.Next
	}
	dummy := &ListNode{
		Next: head,
	}
	p0 := dummy
	var pre, cur, nxt *ListNode = nil, p0.Next, nil

	for i := 0; i+k-1 < cnt; i += k {
		j := 0
		for j < k {
			nxt = cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
			j++
		}

		nxt = p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}
```

## 相交链表
链表A 和 链表B要么不相交，要么从某一点开始共享之后的节点，返回空或者相交节点
```
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
```
## 排序链表
```
func mergeLists(a, b *ListNode) *ListNode {
	var va, vb int
	dummy := &ListNode{
		Next: nil,
	}
	cur := dummy
	for a != nil || b != nil {
		if a == nil {
			va = math.MaxInt
		} else {
			va = a.Val
		}
		if b == nil {
			vb = math.MaxInt
		} else {
			vb = b.Val
		}
		if va <= vb {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, slow, fast *ListNode = nil, head, head
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	pre.Next = nil
	curA := sortList(head)
	curB := sortList(slow)
	cur := mergeLists(curA, curB)
	return cur
}
```
## 重排链表
将表从中间拆分，然后后半部分倒序后插入到前半部分的间隙（两个节点之间）里
```
func findHeadOf2thHalf(head *ListNode) *ListNode {
	var pre, slow, fast *ListNode = nil, head, head
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	pre.Next = nil

	return slow
}

func reverse2thHalf(head *ListNode) *ListNode {
	var pre, cur, nxt *ListNode = nil, head, nil
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func merge(a, b *ListNode) {
	curA, curB := a, b
	var nxtA, nxtB *ListNode
	for curB != nil {
		nxtA = curA.Next
		nxtB = curB.Next
		curB.Next = curA.Next
		curA.Next = curB
		curA = nxtA
		curB = nxtB
	}
}

func reorderList(head *ListNode) {
	mid := findHeadOf2thHalf(head)
	mid = reverse2thHalf(mid)
	merge(head, mid)
}
```
## 判断回文链表
快慢指针找出中间节点，翻转后半部分指针，同时遍历判断
```
func isPalindrome(head *ListNode) bool {
    
    slow, fast := head, head
    var pres *ListNode
    for fast != nil {
        pres = slow
        slow = slow.Next
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
    }
    pres.Next = nil
    
    cur := slow
    var nxt, pre *ListNode
    for cur != nil {
        nxt = cur.Next
        cur.Next = pre
        pre = cur
        cur = nxt
    }
    
    for head != nil && pre != nil {
        if head.Val != pre.Val {
            return false
        }
        head = head.Next
        pre = pre.Next
    }

    return true
}
```
## 环形链表，找环形入口节点，没有返回空
快慢指针
设非环段长为a，环入口节点到快慢指针相遇的位置长度为b，环的剩余长度为c
```
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
```
## 链表删除某一节点
```
// 感觉还是附带一个头结点比较
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre, cur := dummy, head
	var nxt *ListNode
	for cur != nil {
		nxt = cur.Next
		if cur.Val == val {
			cur.Next = nil
			pre.Next = nxt
			break
		}
		pre = cur
		cur = nxt
	}
	return dummy.Next
}
```
## 复杂链表的赋值
```
type Node struct {
Val    int
Next   *Node
Random *Node
}

func copyRandomList(head *Node) *Node {
hash := make(map[*Node]*Node, 0)
cur := head
for cur != nil {
hash[cur] = &Node{
Val:    cur.Val,
Next:   nil,
Random: nil,
}
cur = cur.Next
}

	cur = head
	for cur != nil {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur = cur.Next
	}
	return hash[head]
}
```
## 合并k个升序链表
```
// 堆
type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *hp) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &hp{}
	heap.Init(h)
	for i := range lists {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	dummy := &ListNode{
		Next: nil,
	}
	var pre, cur *ListNode = dummy, nil
	for h.Len() > 0 {
		cur = heap.Pop(h).(*ListNode)
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return dummy.Next
}
```

```

```

```

```
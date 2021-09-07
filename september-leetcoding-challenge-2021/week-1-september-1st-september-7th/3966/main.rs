#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

impl Solution {
    pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut root = None;
        let mut p = head;

        while let Some(mut node) = p.take() {
            let tmp = node.next.take();
            node.next = root.take();
            root = Some(node);

            p = tmp;
        }
        return root;
    }
}

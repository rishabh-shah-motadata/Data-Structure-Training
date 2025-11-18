### Data Structure Time & Space Complexity

| Data Structure | Access (by Index/Key) | Search (by Value/Key) | Insertion | Deletion | Space | Ordered Traversal / Sorting |
| --- | --- | --- | --- | --- | --- | --- |
| Array | O(1) | O(n) | O(n)¹ | O(n)¹ | O(n) | O(n log n) |
| Dynamic Array | O(1) | O(n) | O(n)¹ (End: O(1) amort.)² | O(n)¹ (End: O(1)) | O(n) | O(n log n) |
| LinkedList (Singly) | O(n) | O(n) | O(1) (Head) / O(n) (Mid) | O(1) (Head) / O(n) (Mid) | O(n) | O(n log n) |
| Doubly LinkedList | O(n) | O(n) | O(1) (Head/Tail) | O(1) (Head/Tail) | O(n) | O(n log n) |
| Circular LinkedList | O(n) | O(n) | O(1) (Head) / O(n) (Mid) | O(1) (Head) / O(n) (Mid) | O(n) | O(n log n) |
| Stack | N/A (Top: O(1)) | O(n) | O(1) (Push) | O(1) (Pop) | O(n) | N/A |
| Queue | N/A (Front: O(1)) | O(n) | O(1) (Enqueue) | O(1) (Dequeue) | O(n) | N/A |
| Circular Queue | N/A (Front: O(1)) | O(n) | O(1) (Enqueue) | O(1) (Dequeue) | O(n) | N/A |
| Doubly Ended Queue | N/A (Ends: O(1)) | O(n) | O(1) (At ends) | O(1) (At ends) | O(n) | N/A |
| Priority Queue | N/A (Top: O(1)) | O(n) | O(log n) | O(log n) (Extract-Top) | O(n) | N/A (Used for sorting) |
| Hashmap | O(1) (Avg) / O(n) (Worst) | O(1) (Avg) / O(n) (Worst) | O(1) (Avg) / O(n) (Worst) | O(1) (Avg) / O(n) (Worst) | O(n) | N/A (Unordered) |
| BST (Binary Search Tree) | O(log n) (Avg) / O(n) (Worst) | O(log n) (Avg) / O(n) (Worst) | O(log n) (Avg) / O(n) (Worst) | O(log n) (Avg) / O(n) (Worst) | O(n) | O(n) |
| AVL Tree (Balanced BST) | O(log n) | O(log n) | O(log n) | O(log n) | O(n) | O(n) |
| B-Tree (Balanced Tree) | O(log n) | O(log n) | O(log n) | O(log n) | O(n) | O(n) |

---

### Data Structure Cache & Memory Profile

| Structure | Cache Locality | Memory Overhead |
| --- | --- | --- |
| Array / Dynamic Array | ⭐ Excellent | Very Low |
| LinkedList | 😞 Very Poor | Moderate (1 pointer/item) |
| Doubly LinkedList | 😞 Very Poor | High (2 pointers/item) |
| BST / AVL Tree | 😐 Poor | High (2+ pointers/item) |
| Hashmap | 😐 (Depends) | High (Load factor + entries) |
| B-Tree | ⭐ Excellent (for disk) | High (but efficient) |
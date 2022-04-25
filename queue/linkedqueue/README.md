Package linkedqueue implements a queue using a singly linked list.

Supported operations:
* Enqueue -> O(1)
* Dequeue -> O(1)
* First	 -> O(1)
* Size	 -> O(1)
* IsEmpty -> O(1)
* String  -> O(n)

Factory functions:
* New	 -> O(1)

```
import "github.com/MehdiEidi/gods/queue/linkedqueue"
```

```
q := linkedqueue.New[int]()

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

fmt.Println(q.Dequeue())

fmt.Println(q)
```
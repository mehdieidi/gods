## Circular Queue

Package circularqueue implements a circular queue using a circular linked list.

The key point of a circular queue is the Rotate operation which can move the front element to the
end of the queue.

Supported operations:
* Enqueue -> O(1)
* Dequeue -> O(1)
* First	 -> O(1)
* Size	 -> O(1)
* IsEmpty -> O(1)
* Rotate	 -> O(1)
* String  -> O(n)

Factory functions:
* New	 -> O(1)

```
import "github.com/MehdiEidi/gods/queue/circularqueue"
```

```
q := circularqueue.New[int]()

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

q.Rotate()

fmt.Println(q)
```
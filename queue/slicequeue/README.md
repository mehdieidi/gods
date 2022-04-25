Package slicequeue implements a queue using a slice.

Despite the worse running time of Enqueue operation of slice-queue than linked-queue, the slice-queue
has the help of the CPU cache and consumes less auxiliary memory.

Supported operations:
* Enqueue -> O(1) if there is capacity, O(n) if capacity runs out and new array needs to be allocated.
* Dequeue -> O(1)
* First	 -> O(1)
* Size	 -> O(1)
* IsEmpty -> O(1)
* String

Factory functions:
* New	 -> O(1)

```
import "github.com/MehdiEidi/gods/queue/slicequeue"
```

```
q := slicequeue.New[int]()

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

fmt.Println(q.Dequeue())

fmt.Println(q)
```
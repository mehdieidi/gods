# gods
Library of generic data structures for Go.

* [priority queue](./priorityqueue/)
  * [sorted list priority queue](./priorityqueue/sortedpq/)
  * [unsorted list priority queue](./priorityqueue/unsortedpq/)
  * [heap priority queue](./priorityqueue/heappq/)
  * [adaptable heap priority queue](./priorityqueue/adaptablepq/)
* [tree](./tree/)
  * [binary tree](./tree/binarytree/)
  * [general tree](./tree/generaltree/)
* [queue](./queue/)
  * [queue using linked list](./queue/linkedqueue/)
  * [queue using slice](./queue/slicequeue/)
  * [circular queue](./queue/circularqueue/)
* [stack](./stack/)
  * [stack using linked list](./stack/linkedstack/)
  * [stack using slice](./stack/slicestack/)
* [bitset](./bitset/)
* [linked list](./linkedlist/)
  * [singly linked list](./linkedlist/singly/)
  * [doubly linked list](./linkedlist/doubly/)
  * [circularly linked list](./linkedlist/circularly/)
* [set](./set/)
* [deque](./deque/)
* [positional list](./positionallist/)


  
  --- more coming soon :D ---

## Install
```
$ go get github.com/MehdiEidi/gods
```

## Import
```
import "github.com/MehdiEidi/gods"
```

## Usage
Either use convenience helper functions in gods package like this:<br>
Example: Creating a new set:
```
s := gods.NewSet[string]()

s.Add("awesome :D")
```
or directly import the intended data structure and use its factory function:
```
import "github.com/MehdiEidi/gods/set"
.
.
.

s := set.New[string]()
s.Add("again awesome :D")
```


# gods
Library of generic data structures for Go.


* [bitset](./bitset/)
* [linked list](./linkedlist/)
  * [singly linked list](./linkedlist/doubly/)
  * [doubly linked list](./linkedlist/doubly/)
  * [circularly linked list](./linkedlist/circularly/)
* [set](./set/)
* [stack](./stack/)
  * [stack using linked list](./stack/)
  * [stack using slice](./stack/)
* [queue](./queue/)
  * [queue using linked list](./queue/)
  * [queue using slice](./queue/)
  * [circular queue](./queue/)
* [deque](./deque/)
  
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


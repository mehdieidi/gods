// Package slicequeue implements a queue using a slice.
//
// Supported operations:
// * Enqueue -> O(1) if there is capacity, O(n) if capacity runs out and new array needs to be allocated.
// * Dequeue -> O(1)
// * First	 -> O(1)
// * Size	 -> O(1)
// * IsEmpty -> O(1)
// * String
//
// Factory functions:
// * New	 -> O(1)
package slicequeue

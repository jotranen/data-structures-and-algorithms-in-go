package ch06

// Given K input stream of number in sorted order. You need to make a single output stream, which contains all the elements of the K streams in sorted order. The input stream support ReadNumber() operation and output stream support WriteNumber() operation.
// Hint:
// a) Read the first number form all the K input streams and add them to a Priority Queue. ()Nodes should keep track of the input stream).
// b) Dequeue on element at the time form PQ, put this element value to the output stream, Read the input stream number and from the same input stream add another element to PQ.
// c) if the stream is empty, just continue
// d) Repeat until PQ is empty


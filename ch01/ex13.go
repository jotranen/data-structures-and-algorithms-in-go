package ch01

// Find the duplicate elements in a list of size n where each element is inthe range 0 to n-1
//Hint: Approach 1: Compare each element with all the elements of the list (using two loops O(n2) solution)
//Approach 2: Maint a Hash-Table. Set the has value to 1 if we encounter the element for the first time. When we same value again we can see that the hash value is already 1 so we can print tht value
//O(n) solutio, but additional space if required
//Approach 3: We will explit the constraing "every element is tin the tnage 0 to n-1" We can tkae a list arr[] or size n and set all the elements to 0. Whenever we get a value say val1.
//	We will inclement the value at arr[val1] index by 1. Int then, we can traverse the list arr and print repeated values.
//	Additions space complexity will be O(n) whicl will be less than Hash-Table approach.

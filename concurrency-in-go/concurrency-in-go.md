# Concurrency in Go

## Notes from [this book](https://www.amazon.in/Concurrency-Go-Tools-Techniques-Developers-ebook/dp/B0742NH2SG).

### Race Condition
A race condition occurs when two or more operations must execute in the correct order, but the program has not been 
written so that this order is guaranteed to be maintained.

### Data Race
Data race happens where one concurrent operation attempts to read a variable while at some undetermined time another 
concurrent operation is attempting to write to the same variable.

### Sleeps
Introducing sleeps into your code can be a handy way to debug concurrent programs, but they are not a solution.

### Critical Section
In fact, there’s a name for a section of your program that needs exclusive access to a shared resource. 
This is called a critical section.

### `Lock` Makes our Program Slow
The calls to Lock you see can make our program slow.

### Deadlocks, Livelocks and Starvation
These issues all concern ensuring your program has something useful to do at all times. If not handled properly, your 
program could enter a state in which it will stop functioning altogether.

### Deadlocks
A deadlocked program is one in which all concurrent processes are waiting on one another. In this state, the program 
will never recover without outside intervention.

### Coffman Conditions
Coffman Conditions and are the basis for techniques that help detect, prevent, and correct deadlocks.

The Coffman Conditions are as follows:
#### Mutual Exclusion
A concurrent process holds exclusive rights to a resource at any one time.
#### Wait For Condition
A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
#### No Preemption
A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
#### Circular Wait
A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting 
on it (P1), so it fulfills this final condition too.

### Livelock
Livelocks are programs that are actively performing concurrent operations, but these
operations do nothing to move the state of the program forward.

Have you ever been in a hallway walking toward another person? She moves to one
side to let you pass, but you’ve just done the same. So you move to the other side, but
she’s also done the same. Imagine this going on forever, and you understand livelocks.

A very common reason livelocks are written: two or more concurrent processes attempting to prevent a deadlock without coordination.
In my opinion, livelocks are more difficult to spot than deadlocks simply because it can appear as if the program is doing work.
Livelocks are a subset of a larger set of problems called starvation.

### Starvation
Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.
Livelocks warrant discussion separate from starvation because in a livelock, all the concurrent processes 
are starved equally, and no work is accomplished
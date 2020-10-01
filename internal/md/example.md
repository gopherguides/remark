class: center, middle, inverse
duration: 1h30m
exercises: 1 Shared Lab, 2 student Labs

# Concurrency With The Sync Package

---

overview: true

## Overview

Go promotes sharing memory by communicating. This chapter covers goroutines and how to synchronize communication between them. Mechanics for synchronization such as `WaitGroups` and `Mutexes` are explored along with the corresponding patterns for each.

---

## Synchronization Between Goroutines

Since goroutines run in the same address space, access to shared memory must be synchronized.

The `sync` package provides useful primitives to accomplish this, but there are also other primitives such as channels as well.

---

## The Problem

Run the following program:

<code src="./src/examples/goroutine/main.go"></code>

---

## Concurrency & Goroutines

Goroutines execute concurrently. In this case, `main()` and `sayHello()` were running at the same time. `main()` finished and `sayHello()` never got a chance to run!

<code src="./src/examples/goroutine/main.go"></code>

To solve this problem you need to use some sort of communication, or synchronization device, to signal that a goroutine is complete.

We typically do that with channels, or a `WaitGroup`, depending on whether you need to return a value from the function you're running.

.footnote[Using the `go` keyword in front of a function only `schedules` the function to be run, it does not run it immediately.]


---

## WaitGroup

A `WaitGroup` is defined in the `sync` package in the standard library. They're a simple, and powerful, way to ensure that all of the goroutines you launch have completed before you move on to do other things.

.footnote[[https://golang.org/pkg/sync/#WaitGroup](https://golang.org/pkg/sync/#WaitGroup) ]

---

### Understanding A WaitGroup

A `WaitGroup` is, essentially, an atomic counter, that can be safely accessed concurrently. A `WaitGroup` had three methods:

* `Add(i int)`
* `Done()`
* `Wait()`

When a new `WaitGroup` is created its internal counter is set to `0`.

---

#### WaitGroup.Add

To add more items to a `WaitGroup` we can use the `Add(n int)` method. You can call `Add` as many times as you need to.

You can add one item at a time:

```go
wg.Add(1)
```

Or add more than one:

```go
wg.Add(42)
```

The number you add will almost always be the total number of goroutines you intend to launch.

If you create a new `WaitGroup`, and add `5` to it:

```go
wg.Add(5)
```

And then later on add `42`:
```go
wg.Add(42)
```

The internal counter of the waitgroup  will now be a total of `47`.


.footnote[It is important to remember that while you can call `Add` as many times as needed, once `Wait` is called, `Add` should no longer be called. Doing so may create a race condition in your code and will result in potentially non-deterministic code.]

---

#### WaitGroup.Done

To remove an item from the group, we can call the `Done()` method. This method will decrement the counter by `1`. So, if the `WaitGroup` contains `42` items, calling `wg.Done()`, will set the counter to `41`.

You will almost always call `Done()` from within the Go Routine you launched.  This signals to your program that the Go Routine has finished.

---

#### WaitGroup.Wait

The `Wait()` function on a `WaitGroup` will block execution until the internal counter is `0`. This is always done after all `Add`'s have been called.

The common flow of using a `WaitGroup` will look like this:

1. Declare your `WaitGroup`
1. Call `Add` as many times as necessary and launch corresponding goroutines
1. Call `Wait` from your current routine
1. Call `Done` from inside each of the launched goroutines from your second step

Using the methods of a `WaitGroup` in this order will help ensure proper code syncronization.

---

### WaitGroup Example

<code src="./src/examples/waitgroup/main.go"></code>

---

### Using Closures With A WaitGroup

A `WaitGroup` is also especially powerful when combined with closures (anonymous functions), however, care must be used to prevent subtle bugs.

Below is a bug that is common to new developers in Go:

<code src="./src/examples/closure/bug/main.go" section="main"></code>

---

#### Closure Bug Explanation

The first example has the scope of `i` and `n` at the outer scope. Since the goroutines are being added to the execution stack, but not run, by the time they run, and access that memory space, `i` and `n` have already been changed.

Sending `i` and `n` in as an argument creates a new scope and copies the value to a new memory location to preserve the proper value.

<code src="./src/examples/closure/solution/main.go"></code>

.footnote[When you see a loop that launches goroutines, you should pay careful attention to if any variables need to be captured in a new scope to avoid possible bugs.]

---

#### Go Vet Closure Detection

If you run `go vet` on your code, it can now detect any closure variables that are captured:

```sh
main.go:19:47: loop variable n captured by func literal
main.go:19:50: loop variable i captured by func literal
```

---

## Using Error Groups

The [`golang.org/x/sync/errgroup`](https://godoc.org/golang.org/x/sync/errgroup) package allows for a simpler way to manage goroutines, while also allowing for error handling.

<code src="./src/examples/errgroup/main.go" section="main"></code>

---

### Error Groups Behavior

While using an error group may simplify your program's ability to retrieve the error, it will only report the first error.  You won't be able to determine if you had other errors, or what they were.  This should be taken into account when using the `errgroup.Group`.

<code src="./src/examples/errgroup/multiple/main.go" section="main"></code>

Output:
<code src="./src/examples/errgroup/multiple/main.go" section="output"></code>

---

### Context With ErrorGroup

The [`golang.org/x/sync/errgroup#WithContext`](https://godoc.org/golang.org/x/sync/errgroup#WithContext) function will create a new `errgroup.Group` as well as return a `context.Context` that others can listen to for cancellation should the `errgroup.Group` return an `error`.

<code src="./src/examples/errgroup-cancel/main.go" section="main"></code>

---

### ErrorGroup Vs. WaitGroup

The question of when to use which is a matter of personal preference, as well as, the task at hand.

An `ErrorGroup` takes away a lot of the granular control that you with a `WaitGroup`, but they give you a simpler API and easier error handling.

A `WaitGroup` provides total control, but you have to plumb in your error handling.

---

## Code Walkthrough (Optional)

Now that we have a basic understanding of how a `WaitGroup` works, we can take a process that is slow and sequential, and make it fast and efficient.

<hr></hr>

What if We are asked to download the top 100 sites each day and monitor how long their response times are.  We want this to process as fast as possible.

Given the following code, let's work through the problem and use our new knowledge of concurrency:

<code src="./src/examples/walkthrough/sites/v1/main.go"></code>

---

### Refactor

To simplify the code, let's go ahead and refactor the actual retrieving of the website into its own function:

<code src="./src/examples/walkthrough/sites/v2/main.go" section="code"></code>

Here is the output:

<code src="./src/examples/walkthrough/sites/v2/main.go" section="output"></code>

---

### Record Timing

Now that we have a basic structure, let's record the time it takes for each site to be retrieved, as well as how much time it takes to run the entire process:

<code src="./src/examples/walkthrough/sites/v3/main.go" section="code"></code>

Here is the output:

<code src="./src/examples/walkthrough/sites/v3/main.go" section="output"></code>

Because we are running everything sequentially, it takes longer than we would like to process these sites.  Currently we are only doing 3 sites, so this problem will only get worse as we add more sites to process.

---

### Add Concurrency

We are at a point now that we can add our concurrency primitives and make this code more efficient:

<code src="./src/examples/walkthrough/sites/v4/main.go" section="code"></code>

Here is the output:

<code src="./src/examples/walkthrough/sites/v4/main.go" section="output"></code>

Now the process only takes about as long as the longest site takes to retrieve.  This is a great improvement.

---

### Processing 100 sites

Next, let's process all the sites. We can use a list of the top 100 sites. Here is a sample of the list we will process:

<code src="./src/examples/walkthrough/sites/v5/sites.txt" end="5"></code>

First, we'll have to get the arguments from the input now. We can use that by just asking for `os.Args[1:]`.  We slice off the first argument, as that is always the name of the program that is running.

To process this file, I can use the following command:

```command
go run main.go "$(< sites.txt)"
```

<code src="./src/examples/walkthrough/sites/v5/main.go" section="code"></code>

Here is a sample of the output of all 100 sites:


<code src="./src/examples/walkthrough/sites/v5/sites.output" end="5"></code>
<code src="./src/examples/walkthrough/sites/v5/sites.output" start="95"></code>

---

### Outstanding Issues

While we were able to process all the files faster, we still have some issues we need to address to use this in a real production scenario:

* Processing 100 sites may work, but what happens if we have 1,000, or 1,000,000?  This solution won't scale properly in that scenario.
* Currently some sites don't load and time out. The default timeout is 10 minutes, which we don't want to wait for. We could control this with the http client, but using other concurrency primitives we can add our own timeout.
* All the errors are happening inside the `get` function.  Ideally we would want to pass these errors back to be properly handled upstream in our application.

---

## WaitGroup Exercise (10 min)

The following program currently works but could be improved by adding concurrency.

Make the following changes in the program:
* Launch `charCount` in a go routine.
* Use a `sync.WaitGroup` to track each time you launch a new go routine.
* Make sure to wait for each go routine to finish before exiting the program.
* Make sure to mark each go routine as `Done`.

Note: You can download this file and the files it uses with the `Download Source` button on the left.

Be sure to run the example from the same directory that `main.go` is in to ensure it can find the downloaded files.

<code src="./src/exercises/01-waitgroup/main.go"></code>

---

### WaitGroup Solution

<code src="./src/solutions/01-waitgroup/main.go"></code>

---

class: center, middle, inverse

# Protecting Data With A Mutex

---

## Protecting Data With A Mutex

A `Mutex` allows you to synchronize goroutines for safe access to the same shared memory. Go provides two `Mutex` implementations, a [full lock](https://golang.org/pkg/sync/#Mutex), and a [read/write](https://golang.org/pkg/sync/#RWMutex) lock.

```go
var m sync.Mutex // Make a new mutex

m.Lock() // lock the mutex

// Code in here that is protected by this mutex is safe to
// read or write for all goroutines observing this mutex.

m.Unlock() // unlock the mutex
```

.footnote[[https://golang.org/pkg/sync/#Mutex](https://golang.org/pkg/sync/#Mutex) ]

.footnote[[https://golang.org/pkg/sync/#RWMutex](https://golang.org/pkg/sync/#RWMutex) ]

---

## Running Without A Mutex

When trying to access a `map` concurrently, without a mutex, the app will `panic` and crash.

Here are two functions that access the same map:
<code src="./src/examples/mutex/broken/main.go" section="funcs"></code>

Call them concurrently without a mutex to synchronize access:
<code src="./src/examples/mutex/broken/main.go" section="main"></code>

If Go detects access to the same memory in different goroutines (specifically in a map), it will panic:
```bash
fatal error: 9
concurrent map writes

goroutine 54 [running]:
runtime.throw(0x10c4c4b, 0x15)
/usr/local/go/src/examples/runtime/panic.go:608 +0x72 fp=0xc0000b4760 sp=0xc0000b4730 pc=0x1027092
runtime.mapassign_fast64(0x10a7f00, 0xc000078120, 0x12, 0x0)
```

---

### Synchronizing With A Mutex

By using a `Mutex` in the `get` and `set` functions that wrap the `map`, we are able to synchronize access to the map and prevent a `panic` in the application.

<code src="./src/examples/mutex/simple/main.go" section="funcs"></code>

---

### Synchronizing With A RWMutex

We can make use of a `RWMutex` as well so that we can have multiple read locks open allowing for less lock contention in a read heavy application:

<code src="./src/examples/mutex/simple-rw/main.go" section="funcs"></code>

---

### Embedding A Mutex

You can embed a mutex in your own struct to synchronize access to internal data:
<code src="./src/examples/embed/embed/main.go" section="embed"></code>

However, it's more common to use a mutex as a field to prevent method promotion and ensure that only your struct can call the methods on the mutex.
<code src="./src/examples/embed/field/main.go" section="embed"></code>

---

### Mutex Lock Order

Mutexes can run in two modes: `normal` and `starvation`.

* In normal mode, waiters are queued in `FIFO` order.
* In starvation mode, ownership of the mutex is directly handed off from the unlocking goroutine to the waiter at the front of the queue.

From the [source](https://golang.org/src/sync/mutex.go?#L42):

```go
// Mutex fairness.
//
// Mutex can be in 2 modes of operations: normal and starvation.
// In normal mode waiters are queued in FIFO order, but a woken up waiter
// does not own the mutex and competes with new arriving goroutines over
// the ownership. New arriving goroutines have an advantage -- they are
// already running on CPU and there can be lots of them, so a woken up
// waiter has good chances of losing. In such case it is queued at front
// of the wait queue. If a waiter fails to acquire the mutex for more than 1ms,
// it switches mutex to the starvation mode.
//
// In starvation mode ownership of the mutex is directly handed off from
// the unlocking goroutine to the waiter at the front of the queue.
// New arriving goroutines don't try to acquire the mutex even if it appears
// to be unlocked, and don't try to spin. Instead they queue themselves at
// the tail of the wait queue.
//
// If a waiter receives ownership of the mutex and sees that either
// (1) it is the last waiter in the queue, or (2) it waited for less than 1 ms,
// it switches mutex back to normal operation mode.
//
// Normal mode has considerably better performance as a goroutine can acquire
// a mutex several times in a row even if there are blocked waiters.
// Starvation mode is important to prevent pathological cases of tail latency.
```

---

## Use Mutexes To Build A WaitGroup

Using a simple implementation of a `WaitGroup` we can demonstrate how a `Mutex` can help to prevent race conditions in our code.

<code src="./src/examples/mutex/waitgroup/broken/main.go" section="wg"></code>


---

### Enabling Race Detection

Data races occur when multiple goroutines access the same shared memory.  If you are writing concurrent code, you should always enable the race detector.

You can do this both via testing, or with the `go run` command.

To turn on the race detector, pass in the `-race` flag:

```sh
go run -race ./main.go
```

If a race is detected, it will both fail the test and output where the race was detected.  The race detector makes one promise: it will never report a false positive.  However, due to how the race detector is implemented, it will ONLY detect races that occur in the current running code.

---

### Detecting Races

The current implementation of the waitgroup does not have a mutex to synchronize access to shared memory.  This creates the potential for a data race.

We can see this by turning on the race detection with the `-race` flag:

```sh
go run -race ./broken.go
```


```bash
==================
WARNING: DATA RACE
Read at 0x00c000012058 by goroutine 9:
  main.(*WaitGroup).Done()
      main.go:16 +0x47
  main.main.func1()
      main.go:47 +0xcf

Previous write at 0x00c000012058 by main goroutine:
  main.main()
      main.go:25 +0xb3

Goroutine 9 (running) created at:
  main.main()
      main.go:43 +0xe1
```

---

#### Adding A Mutex

Adding a mutex to the WaitGroup will allow us to synchronize access to memory.

<code src="./src/examples/mutex/waitgroup/main.go" section="wg"></code>

---

#### Using A Read/write Mutex

Using `sync.RWMutex` we can lock access to data based on how that data is being used.

An `sync.RWMutex` is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer.

<code src="./src/examples/mutex/rw/main.go" section="funcs"></code>

---

## Mutex Bug

The following code has a subtle bug that is common when using mutexes. Can you spot the bug?

<code src="./src/examples/mutex/bug/bug/main.go" section="main"></code>

---

### Mutex Bug Solution

The `defer` keyword only executes once it exits the function it was called in. Because it is in a loop, the `defer` calls get put on the execution stack, but never gets called because the loop is running within the function scope.

In these scenarios, you will need to handle each unlock without the use of the `defer` keyword.

<code src="./src/examples/mutex/bug/solution/main.go" section="main"></code>

---

### Mutex Bug Common Approach

It is much more common to launch many goroutines inside a for loop, than to launch a for loop in a goroutine.

The following is more idiomatic for handling the mutex with defers:

<code src="./src/examples/mutex/bug/common/main.go" section="main"></code>

---

## Debugging Mutexes

On occasion, you may end up in a deadlock on a `Mutex`. They tend to be very difficult to track down.

The following packages can assist in debugging the offending lock. They are utilities you don't need until
you've done something **totally wrong** and desperately need help!

* [syncdebug](https://github.com/go4org/go4/tree/master/syncutil/syncdebug) - Package by Brad Fitzpatrick.
* [Debug Mutex Package](https://godoc.org/github.com/apptimistco/asn/debug/mutex)

---

### Debugging Our Bug

We can see that if we want to debug our offending code from the previous example, we can use the `syncdebug` package.

First, get the package:

<code src="./src/examples/mutex/bug/debug/main.go" section="get"></code>

Here is the code using the debug package:

<code src="./src/examples/mutex/bug/debug/main.go" section="main"></code>

And here is the output when we run the program:

<code src="./src/examples/mutex/bug/debug/main.go" section="output"></code>

Notice that it no longer deadlocks.  This is because it starts up it's own go routine to monitor how many locks are currently acquired and waiting.

---

## Mutex Panics

Care must be taken to make sure not to try and unlock a `Mutex` that is not currently locked.

<code src="./src/examples/unlocked/main.go"></code>

```bash
fatal error: sync: unlock of unlocked mutex
```

---

## Mutex Exercise (10 Mins)

Add a `sync.RWMutex` to the code below to prevent a race condition.

If you run the code with the race detector on, you will see it finds a race condition:

```
go run -race ./main.go
```

<code src="./src/exercises/02-mutex/main.go"></code>

---

### Mutex Solution

<code src="./src/solutions/02-mutex/main.go"></code>

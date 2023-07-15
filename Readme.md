Nana's project idea version wise:

version 10:
# Project Name: Go Conference (Booking system)
-----------------------------------------------

---

Task 1. Apply concurrency using synchronous goroutines

### Here if we use normal concurrency mean without synchronizing then for infinite loop we will see our expectd output flow but if we do not use infinite loop and waitgroup then in this case we will not able to see our desire output flow because "By default the main goroutine does not wait for other goroutines" So the solution of this problem is synchonizing go routines using waitgroups.

---

To do so at first we have create waitgroup then we have to place three function of it in the right place.

1. Creating wait group:
```var wg = Sync.Waitgroup{}```

2. Three functions of waitgroup are 
   i. wg.Add(place the number of thread) i.e: wg.Add(1) means 1 thread. We need to put this line before the keyword go func() call.
   
   ii. wg.Wait() we need to put this function at the end of the main thread means main func

   iii. wg.Done() we need to put this function as end line of our delay function. In our case it is func sendTickets()

---

# More specifically:
-------------------
To implement this, we need to perform the following steps:
```
=> Create a waitgroup:
var wg = sync.WaitGroup
Use three functions of the waitgroup:
wg.Add() - Place the number of threads to be added.
Example: wg.Add(1) means adding 1 thread.

We need to put this line before calling the goroutine using the go keyword.
wg.Wait() - Place this function at the end of the main function to make the main goroutine wait for other goroutines to complete.
wg.Done() - Place this function as the last line of the delay function (in this case, func sendTickets()).
--------------------------------------------------------------------

Here is an example implementation:
-----------------------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

func sendTickets() {
	// Delay function implementation
	time.Sleep(1 * time.Second)
	fmt.Println("Tickets sent!")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go sendTickets()

	wg.Wait()
	fmt.Println("All goroutines completed!")
}
```
By using the waitgroup, we ensure that the main goroutine waits for the sendTickets goroutine to complete before printing "All goroutines completed!" at the end.

## Happy coding!

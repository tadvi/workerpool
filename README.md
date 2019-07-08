# workerpool

WorkerPool for basic goroutine management

	wp := workerpool.New(100)

	for i := 1; i < 1000; i++ {
		log.Println("started on", n)

		wp.Go(func() {
			// Do some work.
		})
	}

	// Make sure all workers done.
	wp.Wait()

Simple use with closures:

    wp := workerpool.Run(4, func() {
        fmt.Println("works")
    })

    // Make sure all workers done.
    wp.Wait()
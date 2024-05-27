Package used for concurrent jobs

Usage:

```
wp := workerpool.NewWorkerPool(10, 10)
wp.Run()

for i := range tasks {
	wp.AddTask(func() {
		tasks.Execute()
	})
}

```

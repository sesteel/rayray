rayray
=========
#### A Goroutine Experiment

This project demonstrates the use of WaitGroups and goroutines to divide and conquer the rendering of a single ray traced frame.  By changing the the number of sections in use by the algorithm, we see various performance characteristics as listed here:

| # Sections | CPU Load | #Threads | Time (ms)   |
| ---------- | -------- | -------- | ----------- |
| 1          | 100%     | 9        | ~ 524       |
| 2          | 198%     | 9        | ~ 275       |
| 5          | 337%     | 9        | ~ 145       |
| 10         | 571%     | 9        | ~ 145       |
| 20         | 617%     | 9        | ~ 129       |
| 50         | 733%     | 9        | ~ 127       |

In the above, each frame is rendered 200 times and the average
frame render time is recorded.  

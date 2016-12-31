rayray
=========
#### A Goroutine Experiment

This project demonstrates the use of WaitGroups and goroutines to divide and conquer the rendering of a single ray traced frame.  By changing the the number of sections in use by the algorithm, we see various performance characteristics as listed here:

| # Sections | CPU Load | #Threads | Time (ms)   |
| ---------- | -------- | -------- | ----------- |
| 1          | 100%     | 9        | ~ 159       |
| 2          | 198%     | 9        | ~ 92        |
| 5          | 337%     | 9        | ~ 180       |
| 10         | 571%     | 9        | ~ 120       |
| 20         | 617%     | 9        | ~ 120       |

In the above, each frame is rendered 200 times and the average
frame render time is recorded.  There was high variance between
runs in the 5x range.  I chose the fastest times for recording,
but the lack of performance consistency forces me to conclude
the setup is not to be completely trusted; though it does
represent a real world rig.

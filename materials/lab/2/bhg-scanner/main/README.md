In here? try > 
- go build
- time ./main

1)The code scans a number of ports and returns the number of open and closed ports based on the presence of error using the net.DialTimeout package.
2) TODO 1: method to build and run the code 
   TODO 2: replace the default timeout with a designated timeout 
   TODO 3: add close ports to track the closed ports along with the open ports
   TODO 4: used for tuning. the time for 1000 was much higher (600.08 s) compared to 1 (2.154 s)
   TODO 5: take a variable for the number of ports to scan and use comma to separate the output
   TODO 6: return the total number of open and closed ports scanned
3) The modification I chose was to use a variable which can be altered as an input from the user to assign the number of ports instead of using the fixed value of 1024.
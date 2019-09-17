# frog_problem
Here is the frog problem code I wrote in an Edinburgh pub.

Feel free to make better version. This one has not changed since I wrote it. V1
was the same code but saved before I changed it to work for different distances
so I have not bothered upload it as well (it was still very unifinished).

Video is here: https://youtu.be/ZLTyX4zL2Fc

## Iterative Rational Implmentation

Written in Golang by Taylor Wrobel

Makes use of big rationals in golang to get exact expectation for all number of
pads, up to and including the specified target number of pads (including target
bank).  Makes use of an interative approach, which performs better than its
recursive equivalent (calculating all results up to 1,000 pads takes ~12.5
seconds for me).

Outputs to STDOUT in CSV format with the following columns (output for `n=10`
displayed in table form):

n  | numerator | denominator | float64 representation
-- | --------- | ----------- | ----------------------
 1 |         1 |           1 |               1.000000
 2 |         3 |           2 |               1.500000
 3 |        11 |           6 |               1.833333
 4 |        25 |          12 |               2.083333
 5 |       137 |          60 |               2.283333
 6 |        49 |          20 |               2.450000
 7 |       363 |         140 |               2.592857
 8 |       761 |         280 |               2.717857
 9 |      7129 |        2520 |               2.828968
10 |      7381 |        2520 |               2.928968


Building: `go build`
Running: `./frog [n]` (`n` defaults to 10 if not provided)
Outputting to file (on a UNIX system): `./frog 100 > output.csv`
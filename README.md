go-batchrun
===========

A command-line tool for executing n-tasks concurrently


## Usage
    $ ./batchrun
    Usage ./batchrun: prog1 prog2 prog3...
        -concurrency=1: Concurrency
        -help=false: Help
        -logdir=".": Log directory

## Example
	$ ./batchrun -concurrency 2 -logdir=logs "./test.sh 1" "./test.sh 2" "./test.sh 3" "./test.sh 4"
	Starting task :  test.sh.0
	Starting task :  test.sh.1
	Completed task : test.sh.0 in 4.015887286s
	Starting task :  test.sh.2
	Completed task : test.sh.1 in 4.016462683s
	Starting task :  test.sh.3
	Completed task : test.sh.3 in 4.007310343s
	Completed task : test.sh.2 in 4.008505852s
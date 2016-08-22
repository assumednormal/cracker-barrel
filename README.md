# cracker-barrel

Package _cb_ provides a simple interface for playing the Cracker Barrel triangular peg board game. See main/main.go for example usage.

~~~~
$ go run main/main.go -n.games 10000 | sort | uniq -c
  76 1
 644 2
2487 3
4637 4
1447 5
 387 6
 319 7
   3 8
~~~~

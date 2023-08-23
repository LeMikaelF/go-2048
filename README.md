# 2048

Incomplete implementation of the game 2048, written with the 2 hours of laptop battery I had on a transatlantic flight. I had never played it and my only experience with 2048 was the in-flight game.

There's no merging, and no scoring, so far all that works is 2s appearing in blank squares, everything sliding correctly, and losing when the board is full. The UI is a barebones CLI UI.

Here's a sample of the gameplay:
```
Current grid.
0 0 0 2 
0 0 0 2 
2 0 0 0 
0 0 0 0 
Press any arrow, then Enter.
^[[D

Current grid.
2 0 0 0 
2 0 0 0 
2 0 0 2 
0 0 0 0 
Press any arrow, then Enter.
^[[C

Current grid.
0 0 0 2 
0 0 0 2 
0 0 2 2 
0 0 2 0 
Press any arrow, then Enter.
```

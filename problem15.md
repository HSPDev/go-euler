#This didn't need programming :-)

It is solved using combinatorics. I had to look that up, but I did do that formula in our Gymnasie.

To solve a 20 x 20 grid from top left to bottom left, only going left and down, we have 40 possible moves to do. 
We have to select which 20 to go left (or down, whatever you want).
In principle you could go 20 to the left, and then you could only go 20 down.
Or 19 to the left, then you would have 21 moves left, needing to do 20 down, and move 1 more left into it.

So it's combinatorics

We have total 40 moves.
We have to select 20 out of 40.

40! / ((40-20)! * (20!)) = 137846528820

Edu and Alberto are two brothers who work on a construction site. They frequently have to deal with complex mathematical challenges in their jobs. But they don't have any work today, so they're in charge of clearing the piles of sand that other workers have made during the week.

More precisely there are N piles numbered from 1 to N. The i-th (1 ≤ i ≤ N) pile stacks Vi m3 of sand. Fortunately, they can use the bulldozer of infinite buckets.

The bulldozer of infinite buckets is a huge machine that can be plugged into one bucket in an infinite set of buckets. The j-th (1 ≤ j ≤ +∞) bucket has a capacity Cj (the volume of sand it may carry) of 2j-1 m3. That means in the infinite set, the buckets have capacities that go: 1, 2, 4, 8, 16, 32…

The bulldozer

They could use the nth bucket to grab ∞ m3 of sand and finish fast, but Edu suggests that Alberto play a game instead. The game has the following rules.

They take turns removing a positive amount of sand from a single pile of sand. Edu makes the first move. Alberto goes next and then Edu makes another move, and so on.
On each turn a player can use any bucket in the infinite set of buckets to remove exactly Cj m3 of sand from any single pile. They have to fill whatever bucket they use completely. That means they can’t use the 4 m3 bucket to remove only 3 m3 of sand.
The player who removes all the remaining sand from the last non-empty pile wins the game. In other words, a player who can’t remove any sand in his turn loses the game.
If both brothers play optimally who will be the winner?

Input
The first line will have an integer T, which is the number of cases for the problem. It’s followed by a description of T cases. Every case has two lines. The first line of each case has an integer N, which is the number of piles of sand. The second line of each case has N integers: V1, V2, V3, … VN, which is the amount of sand in each pile.

Output
For each case, there should be a line starting with "Case #x: y", where x is the test case number (starting from 1) and y is “Edu” or “Alberto”, whoever wins the game.

Limits
T = 10 for test input and T = 100 for submission input.
1 ≤ N ≤ 1000
1 ≤ Vi ≤ 2500
Sample Input
3
1
8
1
3
2
1 1
Sample Output
Case #1: Edu
Case #2: Alberto
Case #3: Alberto
In the first test case, there is only one pile of 8 m3 of sand. Edu can use the 4th bucket to remove 8 m3 and win.
In the second test case, there is only one pile of 3 m3 of sand. Edu is only allowed to use the first or the second bucket to remove 1 m3 or 2 m3. He can’t use the 4 m3 bucket because he has to fill every bucket he uses. It doesn’t matter which bucket he uses. Alberto will clean the pile in the next turn.
In the third test case there are two piles. They are each 1 m3. Edu needs to remove the remaining sand from one of them and leave an empty pile and a 1 m3 pile. Alberto clears the other pile and wins.
Note that Edu can’t use the 2 m3 bucket to remove both piles in one step, because he can only remove sand from one pile at a time.

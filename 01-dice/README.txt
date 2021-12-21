Let's play a very simple dice game. In the game, two players throw two dice and the score is the sum of the dice throws. The player with the highest score wins the game. If both players get the same score the game is a draw.

Roll the dice!

Input
The first line will have an integer N, which is the number of cases for the problem. Each case has the result of the dice roll of the first player represented by two numbers (between 1 and 6) separated by ':'.

Output
For each case, there should be a line starting with "Case #x: " followed by the minimum score the second player needs to win the game, or '-' when it's impossible to win (there can only be a draw after scoring all 6 dice).

Sample Input
8
1:3
3:3
6:5
5:6
6:6
1:1
2:1
4:4
Sample Output
Case #1: 5
Case #2: 7
Case #3: 12
Case #4: 12
Case #5: -
Case #6: 3
Case #7: 4
Case #8: 9
__
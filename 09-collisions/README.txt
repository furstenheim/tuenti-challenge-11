We’re building the collision detection module for our 2D game engine. We want to be able to process a lot of sprites at the same time and detect all the collisions. We'll provide the bitmask for each sprite, where 0 indicates transparency and 1 indicates a solid part of the sprite. If two solid parts of different sprites share the same location then it's counted as a collision. If two sprites collide at different points a single collision is counted. A sprite can collide with multiple other sprites.

Given a set of sprites, return how many collisions are detected.


Limits
T ≤ 10
D ≤ 10
W, H ≤ 512
P ≤ 50000
X, Y ≤ 100000
Input
The first line will have an integer T, which is the number of cases for the problem. The next line has an integer D, which is the number of sprite definitions. D sprite bitmask definitions follow. Each sprite bitmask definition starts with a line with two integers W, H, which are the width and height of the sprite followed by H lines of length W of 0's or 1's. T test cases follow. Each case starts with an integer P, which is the number of sprite positions in the test. P lines follow and each line has 3 integers: I, X, Y, where I is the sprite identifier (from 0 to D-1), and X, Y are the coordinates of the sprite.

The (0,0) coordinate is the top-left corner of the display and sprites.

Output
For each case, there should be a line starting with "Case #x: " followed by the number of collisions for that test case.

Sample Input
3
2
16 8
0000100000100000
0000010001000000
0000111111100000
0001101110110000
0011111111111000
0010111111101000
0010100000101000
0000011011000000
4 4
0110
1111
1111
0110
2
0 0 0
1 2 2
5
0 0 2
1 0 2
1 5 0
1 11 2
1 12 8
6
0 0 2
1 7 0
1 1 2
1 10 2
1 12 5
1 2 2
Sample Output
Case #1: 1
Case #2: 0
Case #3: 6
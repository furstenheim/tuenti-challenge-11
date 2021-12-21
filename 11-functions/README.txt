Our team is writing a new JS utils library. The project has N functions. We want to split and distribute the functions in different files so every file exposes the same K number of functions. But we want to do it in a very specific way.

We have a linting tool that gives a score to our project and we want to maximize our score.

How is the score computed? The score of the project is the sum of the scores of its files.

How is the score for a file computed? The score for a file is the length of the longest common prefix of all the function names the file exposes. Let's see an example. A file that exposes the following three functions - getHeight, getWidth and getDepth - will score 3 points, because their longest common prefix is “get”, which has three characters

However, a file that exposes the following functions - pairs, invert and invoke - will score no points because the names have no common prefix.

So given a set of function names and K, the number of functions a file should contain, what would be the maximum score possible by spreading the functions into N/K files, with K functions in each file?

Note that you don’t need to say which functions go into which files. You just need to know the score they would have.

Input
The first line will have an integer T, which is the number of cases for the problem. It is followed by a description of T cases. Every case has a line with two integers N and K. There are N lines following that each have a function name as a string.

Output
For each case, there should be a line starting with "Case #x: y", where x is the test case number (starting with 1) and y is the maximum score that can be achieved.

Limits
T = 100
1 ≤ K ≤ N ≤ 1000
N mod K == 0
function names length ≤ 50
Sample Input
1
6 3
getHeight
pairs
getWidth
invert
getDepth
invoke
Sample Output
Case #1: 3
In the first test case, the best arrangement would be to put the getHeight, getWidth, getDepth functions into one file, which gives 3 points. And then put the rest of the functions in another file, which would not add any points because they have no common prefix.


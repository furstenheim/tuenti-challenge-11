This is an interactive problem. You can talk with the judge at codechallenge-daemons.0x14.net:7162

The judge will generate a random permutation A of N elements when the connection is created. A permutation of N elements is a sequence of N different elements that contains all elements from 1 to N exactly once. [1, 2, 4, 3] is example of a permutation of 4 elements, while [2, 3, 4] or [1, 1, 2] are not permutations.

Permutation A is kept secret from you. But you can ask the server to tell you the result of gcd(Aa, Ab), given two indices a and b (a ≠ b). Here, gcd(X, Y) means the Greatest Common Divisor of X and Y.

Your task is to find the indices of the array (from 1 to N) that contain numbers with no more than 2 divisors in no more than Q queries.

Communication with the judge
The first message will be from the judge. It will be a single line with two integers that are the values of N and Q, which are the maximum number of queries you may send. After that, you can send up to Q messages in the following format.

“? a b”: Asks the judge what the value of gcd(Aa, Ab) is. The two positions a and b must be different and must contain 1 ≤ a ≤ N and 1 ≤ b ≤ N. Otherwise, the judge will reply with an error message and the connection will be closed.
“! p1 p2 p3 …": Gives the answer to the judge. These are positions in the array that contain a number with no more than 2 divisors. If your answer is wrong, the judge will reply with an error message and the connection will be closed. If your answer is correct the judge will respond with the password.
If you make more than Q queries the server will reply with an error message and the connection will be closed.

Sample Connection
Imagine the judge generates the following permutation of N = 8 (the numbers with no more than 2 divisors are highlighted).

[8, 4, 6, 3, 2, 7, 5, 1]
A possible conversation between a player and the judge would be:

Judge's messages	Player's messages	Explanation
8 15		Judge specifies N = 8 and Q = 15
? 1 2	Player asks gcd(A1, A2). Queries made: 1
4		The judge responses with gcd(8, 4) = 4 Now the player knows that A1 and A2 both have more than 2 divisors. They can be divided by 1, 2 and 4. More precisely, she knows the numbers are 4 and 8, in some order. She needs to keep asking.
? 3 4	Player asks gcd(A3, A4). Queries made: 2
3		The judge responses with gcd(6, 3) = 3 The player knows both A3 and A4 are divisible by 1 and 3. She needs to keep asking.
? 3 5	Player asks gcd(A3, A5). Queries made: 3
2		The judge responses with gcd(6, 2) = 2 The player was lucky. Now she knows that A3 is divisible by 1, 2 and 3. It must be 6.
! 4 5 6 7 8	She knows where 4, 6 and 8 are. The rest of the numbers in the permutation have 2 or fewer divisors. Queries made: 4
Congratulations you're the winner!		The judge responds with the password
Limits
N = 100
Q = 1500

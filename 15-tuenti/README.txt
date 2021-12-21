This is bad! After we wrote a problem about the Tuentistic Numbers, the non-Tuentistic numbers became angry and invaded our offices! They want you to show their strength, so you need to calculate the product of multiplying all the non-Tuentistic numbers from 1 to the given n, modulo 1e8+7.

Remember that a Tuentistic number is any number that, when written in English, contains the word “twenty” (for example, 20, 21 or 120000).

Input
The first line has the number of cases C. Then C lines follow, each with a number N.

Output
Output 'Case #X: P' for each case, where X is the case number (the first case has number 1) and P is the product of multiplying all the non-Tuentistic numbers from 1 to N, modulo 1e8+7 (100000007).

Limits
1 ≤ C ≤ 1000
1 ≤ N ≤ 262
Sample Input
3
19
25
100000
Sample Output
Case #1: 93675574
Case #2: 93675574
Case #3: 46066740
Explanation
In the first case, there are no Tuentistic numbers lower than 20. So, you just multiply all the numbers up to 19 and do the modulus.
In the second case, the result is the same as in the first case since all the numbers from 20 to 25 are Tuentistic
This is a famous cryptarithmetic puzzle. To solve it, you need to replace each letter with a different digit

SEND + MORE = MONEY

Input
The first line has an integer N, which is the number of cases for the problem. Each case is a cryptarithmetic puzzle. The puzzle can have multiple solutions. Arithmetic operations can use the following symbols: {+, -, *, /, =}.

Output
For each case, there should be a line starting with "Case #x: " followed by the semicolon separated list of solutions to the puzzle or IMPOSSIBLE if the puzzle doesn't have a solution. If the puzzle has multiple solutions order them in lexicographic order.

A number cannot start with ‘0’. So, the following numbers are invalid: 01, 012, 0457,...

Sample Input
5
SEND + MORE = MONEY
TWELVE / SIX = TWO
FALSE = TRUE
ABCD * 4 = DCBA
DUCK + DUCK = GOOSE
Sample Output
Case #1: 9567 + 1085 = 10652
Case #2: 160380 / 972 = 165;210370 / 965 = 218;340170 / 986 = 345
Case #3: IMPOSSIBLE
Case #4: 2178 * 4 = 8712
Case #5: 8327 + 8327 = 16654;8345 + 8345 = 16690;9435 + 9435 = 18870;9436 + 9436 = 18872

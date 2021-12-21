You've been contacted by a company called Awesome Sales Inc. They recently bought some new ultra-modern AI based software for creating travel plans for their sales force. But, it seems to have a huge problem and the developers have gone MIA. So they need you to figure out how to fix it.

The "awesome" software works with two modules. The first module (AICityPicker) picks the best cities to be visited by each employee based on a number of the employee’s records and traits. It also takes another important factor into account, which is that employees can move between cities on an ASCB (Awesome Sales Connection Bus).

Awesome Sales Connection Bus

Once the cities have been picked by the AI, a plan is generated in the form of bus tickets that can be used multiple times to move between cities. For example, t1: [A, B] is a ticket that can be used to go from city A to city B and vice versa. In the plans, employees can always reach one city from another using their tickets. For example, if an employee has to visit cities A, B and C, there will always be a way to go from A to C, even if they have to go from A to B and then from B to C.

The other module (AICityRemover) is the one with the problems. Once a plan is generated, this module evaluates budgets, costs and other things and picks cities from the initial plan to remove. It cancels the bus tickets for them and everything else – all automatically.

The biggest problem is that this module doesn't take the ASCB into account. Cities are removed and employees can no longer reach all their destinations. For example, if an employee has to visit A, B and C, the module may remove B, even if none of the ASCB can go from A to C directly. So, when an employee finds themselves in that situation they have to find a way to reach their destination. Sometimes they pay for a regular bus ticket or taxi out of their own pocket or they might even ask a stranger for a ride. Not good.

Of course, you don't have access to the source code. But, you manage to reverse engineer the AICityRemover module and find the root cause of the problem. A function called currentPlan.getCriticalCitiesThatCannotBeRemoved() wasn't implemented and always returns an empty list! Luckily, the documentation has some examples so you can work out your own implementation.

You need to write the getCriticalCitiesThatCannotBeRemoved() function so that for each plan it returns a list of cities that can't be removed. In other words, the list of cities where if any city is removed the initial plan breaks and the employee can't reach the rest of the cities using their ASCB tickets.

Cities are removed one by one, so when a city is removed from the initial plan, other cities might become critical now but that would be a whole new case. You need to provide the list of cities that are critical from the initial plan.

Input
The first line will have an integer C, which is the number of cases for the problem. The description of the C cases follows. For each case, the first line is an integer T showing the number of tickets. Then T lines follow, each describing a ticket. Each ticket is described in the form "City A,City B", which means the ticket can be used to go from City A to City B and vice versa.

Output
For each case, there should be a line starting with "Case #x: y", where x is the test case number (starting from 1) and y is the list of the critical cities, in the form "City A,City B". The list should be ordered alphabetically. If none of the cities are critical a dash (-) should be printed instead.

Limits
1 ≤ C ≤ 100
1 ≤ T ≤ 1500
Sample Input
3
2
Chicago,Houston
Houston,Dallas
4
Madrid,Barcelona
Barcelona,Zaragoza
Zaragoza,Valencia
Valencia,Madrid
6
París,Lyon
Lyon,Toulouse
Toulouse,Lille
Toulouse,Niza
Nantes,Niza
Niza,Lille
Sample Output
Case #1: Houston
Case #2: -
Case #3: Lyon,Niza,Toulouse
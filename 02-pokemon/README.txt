Something strange has happened. All your pokemons have escaped. Can you catch them all?


You've been given a map (shown with rows and columns) so you can find all the Pokémon. The Pokémon's name can only be found horizontally. But it can be written from left to right or right to left and it can be on more than one line. Once you catch a Pokémon make sure to remove it from the map. Because each Pokémon only appears once and some Pokémon can be hidden within other Pokémon.

Input
The first line will have an Integer N, which is the number of cases for the problem. It is followed by a description of T cases. Every case has an Integer P, which is the number of Pokémon to find, an Integer R, which is the number of rows, and an Integer C, which is the number of columns. Then it has P lines with the names of each Pokémon N. And finally R lines with C characters split by an empty space.

Output
For each case, there should be a line starting with "Case #x: " followed by the result of the map without the Pokémon.

Limits
1 ≤ T ≤ 20
1 ≤ P ≤ 50
1 ≤ C, R ≤ 100
1 ≤ N ≤ 100
Sample Input
2
1 4 6
SNORLAX
T A K E C A
S N O R L A
X R E W I T
H V E N O M
2 3 10
PIKACHU
CHARIZARD
N O P O K E M U H C
A K I C H A R I Z A
R D P O N S H E R E
Sample Output
Case #1: TAKECAREWITHVENOM
Case #2: NOPOKEMONSHERE
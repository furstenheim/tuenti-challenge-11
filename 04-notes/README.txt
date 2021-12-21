Did you know there are really 12 notes and not 7? Don't forget the sharps and flats. Those are the black keys on a piano or keyboard. A musical scale is just a series of 7 of the 12 notes. Now, we're going to build some musical scales with different roots.

If we use the English notation notes are represented from A to G (A=La, B=Si, C=Do, etc.). You should also know that the note following most notes, except B and E, is the sharp, for example A and A#. But every sharp note can also be flat for the note after it. So A# can also be called Bb. So, the sequence of notes from A to B can be A-A#-B or A-Bb-B. There is one special rule. Some version of every note must be included in the scale , either natural, flat or sharp. So in some circumstances a modified version of B and E could also be included in the scale.

The root of a musical scale is the first note of the scale. So, if you have a C scale the first note will be C, and so will the 8th one. A sequence of semitone and tone jumps is defined when creating a scale. A semitone jump is just the next note in the sequence of all notes. A full tone jump means 2 semitones. So, you jump 2 notes to find the next note in the scale. For instance, the major scale is built with the sequence tone-tone-semitone-tone-tone-tone-semitone (TTsTTTs). So, the C major scale has the notes CDEFGABC. There are no sharps or flats.

C Major scale

But what happens when a sharp or flat is included in the scale? Should we use the sharp or the flat? It's easy. Note names can only appear once in any given scale. So, if you already have A, for example, you cannot also have A#, so Bb is used.

Input
The first line will have an integer T, which is the number of cases for the problem. It’s followed by a description of T cases. Every case has two lines. The first line of each case has the root of the scale to generate. It can be a simple note or a modified note (sharp or flat). The second line of each case has a string of 7 characters describing the scale steps. "T" means a tone and "s" means a semitone.

Output
For each case, there should be a line starting with "Case #x: y", where x is the test case number (starting with 1) and y is a string with the notes of the scale.

Sample Input
3
G
TTTsTTs
A
sTTsTTT
A
TTsTTsT
Sample Output
Case #1: GABC#DEF#G
Case #2: ABbCDEbFGA
Case #3: ABC#DEF#GA
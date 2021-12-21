Challenge 19 - Code Red Chaos
Good morning/afternoon/evening/night, Agent 0x4C11DB7:

We're contacting you because we're in dire need of your help. As you most likely know, we've been sabotaging the enemy's network for some time now. But, we've been detected.

We've been able to modify their CRC polynomials with any one of our choosing. And since we know the standard interference in their network, we've been supplying them with special CRCs that are unable to detect that interference. That’s resulted in huge gains for us, because it slows down their network and makes them waste effort and resources.

However, they've implemented new features that detect and flag our forged CRC polynomials. We need you to lend us your extensive expertise with CRCs to help the Agency stay ahead of the enemy for a little while longer.

As you know, the interference consists of bit strings that are XOR'd into the message at any place. A single message may have any amount of interference at any place, as long as all the interference is inside the message.

For example, if we detect that a channel usually has interference 0b101 and 0b1001, a 0b0101010101010101 message could be transformed by the channel into any of

0b0101010101010101
0b0101010101000001
0b0101010101111101
0b0101010101110111
0b0101010101101001
0b0101010101001101
but never into
0b0101010101010100
Fortunately for us, we don't need to deal with any reflections in the input or output or preconditioning.

As you explained last time, any CRC polynomial consisting of a 1 followed by any number of 0 bits won't detect those errors. But, unfortunately, now they need the polynomial to end in a 1 bit. Also, to maximize the time these polynomials remain undetected, we need them to be as long as possible (counting the number of bits in their binary representation).

We kindly return the tool you gave us for testing the CRCs with binary strings:

Message in binary	Full poly in binary	Result
Interference in binary
Message with interference		Result of message with interference
Be swift, Agent, you are our only hope!

Input
You get a file that starts with a number of cases followed by the cases themselves.
Each case starts with the quantity of interference and the interference parts follow. Each interference is given as a lowercase HEX string for easier processing. Ignore the leading 0 bits. However, the trailing 0 bits are relevant.

Output
For each case, output

Case #N: Poly
where N is the unpadded index of the case and Poly is the FULL polynomial, also in lowercase HEX

Limits
50 cases for test
100 cases for submit
1 <= Number of interferences for case  <= 10
1 <= Interference Bit Length <= 16 for test
1 <= Interference Bit Length <= 128 for submit
Sample input
5
1
5
2
5
9
2
4
8
2
ab
cd
2
aa
bb
Sample output
Case #1: 5
Case #2: 3
Case #3: 1
Case #4: 1
Case #5: 11
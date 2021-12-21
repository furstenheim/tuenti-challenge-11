We all know that cryptocurrencies is a market that is growing exponentially. That's why your friend Jonathan got the idea of trying to scrape several crypto trading websites to find a bug he can profit from.

His idea is to find out whether you can increase the amount of the coins you initially invested by exchanging different cryptocurrencies between different websites. Your friend Jonathan is a very good programmer, but algorithms aren't his strong suit. That's why he's asking for your help.

Crypto Bubble

He agrees to create a file with all the cryptocurrency exchange information from the different websites. And he's asking you to develop an algorithm to figure out whether it's possible to get rich by exchanging coins.

Jonathan only has one BTC and he wants to get more. Exchanges must start and end in BTC, but any other cryptocurrency can be used in the intermediate steps. Exchanging a cryptocurrency can take a very long time, so you should try to do as few exchanges as possible. If there are multiple paths with the same number of exchanges, you should choose the most profitable one. In other words, you need to find the shortest exchange loop that gives you more bitcoin than when you started even if another longer loop may be more profitable.

Input
After some tough negotiating you reach an agreement on what the file format will be like.

The first line has the number N of different cases
The N cases start after that line:
Each case starts with a number M, which is the number of websites in the case.
Then the websites datai starts:
Each website data block starts with a line that has the website name (a string of characters without white space) and K (the number of available trades).
The K available trades are described after the name.
Each trade is declared on one line with the crypto acronym followed by a dash then an integer D, another dash and finally the other crypto acronym.
An example trade line would be: BTC-2-ETH
What does a trade line mean? In the example above one BTC is exchanged for two ETH. But it doesn't mean two ETH can be exchanged for one BTC.
Output
For each case, there should be a line starting with "Case #N: Z", where N is the case number and Z is a number. Z must be the final amount of BTC you end up with after doing one iteration on the chosen path. If there are no paths that generate a profit the final amount must be the starting amount, in this case “1”.

Limits
1 ≤ N ≤ 100
1 ≤ M ≤ 100
0 ≤ K ≤ 100
0 ≤ D ≤ 20
1 ≤ Z ≤ 1000000000
Sample Input
4
2
coinscoinsmarket 6
BTC-1-ETH
ETH-1-BTC
ETH-1-BTC
ETH-1-BNB
BNB-1-ETH
BNB-1-BTC
crytotome 1
BTC-2-BTC
4
coinscoinsmarket 2
BTC-1-ETH
ADA-1-BNB
crytotome 2
ETH-1-XRP
BNB-1-BTC
yetanothercoin 1
XRP-1-DOT
givemecoins 1
DOT-1-ADA
4
coinscoinsmarket 2
BTC-1-ETH
ADA-1-BNB
crytotome 2
ETH-1-XRP
BNB-1-BTC
yetanothercoin 1
XRP-0-DOT
givemecoins 1
DOT-1-ADA
4
coinscoinsmarket 2
BTC-1-ETH
ADA-1-BNB
crytotome 2
ETH-1-XRP
BNB-1-BTC
yetanothercoin 1
XRP-1-DOT
givemecoins 2
DOT-1-ADA
ADA-2-BNB
Sample Output
Case #1: 2
Case #2: 1
Case #3: 1
Case #4: 2
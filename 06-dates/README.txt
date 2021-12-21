Have you ever wondered what day of the week a date is? I'm sure you have. In this challenge, you have to code a program to figure out how to tell anybody the day of the week (from a list of languages) for a given date.

What day is it?
Input
The first line has an integer N, which is the number of cases for the problem. Each case has a date, the character ':' and a two-letter language code.

The date format can be: YYYY-MM-DD or DD-MM-YYYY
All dates are after 1970-01-01
Input cases may have invalid dates, like an invalid month (MM>13) or an invalid day of a month
The weekday of a date can be given in any of these twenty languages:
CA: catalan	CZ: czech	DE: german	DK: danish	EN: english
ES: spanish	FI: finnish	FR: french	IS: icelandic	GR: greek
HU: hungarian	IT: italian	NL: dutch	VI: vietnamese	PL: polish
RO: romanian	RU: russian	SE: swedish	SI: slovenian	SK: slovak
Output
For each case, there should be a line starting with "Case #x: " followed by the weekday IN LOWERCASE of the day in the indicated language, or 'INVALID_DATE' if given date is invalid, or 'INVALID_LANGUAGE' for an unknown language code.

Sample Input
12
2021-04-01:ES
2021-04-07:ES
01-02-2021:GR
2021-02-07:VI
2021-02-01:DE
01-02-2021:EN
01-02-2021:XX
29-02-2021:FR
15-20-2020:EN
2020-02-29:RU
01-02-2021:CA
2021-02-01:CZ
Sample Output
Case #1: jueves
Case #2: miércoles
Case #3: δευτέρα
Case #4: chủ nhật
Case #5: montag
Case #6: monday
Case #7: INVALID_LANGUAGE
Case #8: INVALID_DATE
Case #9: INVALID_DATE
Case #10: суббота
Case #11: dilluns
Case #12: pondělí

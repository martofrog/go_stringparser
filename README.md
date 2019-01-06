# stringparser
	 
We feed our odds from many different sources, we receive them as strings and manipulate them.

A correct odds string is for example "5/2" which corresponds to 3.5 decimal odds.
Unfortunately our feeds are not always reliable and the string can contain literally anything.
Aim of this task is to provide a small and solid module which parses a string and returns an a valid Odds
object if possible.

fios-finder
=

I was tired of having to fight with Verizon customer service reps to see if "there was service" in an area.
In fact, I wanted to be able to know which cities in a state had FIOS service. I noticed on their 
"availability" tool they only required a zip code. So, I decided to write a script to find the information
that the reps wouldn't give me and now I'm sharing it with you.

Enjoy.

If there are problems with this then I probably wont be able to fix them in a timely manner. Feel free to shoot
me PR requests.

Enjoy.

Zip Code Dataset
==

I used, and setup the code to work with https://www.aggdata.com/node/86 which is a free zipcode set from AggData.
You can modify the CSV struct to support whatever you prefer, it must have a zip code, a city, and a state/state abbr.


Notes
==

The script is not exceptionally fast. It could be made faster by hunting through the data better (since VZN returns
a larger set of data) or by storing the results in some datasource (filesystem or DB). But, this serves my purpose at
the moment.

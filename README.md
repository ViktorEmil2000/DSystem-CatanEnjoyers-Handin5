# DSystem-CatanEnjoyers-Handin5

To run this program you need to use 3 or more individual terminals
You need to run 

server/main.go once to set up our backend

from there you can add you're desired amount of clients/bidders

client/main.go


To see the code that Artificially crashes our main node you need to look in server/main.go
there you will find a kill func starting on line 71

The Auction is set run 100 seconds with the "main" node dying after 40

To clearly see the Replication working try to add more clients after the initial server crashes
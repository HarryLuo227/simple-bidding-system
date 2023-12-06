#!/bin/bash
# Our custom function
cust_func(){
  random_number=$((1 + RANDOM % 100))
  echo $1 $random_number
  curl -X PUT 'http://127.0.0.1:8080/bid/1' -F "latest_bid_price=$random_number"
  sleep 1
}
# For loop 5 times
for i in {1..20}
do
	cust_func $i & # Put a function in the background
done
 
## Put all cust_func in the background and bash 
## would wait until those are completed 
## before displaying all done message
wait 
echo "All done"
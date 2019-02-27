# Commands

## Cow stats 
```
redis-cli -r 10 -i 2 lrange cow-75998f888c-htfnp 0 999999
```
## Dairy stats 

All cows with milk produced
```
redis-cli --raw hgetall dairy|sed 'N;s/\n/,/'
```

Sum
```
redis-cli --raw hgetall dairy|sed 'N;s/\n/,/'|awk -F, '{s+=$2} END {print s}'
```
import redis
import json
import math

# Connection to Redis
redis_client = redis.StrictRedis(
    host='130.245.136.34', 
    port=6379, 
    password='66cfe2a89ba30e1a6c706756', 
    db=0, 
    decode_responses=True
)

# Connect to DB 1 for publishing the response
redis_db1 = redis.StrictRedis(
    host='130.245.136.34',
    port=6379,
    password='66cfe2a89ba30e1a6c706756',
    db=1,
    decode_responses=True
)

# Function to calculate the prime factors of a large number
def prime_factors(n):
    i = 2
    factors = []
    while i * i <= n:
        if n % i:
            i += 1
        else:
            n //= i
            factors.append(str(i))
    if n > 1:
        factors.append(str(n))
    return factors

# Subscribe to channel "hw3" on DB 0
pubsub = redis_client.pubsub()
pubsub.subscribe('hw3')

print("Listening for messages on channel 'hw3'...")

# Backend process loop
for message in pubsub.listen():
    if message['type'] == 'message':
        # Extract the content of the message
        data = json.loads(message['data'])
        bignum = data['BIGNUM']
        response_channel = data['CHANNEL']
        
        # Calculate prime factors of BIGNUM (convert it from string to integer)
        big_number = int(bignum)
        factors = prime_factors(big_number)

        # Prepare the response
        response = {
            "factors": factors,
            "value": bignum
        }
        
        # Publish the response to the specified channel on DB 1
        redis_db1.publish(response_channel, json.dumps(response))

        print(f"Sent prime factors of {bignum} to channel {response_channel}")

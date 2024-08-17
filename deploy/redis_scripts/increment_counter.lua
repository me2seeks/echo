local current = redis.call('GET', KEYS[1])
if not current then
    redis.call('SET', KEYS[1], 1)
    return 1
else
    redis.call('INCR', KEYS[1])
    return current + 1
end

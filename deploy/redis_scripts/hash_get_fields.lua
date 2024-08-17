local result = {}
for i=1, #ARGV do
    table.insert(result, redis.call('HGET', KEYS[1], ARGV[i]))
end
return result

local keys = redis.call('KEYS', ARGV[1])
for i=1,#keys,5000 do
    redis.call('DEL', unpack(keys, i, math.min(i+4999, #keys)))
end
return #keys

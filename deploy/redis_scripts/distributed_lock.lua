if redis.call("SETNX", KEYS[1], ARGV[1]) == 1 then
    redis.call("PEXPIRE", KEYS[1], ARGV[2])
    return 1
else
    return 0
end

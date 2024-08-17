local message = ARGV[1]
if string.match(message, "keyword") then
    redis.call('PUBLISH', KEYS[1], message)
end

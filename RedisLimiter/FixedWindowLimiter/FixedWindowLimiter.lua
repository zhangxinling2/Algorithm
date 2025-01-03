local limit=ARGV[1]
local ttl=ARGV[2]
local counter = tonumber(redis.call("get",KEYS[1]))
if counter==nil then
    counter=0
end
if counter>=limit then
    return 0
end

redis.call("incr",KEYS[1])
if counter==0 then
    redis.call("pexpire",KEYS[1],ttl)
end
return 1

redis.pcall("remrangeByScore",KEYS[1],0,ARGV[1])
local res=redis.pcall("zcard",KEYS[1])
if (res==nil) or (res<tonumber(ARGV[3])) then
    redis.call("zadd",KEYS[1],ARGV[2],ARGV[4])
    redis.call("pexpire",KEYS[1],ARGV[2])
    return 1
end
else return 0 end
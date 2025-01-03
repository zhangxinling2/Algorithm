-- ARGV[1]: 最高水位
-- ARGV[2]: 水流速度/秒
-- ARGV[3]: 当前时间（秒）
local peakLevel=tonumber(ARGV[1])
local currentV=tonumber(ARGV[2])
local now=tonumber(ARGV[3])
local lastTime=tonumber(redis.call("hget",KEYS[1],"lastTime"))
local curLevel=tonumber(redis.call("hget",KEYS[1],"curLevel"))
--初始化
if lastTime==nil then
    lastTime=now
    curLevel=0
    redis.call("hmset",KEYS[1],"lastTime",now,"curLevel",0)
end
--尝试放水
local interval =lastTime-now
if interval>0 then
    local newLevel=curLevel-currentV*interval
    if newLevel <0 then
        newLevel=0
    end
    redis.call("hset",KEYS[1],"curLevel",newLevel)
end
if curLevel>=peakLevel then
    return 0
end
--增加水位
redis.call("hincrby",KEYS[1],"curLevel",1)
redis.call("expire",KEYS[1],peakLevel/currentV)
return 1
---------------------  redis debugger setting  ------------------
local redis = require 'redis'
local host = "127.0.0.1"
local port = 6379
client = redis.connect(host, port)
redis.call = function(cmd, ...)
    return assert(loadstring('return client:'.. string.lower(cmd) ..'(...)'))(...)
end
-----------------------  methods theme  ------------------------
local function redisOperator(a, b)
    redis.call('set', 'a', 'a')
    redis.call('set', 'b', 'b')
    redis.call('set', 'lua:a', a)
    redis.call('set', 'lua:b', b)
    local keys = redis.call('keys', '*')
    return keys
end

----------------------  execute theme  --------------------------
local a = 2
local b = 33
local result = redisOperator(a, b)
print(table.concat(result,', '))

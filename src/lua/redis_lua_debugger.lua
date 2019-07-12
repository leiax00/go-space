---------------------  redis debugger setting  ------------------
local redis = require 'redis'
local host = "127.0.0.1"
local port = 6379
client = redis.connect(host, port)
client:auth('root')  -- 鉴权语句
redis.call = function(cmd, ...)
    return assert(loadstring('return client:'.. string.lower(cmd) ..'(...)'))(...)
end
-----------------------  methods theme  ------------------------
local function redisOperator(a, b)
    -- redis-lua中支持的语法，但在redis中还不支持
    --client:set('a', 'a')
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

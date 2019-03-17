local cjson = require 'cjson'
local jsonStr = '{"a": "a", "b": {"b1": "b1"}, "c": "c"}'

local jsonObj = cjson.decode(jsonStr)

print(cjson.encode(jsonObj))
local stepNum = 10000
local function build(str)
    local n, c, s = math.floor(#str / 4), {}, 1
    if #str % 4 ~= 0 then
        c[n + 1], s = tonumber(string.sub(str, 1, #str % 4)), #str % 4 + 1
    end
    for i = n, 1, -1 do
        c[i], s = tonumber(string.sub(str, s, s + 3)), s + 4
    end
    return c
end

local function get(c)
    local r = ""
    for i = #c, 1, -1 do
        if tonumber(c[#c]) < 0 and tonumber(c[i]) > 0 then
            local v = math.pow(10, #c[i]) - tonumber(c[i])
            r = string.sub(r, 1, #r - 1) .. tostring(tonumber(string.sub(r, #r)) - 1) .. tostring(v)
        else
            if r ~= '' or tonumber(c[i]) ~= 0 then
                r = r .. c[i]
            end
        end
    end
    return r ~= '' and r or '0'
end

local function add(a, b)
    local c, d = {}, 0
    a, b = build(a), build(b)
    local count = math.max(#a, #b)
    for i = 1, count, 1 do
        local t = d + (a[i] or 0) + (b[i] or 0)
        c[i], d = tostring(t), 0
        if i ~= count then
            c[i], d = tostring(t % stepNum), math.floor(t/ stepNum)
        end
    end
    return get(c)
end

local function sub(a, b)
    local c, d = {}, 0
    a, b = build(a), build(b)
    local count = math.max(#a, #b)
    for i = 1, count, 1 do
        local t = d + (a[i] or 0) - (b[i] or 0)
        d = 0
        if i ~= count and t < 0 then
            t, d = t + stepNum, -1
        end
        c[i] = tostring(t)
    end
    return get(c)
end

local function multi(a, b)

end

local function div(a, b)

end

local c = add("1234567891234567890", "1234000000000000220")
print(c)
print(add('1200', '130000000'))
print(add('1200', '1300'))
print(sub('1200', '1300'))
print(sub('1300', '1200'))
print(sub('123456789', '123456789'))
print(sub('123456789', '1234567890123'))
print(sub('1234567890123', '123456789'))

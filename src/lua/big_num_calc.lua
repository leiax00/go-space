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
        if r ~= '' or tonumber(c[i]) ~= 0 then
            r = r .. c[i]
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
        c[i], d = i ~= count and tostring(t % 10000), math.floor(t/10000) or tostring(t), 0
    end
    return get(c)
end

local function sub(a, b)
    local c, d = {}, 0
    a, b = build(a), build(b)
    local delta, aMax = math.abs(#a - #b), #a > #b
    for i = math.max(#a, #b), 1, -1 do
        local aa = aMax and (a[i] or 0) or (a[i - delta] or 0)
        local bb = aMax and (b[i - delta] or 0) or (b[i] or 0)
        local t = d + aa - bb
        if i ~= 1 and t < 0 then
            t, d = t + math.pow(10, #tostring(aa)), -1
        else
            d = 0
        end
        c[i] = tostring(t)
    end
    return get(c)
end

local function multi(a, b)

end

local function div(a, b)

end

--local c = add("1234567891234567890", "1234000000000000220")
--print(c)
print(add('1200', '130000000'))
--print(add('1200', '1300'))
--print(sub('1200', '1300'))
--print(sub('1300', '1200'))
--print(sub('123456789', '123456789'))
--print(sub('123456789', '1234567890123'))
--print(sub('1234567890123', '123456789'))

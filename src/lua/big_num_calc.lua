local function build(str)
    local n, c = math.floor(#str / 4), {}
    if #str % 4 ~= 0 then
        c[n + 1] = tonumber(string.sub(str, 4 * n + 1, #str))
    end
    for i = 1, n, 1 do
        c[i] = tonumber(string.sub(str, 4 * i - 3, 4 * i))
    end
    return c
end

local function get(c)
    local r = ""
    for i, t in pairs(c) do
        if tonumber(c[1]) < 0 and tonumber(t) > 0 then
            local v = math.pow(10, #t) - tonumber(t)
            r = string.sub(r, 1, #r - 1) .. tostring(tonumber(string.sub(r, #r)) - 1) .. tostring(v)
        else
            if r ~= '' or tonumber(t) ~= 0 then
                r = r .. t
            end
        end
    end
    return r ~= '' and r or '0'
end

local function add(a, b)
    local c, d = {}, 0
    a, b = build(a), build(b)
    local delta, aMax = math.abs(#a - #b), #a > #b
    for i = math.max(#a, #b), 1, -1 do
        local aa = aMax and (a[i] or 0) or (a[i - delta] or 0)
        local bb = aMax and (b[i - delta] or 0) or (b[i] or 0)
        local t = d + aa + bb
        if i ~= 1 and #tostring(t) > #tostring(math.max(aa, bb)) then
            c[i], d = string.sub(tostring(t), 2), 1
        else
            c[i], d = tostring(t), 0
        end
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

local c = add("1234567891234567890", "1234000000000000220")
print(c)
print(add('1200', '130000000'))
print(add('1200', '1300'))
print(sub('1200', '1300'))
print(sub('1300', '1200'))
print(sub('123456789', '123456789'))
print(sub('123456789', '1234567890123'))
print(sub('1234567890123', '123456789'))

--2468567891234568110
--130001200
--2500
--9900
--100
--000000000
--765555566666
--1234444433334
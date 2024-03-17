local json = require("dkjson")

local function readjson(path)
	local file = io.open(path)
	if file == nil then
		return nil
	end
	local content = file:read("*a")
	local tab = json.decode(content)
	return tab
end

local function timeToMinutes(timeStr)
	local hours, minutes = timeStr:match("(%d+):(%d+)")
	return tonumber(hours) * 60 + tonumber(minutes)
end
local function formatTime(minutes)
	local hours = math.floor(minutes / 60)
	local remainderMinutes = math.floor(((minutes / 60) % 1) * 60)
	if hours == 0 then
		return string.format("%02d", remainderMinutes) .. " minutes"
	else
		return tostring(hours) .. ":" .. string.format("%02d", remainderMinutes) .. " hours"
	end
end

local function getClosestPrayer(currentTime, times)
	local currentMinutes = timeToMinutes(currentTime)
	local nextPrayer = nil
	local closestDifference = math.huge
	for prayer, time in pairs(times) do
		local prayerTimeMinutes = timeToMinutes(time)
		local difference = prayerTimeMinutes - currentMinutes
		if difference > 0 and difference < closestDifference then
			closestDifference = difference
			nextPrayer = prayer
		end
	end
	return nextPrayer, formatTime(closestDifference)
end

local function refresh()
	local timeTable = readjson("/home/joe/.local/share/timings.json")
	local currentTime = os.date("%H:%M")
	local prayer, time = getClosestPrayer(currentTime, timeTable)
	print(string.format("<b>%s in %s</b>", prayer, time))
end

refresh()

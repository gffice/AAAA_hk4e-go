-- 基础信息
local base_info = {
	group_id = 133106619
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 619001, gadget_id = 70500000, pos = { x = -795.615, y = 296.394, z = 2064.395 }, rot = { x = 0.000, y = 113.672, z = 0.000 }, level = 36, point_type = 2028, area_id = 19 },
	{ config_id = 619002, gadget_id = 70500000, pos = { x = -793.738, y = 295.916, z = 2061.913 }, rot = { x = 350.376, y = 333.127, z = 349.379 }, level = 36, point_type = 2028, area_id = 19 },
	{ config_id = 619003, gadget_id = 70500000, pos = { x = -643.366, y = 352.930, z = 2059.879 }, rot = { x = 335.155, y = 311.794, z = 11.679 }, level = 36, point_type = 2004, area_id = 19 },
	{ config_id = 619004, gadget_id = 70500000, pos = { x = -669.934, y = 355.955, z = 2056.023 }, rot = { x = 25.519, y = 128.699, z = 346.950 }, level = 36, point_type = 2004, area_id = 19 },
	{ config_id = 619005, gadget_id = 70500000, pos = { x = -692.802, y = 361.113, z = 2059.332 }, rot = { x = 22.273, y = 129.693, z = 349.999 }, level = 36, point_type = 2001, area_id = 19 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { 619001, 619002, 619003, 619004, 619005 },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================
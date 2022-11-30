-- 基础信息
local base_info = {
	group_id = 133308374
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
	{ config_id = 374001, gadget_id = 70500000, pos = { x = -1827.105, y = 118.412, z = 4850.343 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, point_type = 1002, area_id = 26 },
	{ config_id = 374002, gadget_id = 70500000, pos = { x = -1829.054, y = 117.749, z = 4848.865 }, rot = { x = 0.000, y = 297.730, z = 0.000 }, level = 32, point_type = 1002, area_id = 26 },
	{ config_id = 374003, gadget_id = 70500000, pos = { x = -1828.561, y = 117.770, z = 4850.472 }, rot = { x = 0.000, y = 28.153, z = 0.000 }, level = 32, point_type = 1001, area_id = 26 },
	{ config_id = 374004, gadget_id = 70500000, pos = { x = -1825.532, y = 119.044, z = 4850.559 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 32, point_type = 1001, area_id = 26 }
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
		gadgets = { 374001, 374002, 374003, 374004 },
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
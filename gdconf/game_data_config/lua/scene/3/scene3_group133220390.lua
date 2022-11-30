-- 基础信息
local base_info = {
	group_id = 133220390
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
	{ config_id = 390001, gadget_id = 70710078, pos = { x = -2789.356, y = 200.405, z = -4590.496 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 390002, gadget_id = 70710078, pos = { x = -2784.137, y = 200.466, z = -4590.652 }, rot = { x = 0.000, y = 190.617, z = 0.000 }, level = 1, area_id = 11 },
	{ config_id = 390003, gadget_id = 70710078, pos = { x = -2786.719, y = 200.288, z = -4596.886 }, rot = { x = 0.000, y = 3.001, z = 0.000 }, level = 1, area_id = 11 }
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
		gadgets = { 390001, 390002, 390003 },
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
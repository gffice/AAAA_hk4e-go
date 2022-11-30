-- 基础信息
local base_info = {
	group_id = 133103369
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 369001, monster_id = 26030101, pos = { x = -104.888, y = 322.065, z = 1680.126 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 27, drop_tag = "幼岩龙蜥", area_id = 6 },
	{ config_id = 369002, monster_id = 26030101, pos = { x = -94.033, y = 326.152, z = 1693.703 }, rot = { x = 0.000, y = 215.202, z = 0.000 }, level = 27, drop_tag = "幼岩龙蜥", disableWander = true, area_id = 6 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
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
		monsters = { 369001, 369002 },
		gadgets = { },
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
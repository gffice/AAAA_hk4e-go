-- 基础信息
local base_info = {
	group_id = 133008590
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 590001, monster_id = 21030501, pos = { x = 1017.197, y = 349.842, z = -643.150 }, rot = { x = 0.000, y = 205.348, z = 0.000 }, level = 30, drop_tag = "丘丘萨满", pose_id = 9012, climate_area_id = 1, area_id = 10 },
	{ config_id = 590002, monster_id = 21030501, pos = { x = 1015.390, y = 350.319, z = -650.567 }, rot = { x = 0.000, y = 312.259, z = 0.000 }, level = 30, drop_tag = "丘丘萨满", pose_id = 9012, climate_area_id = 1, area_id = 10 },
	{ config_id = 590003, monster_id = 21020601, pos = { x = 1017.525, y = 350.448, z = -647.696 }, rot = { x = 0.000, y = 276.558, z = 0.000 }, level = 30, drop_tag = "丘丘暴徒", pose_id = 401, climate_area_id = 1, area_id = 10 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 590004, gadget_id = 70300087, pos = { x = 1018.962, y = 350.465, z = -645.630 }, rot = { x = 6.133, y = 0.913, z = 16.917 }, level = 1, area_id = 10 },
	{ config_id = 590005, gadget_id = 70220060, pos = { x = 1015.805, y = 349.006, z = -641.954 }, rot = { x = 0.000, y = 282.710, z = 344.908 }, level = 1, area_id = 10 },
	{ config_id = 590006, gadget_id = 70220060, pos = { x = 1009.476, y = 349.571, z = -650.882 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, area_id = 10 },
	{ config_id = 590007, gadget_id = 70220060, pos = { x = 1009.088, y = 348.589, z = -650.132 }, rot = { x = 0.000, y = 315.082, z = 14.529 }, level = 1, area_id = 10 }
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
		monsters = { 590001, 590002, 590003 },
		gadgets = { 590004, 590005, 590006, 590007 },
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
-- 基础信息
local base_info = {
	group_id = 250004009
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
	{ config_id = 109, gadget_id = 70900208, pos = { x = 650.103, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 120, gadget_id = 70900228, pos = { x = 665.700, y = -11.208, z = -123.716 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1, route_id = 50004002 },
	{ config_id = 121, gadget_id = 70900230, pos = { x = 665.729, y = -11.060, z = -123.695 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 233, gadget_id = 70900208, pos = { x = 652.073, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 234, gadget_id = 70900208, pos = { x = 654.069, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 235, gadget_id = 70900208, pos = { x = 656.067, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 236, gadget_id = 70900207, pos = { x = 669.671, y = -11.770, z = -130.763 }, rot = { x = 0.000, y = 0.000, z = 0.000 }, level = 1 },
	{ config_id = 266, gadget_id = 70900208, pos = { x = 658.069, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 272, gadget_id = 70900208, pos = { x = 660.064, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 273, gadget_id = 70900208, pos = { x = 662.060, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 274, gadget_id = 70900208, pos = { x = 648.114, y = -10.877, z = -130.751 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 275, gadget_id = 70900208, pos = { x = 650.108, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 276, gadget_id = 70900208, pos = { x = 652.078, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 277, gadget_id = 70900208, pos = { x = 654.074, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 278, gadget_id = 70900208, pos = { x = 656.072, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 279, gadget_id = 70900208, pos = { x = 658.074, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 280, gadget_id = 70900208, pos = { x = 660.069, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 281, gadget_id = 70900208, pos = { x = 662.065, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 },
	{ config_id = 282, gadget_id = 70900208, pos = { x = 648.119, y = -10.866, z = -131.194 }, rot = { x = 0.000, y = 0.000, z = 90.000 }, level = 1 }
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
	rand_suite = true
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
		gadgets = { 109, 120, 121, 233, 234, 235, 236, 266, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282 },
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
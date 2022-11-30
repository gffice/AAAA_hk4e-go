-- 基础信息
local base_info = {
	group_id = 155005351
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
	{ config_id = 351001, gadget_id = 70211111, pos = { x = 505.597, y = 184.300, z = 642.495 }, rot = { x = 0.000, y = 79.870, z = 0.000 }, level = 16, drop_tag = "解谜中级稻妻", isOneoff = true, persistent = true, explore = { name = "chest", exp = 1 }, area_id = 200 }
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
		gadgets = { 351001 },
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
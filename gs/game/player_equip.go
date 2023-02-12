package game

import (
	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

// SetEquipLockStateReq 设置装备上锁状态请求
func (g *GameManager) SetEquipLockStateReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user set equip lock, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.SetEquipLockStateReq)

	// 获取目标装备
	equipGameObj, ok := player.GameObjectGuidMap[req.TargetEquipGuid]
	if !ok {
		logger.Error("equip error, equipGuid: %v", req.TargetEquipGuid)
		g.SendError(cmd.SetEquipLockStateRsp, player, &proto.SetEquipLockStateRsp{}, proto.Retcode_RET_ITEM_NOT_EXIST)
		return
	}
	switch equipGameObj.(type) {
	case *model.Weapon:
		weapon := equipGameObj.(*model.Weapon)
		weapon.Lock = req.IsLocked
		// 更新武器的物品数据
		g.SendMsg(cmd.StoreItemChangeNotify, player.PlayerID, player.ClientSeq, g.PacketStoreItemChangeNotifyByWeapon(weapon))
	case *model.Reliquary:
		reliquary := equipGameObj.(*model.Reliquary)
		reliquary.Lock = req.IsLocked
		// 更新圣遗物的物品数据
		g.SendMsg(cmd.StoreItemChangeNotify, player.PlayerID, player.ClientSeq, g.PacketStoreItemChangeNotifyByReliquary(reliquary))
	default:
		logger.Error("equip type error, equipGuid: %v", req.TargetEquipGuid)
		g.SendError(cmd.SetEquipLockStateRsp, player, &proto.SetEquipLockStateRsp{})
		return
	}

	setEquipLockStateRsp := &proto.SetEquipLockStateRsp{
		TargetEquipGuid: req.TargetEquipGuid,
		IsLocked:        req.IsLocked,
	}
	g.SendMsg(cmd.SetEquipLockStateRsp, player.PlayerID, player.ClientSeq, setEquipLockStateRsp)
}

// TakeoffEquipReq 装备卸下请求
func (g *GameManager) TakeoffEquipReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user take off equip, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.TakeoffEquipReq)

	// 获取目标角色
	avatar, ok := player.AvatarMap[player.GetAvatarIdByGuid(req.AvatarGuid)]
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", req.AvatarGuid)
		g.SendError(cmd.TakeoffEquipRsp, player, &proto.TakeoffEquipRsp{}, proto.Retcode_RET_CAN_NOT_FIND_AVATAR)
		return
	}
	// 确保角色已装备指定位置的圣遗物
	reliquary, ok := avatar.EquipReliquaryMap[uint8(req.Slot)]
	if !ok {
		logger.Error("avatar not wear reliquary, slot: %v", req.Slot)
		g.SendError(cmd.TakeoffEquipRsp, player, &proto.TakeoffEquipRsp{})
		return
	}
	// 卸下圣遗物
	player.TakeOffReliquary(avatar.AvatarId, reliquary.ReliquaryId)
	// 更新玩家装备
	avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByReliquary(avatar, reliquary)
	g.SendMsg(cmd.AvatarEquipChangeNotify, player.PlayerID, player.ClientSeq, avatarEquipChangeNotify)

	takeoffEquipRsp := &proto.TakeoffEquipRsp{
		AvatarGuid: req.AvatarGuid,
		Slot:       req.Slot,
	}
	g.SendMsg(cmd.TakeoffEquipRsp, player.PlayerID, player.ClientSeq, takeoffEquipRsp)
}

// WearEquipReq 穿戴装备请求
func (g *GameManager) WearEquipReq(player *model.Player, payloadMsg pb.Message) {
	logger.Debug("user wear equip, uid: %v", player.PlayerID)
	req := payloadMsg.(*proto.WearEquipReq)

	// 获取目标角色
	avatar, ok := player.AvatarMap[player.GetAvatarIdByGuid(req.AvatarGuid)]
	if !ok {
		logger.Error("avatar error, avatarGuid: %v", req.AvatarGuid)
		g.SendError(cmd.WearEquipRsp, player, &proto.WearEquipRsp{}, proto.Retcode_RET_CAN_NOT_FIND_AVATAR)
		return
	}
	// 获取目标装备
	equipGameObj, ok := player.GameObjectGuidMap[req.EquipGuid]
	if !ok {
		logger.Error("equip error, equipGuid: %v", req.EquipGuid)
		g.SendError(cmd.WearEquipRsp, player, &proto.WearEquipRsp{}, proto.Retcode_RET_ITEM_NOT_EXIST)
		return
	}
	switch equipGameObj.(type) {
	case *model.Weapon:
		weapon := equipGameObj.(*model.Weapon)
		g.WearUserAvatarWeapon(player.PlayerID, avatar.AvatarId, weapon.WeaponId)
	case *model.Reliquary:
		reliquary := equipGameObj.(*model.Reliquary)
		g.WearUserAvatarReliquary(player.PlayerID, avatar.AvatarId, reliquary.ReliquaryId)
	default:
		logger.Error("equip type error, equipGuid: %v", req.EquipGuid)
		g.SendError(cmd.WearEquipRsp, player, &proto.WearEquipRsp{})
		return
	}

	wearEquipRsp := &proto.WearEquipRsp{
		AvatarGuid: req.AvatarGuid,
		EquipGuid:  req.EquipGuid,
	}
	g.SendMsg(cmd.WearEquipRsp, player.PlayerID, player.ClientSeq, wearEquipRsp)
}

// WearUserAvatarReliquary 玩家角色装备圣遗物
func (g *GameManager) WearUserAvatarReliquary(userId uint32, avatarId uint32, reliquaryId uint64) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	avatar := player.AvatarMap[avatarId]
	reliquary := player.ReliquaryMap[reliquaryId]

	// 获取圣遗物配置表
	reliquaryConfig := gdconf.GetItemDataById(int32(reliquary.ItemId))
	if reliquaryConfig == nil {
		logger.Error("reliquary config error, itemId: %v", reliquary.ItemId)
		return
	}
	// 角色已装备的圣遗物
	avatarCurReliquary := avatar.EquipReliquaryMap[uint8(reliquaryConfig.ReliquaryType)]

	if reliquary.AvatarId != 0 {
		// 圣遗物在别的角色身上
		weakAvatarId := reliquary.AvatarId
		weakReliquaryId := reliquaryId
		strongAvatarId := avatarId
		strongReliquaryId := avatarCurReliquary.ReliquaryId
		player.TakeOffReliquary(weakAvatarId, weakReliquaryId)
		player.TakeOffReliquary(strongAvatarId, strongReliquaryId)
		player.WearReliquary(weakAvatarId, strongReliquaryId)
		player.WearReliquary(strongAvatarId, weakReliquaryId)

		weakAvatar := player.AvatarMap[weakAvatarId]
		weakReliquary := weakAvatar.EquipReliquaryMap[uint8(reliquaryConfig.ReliquaryType)]

		avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByReliquary(weakAvatar, weakReliquary)
		g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
	} else if avatarCurReliquary != nil {
		// 角色当前有圣遗物
		player.TakeOffReliquary(avatarId, avatarCurReliquary.ReliquaryId)
		player.WearReliquary(avatarId, reliquaryId)
	}

	avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByReliquary(avatar, reliquary)
	g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
}

// WearUserAvatarWeapon 玩家角色装备武器
func (g *GameManager) WearUserAvatarWeapon(userId uint32, avatarId uint32, weaponId uint64) {
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	avatar := player.AvatarMap[avatarId]
	weapon := player.WeaponMap[weaponId]

	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	if weapon.AvatarId != 0 {
		// 武器在别的角色身上
		weakAvatarId := weapon.AvatarId
		weakWeaponId := weaponId
		strongAvatarId := avatarId
		strongWeaponId := avatar.EquipWeapon.WeaponId
		player.TakeOffWeapon(weakAvatarId, weakWeaponId)
		player.TakeOffWeapon(strongAvatarId, strongWeaponId)
		player.WearWeapon(weakAvatarId, strongWeaponId)
		player.WearWeapon(strongAvatarId, weakWeaponId)

		weakAvatar := player.AvatarMap[weakAvatarId]
		weakWeapon := weakAvatar.EquipWeapon

		weakWorldAvatar := world.GetPlayerWorldAvatar(player, weakAvatarId)
		if weakWorldAvatar != nil {
			weakWorldAvatar.SetWeaponEntityId(scene.CreateEntityWeapon())
			avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByWeapon(weakAvatar, weakWeapon, weakWorldAvatar.GetWeaponEntityId())
			g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
		} else {
			avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByWeapon(weakAvatar, weakWeapon, 0)
			g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
		}
	} else if avatar.EquipWeapon != nil {
		// 角色当前有武器
		player.TakeOffWeapon(avatarId, avatar.EquipWeapon.WeaponId)
		player.WearWeapon(avatarId, weaponId)
	} else {
		// 是新角色还没有武器
		player.WearWeapon(avatarId, weaponId)
	}

	worldAvatar := world.GetPlayerWorldAvatar(player, avatarId)
	if worldAvatar != nil {
		worldAvatar.SetWeaponEntityId(scene.CreateEntityWeapon())
		avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByWeapon(avatar, weapon, worldAvatar.GetWeaponEntityId())
		g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
	} else {
		avatarEquipChangeNotify := g.PacketAvatarEquipChangeNotifyByWeapon(avatar, weapon, 0)
		g.SendMsg(cmd.AvatarEquipChangeNotify, userId, player.ClientSeq, avatarEquipChangeNotify)
	}
}

func (g *GameManager) PacketAvatarEquipChangeNotifyByReliquary(avatar *model.Avatar, reliquary *model.Reliquary) *proto.AvatarEquipChangeNotify {
	reliquaryConfig := gdconf.GetItemDataById(int32(reliquary.ItemId))
	if reliquaryConfig == nil {
		logger.Error("reliquary config error, itemId: %v", reliquary.ItemId)
		return new(proto.AvatarEquipChangeNotify)
	}
	avatarEquipChangeNotify := &proto.AvatarEquipChangeNotify{
		AvatarGuid: avatar.Guid,
		ItemId:     reliquary.ItemId,
		EquipGuid:  reliquary.Guid,
		EquipType:  uint32(reliquaryConfig.ReliquaryType),
		Reliquary: &proto.SceneReliquaryInfo{
			ItemId:       reliquary.ItemId,
			Guid:         reliquary.Guid,
			Level:        uint32(reliquary.Level),
			PromoteLevel: uint32(reliquary.Promote),
		},
	}
	return avatarEquipChangeNotify
}

func (g *GameManager) PacketAvatarEquipChangeNotifyByWeapon(avatar *model.Avatar, weapon *model.Weapon, entityId uint32) *proto.AvatarEquipChangeNotify {
	weaponConfig := gdconf.GetItemDataById(int32(weapon.ItemId))
	if weaponConfig == nil {
		logger.Error("weapon config error, itemId: %v", weapon.ItemId)
		return new(proto.AvatarEquipChangeNotify)
	}
	affixMap := make(map[uint32]uint32)
	for _, affixId := range weapon.AffixIdList {
		affixMap[affixId] = uint32(weapon.Refinement)
	}
	avatarEquipChangeNotify := &proto.AvatarEquipChangeNotify{
		AvatarGuid: avatar.Guid,
		ItemId:     weapon.ItemId,
		EquipGuid:  weapon.Guid,
		EquipType:  uint32(constant.EQUIP_TYPE_WEAPON),
		Weapon: &proto.SceneWeaponInfo{
			EntityId:     entityId,
			GadgetId:     uint32(weaponConfig.GadgetId),
			ItemId:       weapon.ItemId,
			Guid:         weapon.Guid,
			Level:        uint32(weapon.Level),
			PromoteLevel: uint32(weapon.Promote),
			AbilityInfo:  new(proto.AbilitySyncStateInfo),
			AffixMap:     affixMap,
		},
	}
	return avatarEquipChangeNotify
}

func (g *GameManager) PacketAvatarEquipTakeOffNotify(avatar *model.Avatar, weapon *model.Weapon) *proto.AvatarEquipChangeNotify {
	avatarEquipChangeNotify := &proto.AvatarEquipChangeNotify{
		AvatarGuid: avatar.Guid,
	}
	itemDataConfig := gdconf.GetItemDataById(int32(weapon.ItemId))
	if itemDataConfig != nil {
		avatarEquipChangeNotify.EquipType = uint32(itemDataConfig.Type)
	}
	return avatarEquipChangeNotify
}
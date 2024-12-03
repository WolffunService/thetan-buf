// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: thetan/immortal/v1/immortal_shared.proto

package thetan_immortal_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Arena int32

const (
	Arena_NONE        Arena = 0
	Arena_TRAINEE     Arena = 1
	Arena_APPERENTICE Arena = 2
	Arena_PATHFINDER  Arena = 3
	Arena_FIGHTER     Arena = 4
	Arena_ELITE       Arena = 5
	Arena_MASTER      Arena = 6
	Arena_HERO        Arena = 7
	Arena_CHAMPION    Arena = 8
	Arena_LEGEND      Arena = 9
	Arena_IMMORTAL    Arena = 10
	Arena_DEMIGOD     Arena = 11
)

// Enum value maps for Arena.
var (
	Arena_name = map[int32]string{
		0:  "NONE",
		1:  "TRAINEE",
		2:  "APPERENTICE",
		3:  "PATHFINDER",
		4:  "FIGHTER",
		5:  "ELITE",
		6:  "MASTER",
		7:  "HERO",
		8:  "CHAMPION",
		9:  "LEGEND",
		10: "IMMORTAL",
		11: "DEMIGOD",
	}
	Arena_value = map[string]int32{
		"NONE":        0,
		"TRAINEE":     1,
		"APPERENTICE": 2,
		"PATHFINDER":  3,
		"FIGHTER":     4,
		"ELITE":       5,
		"MASTER":      6,
		"HERO":        7,
		"CHAMPION":    8,
		"LEGEND":      9,
		"IMMORTAL":    10,
		"DEMIGOD":     11,
	}
)

func (x Arena) Enum() *Arena {
	p := new(Arena)
	*p = x
	return p
}

func (x Arena) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Arena) Descriptor() protoreflect.EnumDescriptor {
	return file_thetan_immortal_v1_immortal_shared_proto_enumTypes[0].Descriptor()
}

func (Arena) Type() protoreflect.EnumType {
	return &file_thetan_immortal_v1_immortal_shared_proto_enumTypes[0]
}

func (x Arena) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Arena.Descriptor instead.
func (Arena) EnumDescriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{0}
}

type Hero struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroID string `protobuf:"bytes,1,opt,name=heroID,proto3" json:"heroID,omitempty"`
	SkinID int32  `protobuf:"varint,2,opt,name=skinID,proto3" json:"skinID,omitempty"`
	Level  int32  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Type   int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Rarity int32  `protobuf:"varint,5,opt,name=rarity,proto3" json:"rarity,omitempty"`
}

func (x *Hero) Reset() {
	*x = Hero{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hero) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hero) ProtoMessage() {}

func (x *Hero) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hero.ProtoReflect.Descriptor instead.
func (*Hero) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Hero) GetHeroID() string {
	if x != nil {
		return x.HeroID
	}
	return ""
}

func (x *Hero) GetSkinID() int32 {
	if x != nil {
		return x.SkinID
	}
	return 0
}

func (x *Hero) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Hero) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Hero) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

type HeroFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rarity        int32  `protobuf:"varint,2,opt,name=rarity,proto3" json:"rarity,omitempty"`
	HeroID        int32  `protobuf:"varint,3,opt,name=heroID,proto3" json:"heroID,omitempty"`
	DefaultSkinID int32  `protobuf:"varint,4,opt,name=defaultSkinID,proto3" json:"defaultSkinID,omitempty"`
}

func (x *HeroFull) Reset() {
	*x = HeroFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroFull) ProtoMessage() {}

func (x *HeroFull) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroFull.ProtoReflect.Descriptor instead.
func (*HeroFull) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{1}
}

func (x *HeroFull) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeroFull) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

func (x *HeroFull) GetHeroID() int32 {
	if x != nil {
		return x.HeroID
	}
	return 0
}

func (x *HeroFull) GetDefaultSkinID() int32 {
	if x != nil {
		return x.DefaultSkinID
	}
	return 0
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillID uint32 `protobuf:"varint,1,opt,name=skillID,proto3" json:"skillID,omitempty"`
	Level   int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Skill) GetSkillID() uint32 {
	if x != nil {
		return x.SkillID
	}
	return 0
}

func (x *Skill) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type SkillFull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rarity  int32  `protobuf:"varint,2,opt,name=rarity,proto3" json:"rarity,omitempty"`
	SkillID int32  `protobuf:"varint,3,opt,name=skillID,proto3" json:"skillID,omitempty"`
	Class   int32  `protobuf:"varint,4,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *SkillFull) Reset() {
	*x = SkillFull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillFull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillFull) ProtoMessage() {}

func (x *SkillFull) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillFull.ProtoReflect.Descriptor instead.
func (*SkillFull) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{3}
}

func (x *SkillFull) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SkillFull) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

func (x *SkillFull) GetSkillID() int32 {
	if x != nil {
		return x.SkillID
	}
	return 0
}

func (x *SkillFull) GetClass() int32 {
	if x != nil {
		return x.Class
	}
	return 0
}

type SkillRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating     float64 `protobuf:"fixed64,1,opt,name=rating,proto3" json:"rating,omitempty"`
	Rd         float64 `protobuf:"fixed64,2,opt,name=rd,proto3" json:"rd,omitempty"`
	Volatility float64 `protobuf:"fixed64,3,opt,name=volatility,proto3" json:"volatility,omitempty"`
}

func (x *SkillRating) Reset() {
	*x = SkillRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillRating) ProtoMessage() {}

func (x *SkillRating) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillRating.ProtoReflect.Descriptor instead.
func (*SkillRating) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{4}
}

func (x *SkillRating) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *SkillRating) GetRd() float64 {
	if x != nil {
		return x.Rd
	}
	return 0
}

func (x *SkillRating) GetVolatility() float64 {
	if x != nil {
		return x.Volatility
	}
	return 0
}

type PlayerInfoMatchProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID    string `protobuf:"bytes,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	PlayerName  string `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`
	AvatarID    int32  `protobuf:"varint,3,opt,name=avatarID,proto3" json:"avatarID,omitempty"`
	FrameID     int32  `protobuf:"varint,4,opt,name=frameID,proto3" json:"frameID,omitempty"`
	NameColorID int32  `protobuf:"varint,5,opt,name=nameColorID,proto3" json:"nameColorID,omitempty"`
	BattleCount int32  `protobuf:"varint,6,opt,name=battleCount,proto3" json:"battleCount,omitempty"`
	PartyID     string `protobuf:"bytes,7,opt,name=partyID,proto3" json:"partyID,omitempty"`
	// Deprecated: Marked as deprecated in thetan/immortal/v1/immortal_shared.proto.
	TrophyRank int32 `protobuf:"varint,8,opt,name=trophyRank,proto3" json:"trophyRank,omitempty"`
	// Deprecated: Marked as deprecated in thetan/immortal/v1/immortal_shared.proto.
	TrophySearch int32 `protobuf:"varint,9,opt,name=trophySearch,proto3" json:"trophySearch,omitempty"`
	// Deprecated: Marked as deprecated in thetan/immortal/v1/immortal_shared.proto.
	Rank                int32    `protobuf:"varint,10,opt,name=rank,proto3" json:"rank,omitempty"`
	Hero                *Hero    `protobuf:"bytes,11,opt,name=hero,proto3" json:"hero,omitempty"`
	Skills              []*Skill `protobuf:"bytes,12,rep,name=skills,proto3" json:"skills,omitempty"`
	BotBrain            int32    `protobuf:"varint,13,opt,name=botBrain,proto3" json:"botBrain,omitempty"`
	Arena               Arena    `protobuf:"varint,14,opt,name=arena,proto3,enum=thetan.immortal.v1.Arena" json:"arena,omitempty"`
	WinRate             float64  `protobuf:"fixed64,15,opt,name=winRate,proto3" json:"winRate,omitempty"`
	Rating              float64  `protobuf:"fixed64,16,opt,name=rating,proto3" json:"rating,omitempty"`
	Rd                  float64  `protobuf:"fixed64,17,opt,name=rd,proto3" json:"rd,omitempty"`
	Volatility          float64  `protobuf:"fixed64,18,opt,name=volatility,proto3" json:"volatility,omitempty"`
	WinCount            int32    `protobuf:"varint,19,opt,name=winCount,proto3" json:"winCount,omitempty"`
	RecentMatchedHeroes []int32  `protobuf:"varint,20,rep,packed,name=recentMatchedHeroes,proto3" json:"recentMatchedHeroes,omitempty"`
	MaxHeroLevel        int32    `protobuf:"varint,21,opt,name=maxHeroLevel,proto3" json:"maxHeroLevel,omitempty"`
	MaxSkillLevel       int32    `protobuf:"varint,22,opt,name=maxSkillLevel,proto3" json:"maxSkillLevel,omitempty"`
	// premium user
	BattlePassID int32 `protobuf:"varint,23,opt,name=battlePassID,proto3" json:"battlePassID,omitempty"`
}

func (x *PlayerInfoMatchProto) Reset() {
	*x = PlayerInfoMatchProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInfoMatchProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfoMatchProto) ProtoMessage() {}

func (x *PlayerInfoMatchProto) ProtoReflect() protoreflect.Message {
	mi := &file_thetan_immortal_v1_immortal_shared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfoMatchProto.ProtoReflect.Descriptor instead.
func (*PlayerInfoMatchProto) Descriptor() ([]byte, []int) {
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerInfoMatchProto) GetPlayerID() string {
	if x != nil {
		return x.PlayerID
	}
	return ""
}

func (x *PlayerInfoMatchProto) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerInfoMatchProto) GetAvatarID() int32 {
	if x != nil {
		return x.AvatarID
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetFrameID() int32 {
	if x != nil {
		return x.FrameID
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetNameColorID() int32 {
	if x != nil {
		return x.NameColorID
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetBattleCount() int32 {
	if x != nil {
		return x.BattleCount
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetPartyID() string {
	if x != nil {
		return x.PartyID
	}
	return ""
}

// Deprecated: Marked as deprecated in thetan/immortal/v1/immortal_shared.proto.
func (x *PlayerInfoMatchProto) GetTrophyRank() int32 {
	if x != nil {
		return x.TrophyRank
	}
	return 0
}

// Deprecated: Marked as deprecated in thetan/immortal/v1/immortal_shared.proto.
func (x *PlayerInfoMatchProto) GetTrophySearch() int32 {
	if x != nil {
		return x.TrophySearch
	}
	return 0
}

// Deprecated: Marked as deprecated in thetan/immortal/v1/immortal_shared.proto.
func (x *PlayerInfoMatchProto) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetHero() *Hero {
	if x != nil {
		return x.Hero
	}
	return nil
}

func (x *PlayerInfoMatchProto) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *PlayerInfoMatchProto) GetBotBrain() int32 {
	if x != nil {
		return x.BotBrain
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetArena() Arena {
	if x != nil {
		return x.Arena
	}
	return Arena_NONE
}

func (x *PlayerInfoMatchProto) GetWinRate() float64 {
	if x != nil {
		return x.WinRate
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetRd() float64 {
	if x != nil {
		return x.Rd
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetVolatility() float64 {
	if x != nil {
		return x.Volatility
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetWinCount() int32 {
	if x != nil {
		return x.WinCount
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetRecentMatchedHeroes() []int32 {
	if x != nil {
		return x.RecentMatchedHeroes
	}
	return nil
}

func (x *PlayerInfoMatchProto) GetMaxHeroLevel() int32 {
	if x != nil {
		return x.MaxHeroLevel
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetMaxSkillLevel() int32 {
	if x != nil {
		return x.MaxSkillLevel
	}
	return 0
}

func (x *PlayerInfoMatchProto) GetBattlePassID() int32 {
	if x != nil {
		return x.BattlePassID
	}
	return 0
}

var File_thetan_immortal_v1_immortal_shared_proto protoreflect.FileDescriptor

var file_thetan_immortal_v1_immortal_shared_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x68, 0x65, 0x74,
	0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x78,
	0x0a, 0x04, 0x48, 0x65, 0x72, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x6b, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79, 0x22, 0x74, 0x0a, 0x08, 0x48, 0x65, 0x72, 0x6f,
	0x46, 0x75, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x68, 0x65, 0x72, 0x6f, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x6b, 0x69, 0x6e, 0x49, 0x44, 0x22, 0x37,
	0x0a, 0x05, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x67, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x46, 0x75, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x22, 0x55, 0x0a, 0x0b, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x02, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x61, 0x74,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x76, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x96, 0x06, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0a, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79,
	0x52, 0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x0c, 0x74, 0x72, 0x6f, 0x70, 0x68, 0x79, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0c,
	0x74, 0x72, 0x6f, 0x70, 0x68, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x2c, 0x0a, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x52, 0x04, 0x68, 0x65,
	0x72, 0x6f, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x42, 0x72, 0x61, 0x69,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x42, 0x72, 0x61, 0x69,
	0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x61, 0x72, 0x65, 0x6e, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x52, 0x05, 0x61, 0x72, 0x65,
	0x6e, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x77, 0x69, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x02, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x76, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x05, 0x52, 0x13, 0x72,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x48, 0x65, 0x72, 0x6f,
	0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x72, 0x6f, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x48, 0x65, 0x72,
	0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d,
	0x61, 0x78, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x49, 0x44,
	0x2a, 0xa2, 0x01, 0x0a, 0x05, 0x41, 0x72, 0x65, 0x6e, 0x61, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x45, 0x45, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x45, 0x52, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x45,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x54, 0x48, 0x46, 0x49, 0x4e, 0x44, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x49, 0x47, 0x48, 0x54, 0x45, 0x52, 0x10, 0x04, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x4c, 0x49, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41,
	0x53, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x52, 0x4f, 0x10, 0x07,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x41, 0x4d, 0x50, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x0a,
	0x0a, 0x06, 0x4c, 0x45, 0x47, 0x45, 0x4e, 0x44, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d,
	0x4d, 0x4f, 0x52, 0x54, 0x41, 0x4c, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4d, 0x49,
	0x47, 0x4f, 0x44, 0x10, 0x0b, 0x42, 0xd0, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x2e, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x42, 0x13, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x74, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2d,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x74, 0x61,
	0x6e, 0x2f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5f, 0x69, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x54, 0x49, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x2e,
	0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x54, 0x68,
	0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1e, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x5c, 0x49, 0x6d, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x14, 0x54, 0x68, 0x65, 0x74, 0x61, 0x6e, 0x3a, 0x3a, 0x49, 0x6d, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thetan_immortal_v1_immortal_shared_proto_rawDescOnce sync.Once
	file_thetan_immortal_v1_immortal_shared_proto_rawDescData = file_thetan_immortal_v1_immortal_shared_proto_rawDesc
)

func file_thetan_immortal_v1_immortal_shared_proto_rawDescGZIP() []byte {
	file_thetan_immortal_v1_immortal_shared_proto_rawDescOnce.Do(func() {
		file_thetan_immortal_v1_immortal_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_thetan_immortal_v1_immortal_shared_proto_rawDescData)
	})
	return file_thetan_immortal_v1_immortal_shared_proto_rawDescData
}

var file_thetan_immortal_v1_immortal_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thetan_immortal_v1_immortal_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_thetan_immortal_v1_immortal_shared_proto_goTypes = []interface{}{
	(Arena)(0),                   // 0: thetan.immortal.v1.Arena
	(*Hero)(nil),                 // 1: thetan.immortal.v1.Hero
	(*HeroFull)(nil),             // 2: thetan.immortal.v1.HeroFull
	(*Skill)(nil),                // 3: thetan.immortal.v1.Skill
	(*SkillFull)(nil),            // 4: thetan.immortal.v1.SkillFull
	(*SkillRating)(nil),          // 5: thetan.immortal.v1.SkillRating
	(*PlayerInfoMatchProto)(nil), // 6: thetan.immortal.v1.PlayerInfoMatchProto
}
var file_thetan_immortal_v1_immortal_shared_proto_depIdxs = []int32{
	1, // 0: thetan.immortal.v1.PlayerInfoMatchProto.hero:type_name -> thetan.immortal.v1.Hero
	3, // 1: thetan.immortal.v1.PlayerInfoMatchProto.skills:type_name -> thetan.immortal.v1.Skill
	0, // 2: thetan.immortal.v1.PlayerInfoMatchProto.arena:type_name -> thetan.immortal.v1.Arena
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_thetan_immortal_v1_immortal_shared_proto_init() }
func file_thetan_immortal_v1_immortal_shared_proto_init() {
	if File_thetan_immortal_v1_immortal_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thetan_immortal_v1_immortal_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hero); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thetan_immortal_v1_immortal_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeroFull); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thetan_immortal_v1_immortal_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thetan_immortal_v1_immortal_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillFull); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thetan_immortal_v1_immortal_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillRating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thetan_immortal_v1_immortal_shared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerInfoMatchProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thetan_immortal_v1_immortal_shared_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_thetan_immortal_v1_immortal_shared_proto_goTypes,
		DependencyIndexes: file_thetan_immortal_v1_immortal_shared_proto_depIdxs,
		EnumInfos:         file_thetan_immortal_v1_immortal_shared_proto_enumTypes,
		MessageInfos:      file_thetan_immortal_v1_immortal_shared_proto_msgTypes,
	}.Build()
	File_thetan_immortal_v1_immortal_shared_proto = out.File
	file_thetan_immortal_v1_immortal_shared_proto_rawDesc = nil
	file_thetan_immortal_v1_immortal_shared_proto_goTypes = nil
	file_thetan_immortal_v1_immortal_shared_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: idl/quic/game.proto

package quic

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

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Action_TankMove
	//	*Action_GetBulletList
	//	*Action_TankRemove
	//	*Action_BulletRemove
	//	*Action_GetTankList
	//	*Action_GetRockList
	//	*Action_NewBullet
	//	*Action_GetExplosionList
	//	*Action_GameOver
	Type isAction_Type `protobuf_oneof:"type"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_quic_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_idl_quic_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_idl_quic_game_proto_rawDescGZIP(), []int{0}
}

func (m *Action) GetType() isAction_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Action) GetTankMove() *TankMove {
	if x, ok := x.GetType().(*Action_TankMove); ok {
		return x.TankMove
	}
	return nil
}

func (x *Action) GetGetBulletList() *GetBulletList {
	if x, ok := x.GetType().(*Action_GetBulletList); ok {
		return x.GetBulletList
	}
	return nil
}

func (x *Action) GetTankRemove() *TankRemove {
	if x, ok := x.GetType().(*Action_TankRemove); ok {
		return x.TankRemove
	}
	return nil
}

func (x *Action) GetBulletRemove() *BulletRemove {
	if x, ok := x.GetType().(*Action_BulletRemove); ok {
		return x.BulletRemove
	}
	return nil
}

func (x *Action) GetGetTankList() *GetTankList {
	if x, ok := x.GetType().(*Action_GetTankList); ok {
		return x.GetTankList
	}
	return nil
}

func (x *Action) GetGetRockList() *GetRockList {
	if x, ok := x.GetType().(*Action_GetRockList); ok {
		return x.GetRockList
	}
	return nil
}

func (x *Action) GetNewBullet() *NewBullet {
	if x, ok := x.GetType().(*Action_NewBullet); ok {
		return x.NewBullet
	}
	return nil
}

func (x *Action) GetGetExplosionList() *GetExplosionList {
	if x, ok := x.GetType().(*Action_GetExplosionList); ok {
		return x.GetExplosionList
	}
	return nil
}

func (x *Action) GetGameOver() *GameOver {
	if x, ok := x.GetType().(*Action_GameOver); ok {
		return x.GameOver
	}
	return nil
}

type isAction_Type interface {
	isAction_Type()
}

type Action_TankMove struct {
	TankMove *TankMove `protobuf:"bytes,1,opt,name=tankMove,proto3,oneof"`
}

type Action_GetBulletList struct {
	GetBulletList *GetBulletList `protobuf:"bytes,4,opt,name=getBulletList,proto3,oneof"`
}

type Action_TankRemove struct {
	TankRemove *TankRemove `protobuf:"bytes,5,opt,name=tankRemove,proto3,oneof"`
}

type Action_BulletRemove struct {
	BulletRemove *BulletRemove `protobuf:"bytes,6,opt,name=bulletRemove,proto3,oneof"`
}

type Action_GetTankList struct {
	GetTankList *GetTankList `protobuf:"bytes,7,opt,name=getTankList,proto3,oneof"`
}

type Action_GetRockList struct {
	GetRockList *GetRockList `protobuf:"bytes,8,opt,name=getRockList,proto3,oneof"`
}

type Action_NewBullet struct {
	NewBullet *NewBullet `protobuf:"bytes,9,opt,name=newBullet,proto3,oneof"`
}

type Action_GetExplosionList struct {
	GetExplosionList *GetExplosionList `protobuf:"bytes,10,opt,name=getExplosionList,proto3,oneof"`
}

type Action_GameOver struct {
	GameOver *GameOver `protobuf:"bytes,11,opt,name=gameOver,proto3,oneof"`
}

func (*Action_TankMove) isAction_Type() {}

func (*Action_GetBulletList) isAction_Type() {}

func (*Action_TankRemove) isAction_Type() {}

func (*Action_BulletRemove) isAction_Type() {}

func (*Action_GetTankList) isAction_Type() {}

func (*Action_GetRockList) isAction_Type() {}

func (*Action_NewBullet) isAction_Type() {}

func (*Action_GetExplosionList) isAction_Type() {}

func (*Action_GameOver) isAction_Type() {}

type ActionList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *ActionList) Reset() {
	*x = ActionList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_quic_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionList) ProtoMessage() {}

func (x *ActionList) ProtoReflect() protoreflect.Message {
	mi := &file_idl_quic_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionList.ProtoReflect.Descriptor instead.
func (*ActionList) Descriptor() ([]byte, []int) {
	return file_idl_quic_game_proto_rawDescGZIP(), []int{1}
}

func (x *ActionList) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

type GameOver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameOver bool `protobuf:"varint,1,opt,name=gameOver,proto3" json:"gameOver,omitempty"`
}

func (x *GameOver) Reset() {
	*x = GameOver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_quic_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameOver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameOver) ProtoMessage() {}

func (x *GameOver) ProtoReflect() protoreflect.Message {
	mi := &file_idl_quic_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameOver.ProtoReflect.Descriptor instead.
func (*GameOver) Descriptor() ([]byte, []int) {
	return file_idl_quic_game_proto_rawDescGZIP(), []int{2}
}

func (x *GameOver) GetGameOver() bool {
	if x != nil {
		return x.GameOver
	}
	return false
}

var File_idl_quic_game_proto protoreflect.FileDescriptor

var file_idl_quic_game_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x64, 0x6c, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x13, 0x69, 0x64, 0x6c,
	0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x74, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x69, 0x64, 0x6c, 0x2f, 0x71, 0x75, 0x69, 0x63, 0x2f, 0x62, 0x75, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x69, 0x64, 0x6c, 0x2f, 0x71, 0x75, 0x69,
	0x63, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x03,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x61, 0x6e, 0x6b,
	0x4d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x54, 0x61, 0x6e, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x08, 0x74, 0x61,
	0x6e, 0x6b, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x42, 0x75, 0x6c,
	0x6c, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x67, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x54,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x62, 0x75, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x12, 0x35, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x54, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x65, 0x74,
	0x54, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x52,
	0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x42, 0x75, 0x6c,
	0x6c, 0x65, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x44, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x73, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x67, 0x65, 0x74, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4f, 0x76,
	0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65,
	0x4f, 0x76, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x71,
	0x75, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_quic_game_proto_rawDescOnce sync.Once
	file_idl_quic_game_proto_rawDescData = file_idl_quic_game_proto_rawDesc
)

func file_idl_quic_game_proto_rawDescGZIP() []byte {
	file_idl_quic_game_proto_rawDescOnce.Do(func() {
		file_idl_quic_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_quic_game_proto_rawDescData)
	})
	return file_idl_quic_game_proto_rawDescData
}

var file_idl_quic_game_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_idl_quic_game_proto_goTypes = []interface{}{
	(*Action)(nil),           // 0: game.Action
	(*ActionList)(nil),       // 1: game.ActionList
	(*GameOver)(nil),         // 2: game.GameOver
	(*TankMove)(nil),         // 3: game.TankMove
	(*GetBulletList)(nil),    // 4: game.GetBulletList
	(*TankRemove)(nil),       // 5: game.TankRemove
	(*BulletRemove)(nil),     // 6: game.BulletRemove
	(*GetTankList)(nil),      // 7: game.GetTankList
	(*GetRockList)(nil),      // 8: game.GetRockList
	(*NewBullet)(nil),        // 9: game.NewBullet
	(*GetExplosionList)(nil), // 10: game.GetExplosionList
}
var file_idl_quic_game_proto_depIdxs = []int32{
	3,  // 0: game.Action.tankMove:type_name -> game.TankMove
	4,  // 1: game.Action.getBulletList:type_name -> game.GetBulletList
	5,  // 2: game.Action.tankRemove:type_name -> game.TankRemove
	6,  // 3: game.Action.bulletRemove:type_name -> game.BulletRemove
	7,  // 4: game.Action.getTankList:type_name -> game.GetTankList
	8,  // 5: game.Action.getRockList:type_name -> game.GetRockList
	9,  // 6: game.Action.newBullet:type_name -> game.NewBullet
	10, // 7: game.Action.getExplosionList:type_name -> game.GetExplosionList
	2,  // 8: game.Action.gameOver:type_name -> game.GameOver
	0,  // 9: game.ActionList.actions:type_name -> game.Action
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_idl_quic_game_proto_init() }
func file_idl_quic_game_proto_init() {
	if File_idl_quic_game_proto != nil {
		return
	}
	file_idl_quic_tank_proto_init()
	file_idl_quic_bullet_proto_init()
	file_idl_quic_world_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_idl_quic_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_idl_quic_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionList); i {
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
		file_idl_quic_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameOver); i {
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
	file_idl_quic_game_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Action_TankMove)(nil),
		(*Action_GetBulletList)(nil),
		(*Action_TankRemove)(nil),
		(*Action_BulletRemove)(nil),
		(*Action_GetTankList)(nil),
		(*Action_GetRockList)(nil),
		(*Action_NewBullet)(nil),
		(*Action_GetExplosionList)(nil),
		(*Action_GameOver)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_quic_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_idl_quic_game_proto_goTypes,
		DependencyIndexes: file_idl_quic_game_proto_depIdxs,
		MessageInfos:      file_idl_quic_game_proto_msgTypes,
	}.Build()
	File_idl_quic_game_proto = out.File
	file_idl_quic_game_proto_rawDesc = nil
	file_idl_quic_game_proto_goTypes = nil
	file_idl_quic_game_proto_depIdxs = nil
}

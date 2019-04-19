// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order struct {
	OrderId              string             `protobuf:"bytes,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	Status               string             `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	CreatedOn            int64              `protobuf:"varint,3,opt,name=created_on,json=createdOn" json:"created_on,omitempty"`
	OrderItems           []*Order_OrderItem `protobuf:"bytes,4,rep,name=order_items,json=orderItems" json:"order_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_6a247e85c4352899, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (dst *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(dst, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *Order) GetOrderItems() []*Order_OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

type Order_OrderItem struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	UnitPrice            float32  `protobuf:"fixed32,3,opt,name=unit_price,json=unitPrice" json:"unit_price,omitempty"`
	Quantity             int32    `protobuf:"varint,4,opt,name=quantity" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order_OrderItem) Reset()         { *m = Order_OrderItem{} }
func (m *Order_OrderItem) String() string { return proto.CompactTextString(m) }
func (*Order_OrderItem) ProtoMessage()    {}
func (*Order_OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_6a247e85c4352899, []int{0, 0}
}
func (m *Order_OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order_OrderItem.Unmarshal(m, b)
}
func (m *Order_OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order_OrderItem.Marshal(b, m, deterministic)
}
func (dst *Order_OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order_OrderItem.Merge(dst, src)
}
func (m *Order_OrderItem) XXX_Size() int {
	return xxx_messageInfo_Order_OrderItem.Size(m)
}
func (m *Order_OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Order_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_Order_OrderItem proto.InternalMessageInfo

func (m *Order_OrderItem) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Order_OrderItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order_OrderItem) GetUnitPrice() float32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Order_OrderItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type OrderResponse struct {
	IsSuccess            bool     `protobuf:"varint,1,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_6a247e85c4352899, []int{1}
}
func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResponse.Unmarshal(m, b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
}
func (dst *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(dst, src)
}
func (m *OrderResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResponse.Size(m)
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

func (m *OrderResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *OrderResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type OrderFilter struct {
	SearchText           string   `protobuf:"bytes,1,opt,name=search_text,json=searchText" json:"search_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderFilter) Reset()         { *m = OrderFilter{} }
func (m *OrderFilter) String() string { return proto.CompactTextString(m) }
func (*OrderFilter) ProtoMessage()    {}
func (*OrderFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_6a247e85c4352899, []int{2}
}
func (m *OrderFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFilter.Unmarshal(m, b)
}
func (m *OrderFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFilter.Marshal(b, m, deterministic)
}
func (dst *OrderFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFilter.Merge(dst, src)
}
func (m *OrderFilter) XXX_Size() int {
	return xxx_messageInfo_OrderFilter.Size(m)
}
func (m *OrderFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFilter.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFilter proto.InternalMessageInfo

func (m *OrderFilter) GetSearchText() string {
	if m != nil {
		return m.SearchText
	}
	return ""
}

type ServiceDiscovery struct {
	OrderServiceUri      string   `protobuf:"bytes,1,opt,name=order_service_uri,json=orderServiceUri" json:"order_service_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceDiscovery) Reset()         { *m = ServiceDiscovery{} }
func (m *ServiceDiscovery) String() string { return proto.CompactTextString(m) }
func (*ServiceDiscovery) ProtoMessage()    {}
func (*ServiceDiscovery) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_6a247e85c4352899, []int{3}
}
func (m *ServiceDiscovery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceDiscovery.Unmarshal(m, b)
}
func (m *ServiceDiscovery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceDiscovery.Marshal(b, m, deterministic)
}
func (dst *ServiceDiscovery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceDiscovery.Merge(dst, src)
}
func (m *ServiceDiscovery) XXX_Size() int {
	return xxx_messageInfo_ServiceDiscovery.Size(m)
}
func (m *ServiceDiscovery) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceDiscovery.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceDiscovery proto.InternalMessageInfo

func (m *ServiceDiscovery) GetOrderServiceUri() string {
	if m != nil {
		return m.OrderServiceUri
	}
	return ""
}

type EventStore struct {
	AggregateId          string   `protobuf:"bytes,1,opt,name=aggregate_id,json=aggregateId" json:"aggregate_id,omitempty"`
	AggregateType        string   `protobuf:"bytes,2,opt,name=aggregate_type,json=aggregateType" json:"aggregate_type,omitempty"`
	EventId              string   `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	EventType            string   `protobuf:"bytes,4,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	EventData            string   `protobuf:"bytes,5,opt,name=event_data,json=eventData" json:"event_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventStore) Reset()         { *m = EventStore{} }
func (m *EventStore) String() string { return proto.CompactTextString(m) }
func (*EventStore) ProtoMessage()    {}
func (*EventStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_6a247e85c4352899, []int{4}
}
func (m *EventStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStore.Unmarshal(m, b)
}
func (m *EventStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStore.Marshal(b, m, deterministic)
}
func (dst *EventStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStore.Merge(dst, src)
}
func (m *EventStore) XXX_Size() int {
	return xxx_messageInfo_EventStore.Size(m)
}
func (m *EventStore) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStore.DiscardUnknown(m)
}

var xxx_messageInfo_EventStore proto.InternalMessageInfo

func (m *EventStore) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *EventStore) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *EventStore) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventStore) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventStore) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "order.Order")
	proto.RegisterType((*Order_OrderItem)(nil), "order.Order.OrderItem")
	proto.RegisterType((*OrderResponse)(nil), "order.OrderResponse")
	proto.RegisterType((*OrderFilter)(nil), "order.OrderFilter")
	proto.RegisterType((*ServiceDiscovery)(nil), "order.ServiceDiscovery")
	proto.RegisterType((*EventStore)(nil), "order.EventStore")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	// Get all Orders with filter - A server-to-client streaming RPC.
	GetOrders(ctx context.Context, in *OrderFilter, opts ...grpc.CallOption) (OrderService_GetOrdersClient, error)
	// Create a new Order - A simple RPC
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *OrderFilter, opts ...grpc.CallOption) (OrderService_GetOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderService_serviceDesc.Streams[0], "/order.OrderService/GetOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_GetOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderServiceGetOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	// Get all Orders with filter - A server-to-client streaming RPC.
	GetOrders(*OrderFilter, OrderService_GetOrdersServer) error
	// Create a new Order - A simple RPC
	CreateOrder(context.Context, *Order) (*OrderResponse, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_GetOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrderFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).GetOrders(m, &orderServiceGetOrdersServer{stream})
}

type OrderService_GetOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderServiceGetOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrders",
			Handler:       _OrderService_GetOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order.proto",
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_order_6a247e85c4352899) }

var fileDescriptor_order_6a247e85c4352899 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xc7, 0xeb, 0x3c, 0x4a, 0x3c, 0x4e, 0x79, 0xac, 0xaa, 0xca, 0x44, 0x42, 0x04, 0x4b, 0x48,
	0x11, 0x87, 0x08, 0xda, 0x03, 0x37, 0x2e, 0x04, 0x50, 0x4e, 0x41, 0x4e, 0x39, 0x5b, 0x8b, 0x3d,
	0x0a, 0x2b, 0x11, 0xaf, 0xd9, 0x1d, 0x47, 0xcd, 0x95, 0x2f, 0xc4, 0x57, 0x44, 0x3b, 0xbb, 0x71,
	0xdc, 0x4b, 0x34, 0xff, 0xdf, 0xdf, 0xf3, 0xc8, 0xec, 0x40, 0xa2, 0x4d, 0x85, 0x66, 0xd9, 0x18,
	0x4d, 0x5a, 0x8c, 0x59, 0x64, 0x7f, 0x07, 0x30, 0xde, 0xb8, 0x48, 0xbc, 0x84, 0x09, 0xa3, 0x42,
	0x55, 0x69, 0x34, 0x8f, 0x16, 0x71, 0xfe, 0x84, 0xf5, 0xba, 0x12, 0x37, 0x70, 0x69, 0x49, 0x52,
	0x6b, 0xd3, 0x01, 0x1b, 0x41, 0x89, 0x57, 0x00, 0xa5, 0x41, 0x49, 0x58, 0x15, 0xba, 0x4e, 0x87,
	0xf3, 0x68, 0x31, 0xcc, 0xe3, 0x40, 0x36, 0xb5, 0xf8, 0x18, 0x3a, 0x16, 0x8a, 0x70, 0x6f, 0xd3,
	0xd1, 0x7c, 0xb8, 0x48, 0x6e, 0x6f, 0x96, 0x7e, 0x8a, 0xcd, 0xf9, 0x77, 0x4d, 0xb8, 0xcf, 0x41,
	0x9f, 0x42, 0x3b, 0xab, 0x21, 0xee, 0x0c, 0x21, 0x60, 0x54, 0xea, 0x0a, 0xc3, 0x4c, 0x1c, 0x3b,
	0x56, 0xcb, 0x3d, 0x86, 0x71, 0x38, 0x76, 0xc3, 0xb4, 0xb5, 0xa2, 0xa2, 0x31, 0xaa, 0x44, 0x1e,
	0x66, 0x90, 0xc7, 0x8e, 0x7c, 0x77, 0x40, 0xcc, 0x60, 0xf2, 0xa7, 0x95, 0x35, 0x29, 0x3a, 0xa6,
	0xa3, 0x79, 0xb4, 0x18, 0xe7, 0x9d, 0xce, 0x56, 0x70, 0xc5, 0xfd, 0x72, 0xb4, 0x8d, 0xae, 0x2d,
	0xd7, 0x52, 0xb6, 0xb0, 0x6d, 0x59, 0xa2, 0xb5, 0xdc, 0x79, 0x92, 0xc7, 0xca, 0x6e, 0x3d, 0x10,
	0xd7, 0x30, 0x46, 0x63, 0xb4, 0x09, 0xfd, 0xbd, 0xc8, 0x96, 0x90, 0x70, 0x95, 0xaf, 0xea, 0x37,
	0xa1, 0x11, 0xaf, 0x21, 0xb1, 0x28, 0x4d, 0xf9, 0xab, 0x20, 0x7c, 0xa0, 0x30, 0x3e, 0x78, 0x74,
	0x8f, 0x0f, 0x94, 0x7d, 0x82, 0xe7, 0x5b, 0x34, 0x07, 0x55, 0xe2, 0x4a, 0xd9, 0x52, 0x1f, 0xd0,
	0x1c, 0xc5, 0x3b, 0x78, 0xe1, 0x57, 0x66, 0xbd, 0x53, 0xb4, 0x46, 0x85, 0xd4, 0x67, 0x6c, 0x84,
	0x8c, 0x1f, 0x46, 0x65, 0xff, 0x22, 0x80, 0x2f, 0x07, 0xac, 0x69, 0x4b, 0xda, 0xa0, 0x78, 0x03,
	0x53, 0xb9, 0xdb, 0x19, 0xdc, 0x49, 0xc2, 0xf3, 0x1b, 0x26, 0x1d, 0x5b, 0x57, 0xe2, 0x2d, 0x3c,
	0x3d, 0x7f, 0x42, 0xc7, 0xe6, 0xb4, 0xc0, 0xab, 0x8e, 0xde, 0x1f, 0x1b, 0x74, 0x97, 0x80, 0xae,
	0xae, 0xab, 0x32, 0xf4, 0x97, 0xc0, 0x7a, 0x5d, 0xb9, 0xc5, 0x78, 0x8b, 0xb3, 0x47, 0x6c, 0xc6,
	0x4c, 0x38, 0xb3, 0xb3, 0x2b, 0x49, 0x32, 0x1d, 0xf7, 0xec, 0x95, 0x24, 0x79, 0x7b, 0x80, 0xe9,
	0xa6, 0xf7, 0x27, 0xc4, 0x07, 0x88, 0xbf, 0x21, 0x31, 0xb2, 0x42, 0xf4, 0x0f, 0xc3, 0xef, 0x70,
	0x36, 0xed, 0xb3, 0xec, 0xe2, 0x7d, 0x24, 0xee, 0x20, 0xf9, 0xcc, 0x07, 0xe6, 0x8f, 0xf6, 0xd1,
	0x07, 0xb3, 0xeb, 0xbe, 0x3a, 0x3d, 0x66, 0x76, 0xf1, 0xf3, 0x92, 0x4f, 0xfe, 0xee, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0x49, 0x00, 0xdc, 0x01, 0x03, 0x00, 0x00,
}

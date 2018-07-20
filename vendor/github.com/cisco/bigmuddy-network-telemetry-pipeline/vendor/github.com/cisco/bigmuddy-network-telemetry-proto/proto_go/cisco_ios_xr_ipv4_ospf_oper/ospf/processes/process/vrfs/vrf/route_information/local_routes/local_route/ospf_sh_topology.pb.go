// Code generated by protoc-gen-go.
// source: ospf_sh_topology.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_local_routes_local_route is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_topology.proto

It has these top-level messages:
	OspfShTopology_KEYS
	OspfShTopology
	OspfShTime
	OspfShTopCommon
	OspfShTopPath
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_local_routes_local_route

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OSPF Route Information
type OspfShTopology_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	Prefix       string `protobuf:"bytes,3,opt,name=prefix" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *OspfShTopology_KEYS) Reset()                    { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()               {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShTopology_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type OspfShTopology struct {
	// Prefix
	RoutePrefix string `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix" json:"route_prefix,omitempty"`
	// Prefix length
	RoutePrefixLength uint32 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength" json:"route_prefix_length,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// Route type
	RouteType string `protobuf:"bytes,53,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// If true, connected route
	RouteConnected bool `protobuf:"varint,54,opt,name=route_connected,json=routeConnected" json:"route_connected,omitempty"`
	// Route information
	RouteInfo *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo" json:"route_info,omitempty"`
	// List of paths to this route
	RoutePathList []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path_list,json=routePathList" json:"route_path_list,omitempty"`
}

func (m *OspfShTopology) Reset()                    { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()               {}
func (*OspfShTopology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShTopology) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopology) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopology) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopology) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopology) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopology) GetRoutePathList() []*OspfShTopPath {
	if m != nil {
		return m.RoutePathList
	}
	return nil
}

type OspfShTime struct {
	Second     uint32 `protobuf:"varint,1,opt,name=second" json:"second,omitempty"`
	Nanosecond uint32 `protobuf:"varint,2,opt,name=nanosecond" json:"nanosecond,omitempty"`
}

func (m *OspfShTime) Reset()                    { *m = OspfShTime{} }
func (m *OspfShTime) String() string            { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()               {}
func (*OspfShTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

// OSPF Common Route Information
type OspfShTopCommon struct {
	// Area ID
	RouteAreaId uint32 `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId" json:"route_area_id,omitempty"`
	// TE metric
	RouteTeMetric uint32 `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric" json:"route_te_metric,omitempty"`
	// RIB version
	RouteRibVersion uint32 `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion" json:"route_rib_version,omitempty"`
	// SPF version
	RouteSpfVersion uint64 `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion" json:"route_spf_version,omitempty"`
	// Forward distance
	RouteForwardDistance uint32 `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance" json:"route_forward_distance,omitempty"`
	// Protocol source
	RouteSource uint32 `protobuf:"varint,6,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// Last time updated
	RouteUpdateTime *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime" json:"route_update_time,omitempty"`
	// Last time update failed
	RouteFailTime *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime" json:"route_fail_time,omitempty"`
	// SPF priority
	RouteSpfPriority uint32 `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority" json:"route_spf_priority,omitempty"`
	// If true, exclude from TE paths
	RouteAutoExcluded bool `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded" json:"route_auto_excluded,omitempty"`
	// If true, SRTE registered prefix route
	RouteSrtePrefixRegistered bool `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered" json:"route_srte_prefix_registered,omitempty"`
	// SRTE registered neigbhor count on route
	RouteSrteNbrRegistered uint32 `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered" json:"route_srte_nbr_registered,omitempty"`
}

func (m *OspfShTopCommon) Reset()                    { *m = OspfShTopCommon{} }
func (m *OspfShTopCommon) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopCommon) ProtoMessage()               {}
func (*OspfShTopCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShTopCommon) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if m != nil {
		return m.RouteTeMetric
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if m != nil {
		return m.RouteRibVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if m != nil {
		return m.RouteSpfVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if m != nil {
		return m.RouteForwardDistance
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSource() uint32 {
	if m != nil {
		return m.RouteSource
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if m != nil {
		return m.RouteUpdateTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if m != nil {
		return m.RouteFailTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if m != nil {
		return m.RouteSpfPriority
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if m != nil {
		return m.RouteAutoExcluded
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if m != nil {
		return m.RouteSrtePrefixRegistered
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if m != nil {
		return m.RouteSrteNbrRegistered
	}
	return 0
}

// OSPF Route Path Information
type OspfShTopPath struct {
	// Next hop Interface
	RouteInterfaceName string `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName" json:"route_interface_name,omitempty"`
	// Nexthop IP address
	RouteNextHopAddress string `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress" json:"route_next_hop_address,omitempty"`
	// IP address of source of route
	RouteSource string `protobuf:"bytes,3,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// LSA ID, see RFC2328
	RouteLsaid string `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid" json:"route_lsaid,omitempty"`
	// Multicast-intact path
	RoutePathIsMcastIntact bool `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact" json:"route_path_is_mcast_intact,omitempty"`
	// UCMP path
	RoutePathIsUcmpPath bool `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath" json:"route_path_is_ucmp_path,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,7,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// LSA type, see RFC2328 etc.
	LsaType uint32 `protobuf:"varint,8,opt,name=lsa_type,json=lsaType" json:"lsa_type,omitempty"`
	// Area ID
	AreaId uint32 `protobuf:"varint,9,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	// Area format IP or uint32
	AreaFormat bool `protobuf:"varint,10,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
}

func (m *OspfShTopPath) Reset()                    { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()               {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OspfShTopPath) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPath) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPath) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPath) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPath) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPath) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPath) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopPath) GetAreaFormat() bool {
	if m != nil {
		return m.AreaFormat
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.local_routes.local_route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.local_routes.local_route.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.local_routes.local_route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.local_routes.local_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.local_routes.local_route.ospf_sh_top_path")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x6f, 0x23, 0x35,
	0x14, 0xd6, 0x6c, 0x4b, 0x92, 0x3a, 0x0d, 0x2c, 0x66, 0xe9, 0x4e, 0x11, 0xb0, 0x21, 0x48, 0x10,
	0x21, 0x34, 0x42, 0x6d, 0xf9, 0x79, 0x41, 0x15, 0x6c, 0x44, 0x44, 0xb7, 0x5a, 0x4d, 0xbb, 0x48,
	0x9c, 0x2c, 0xc7, 0xe3, 0x69, 0x2c, 0xcd, 0x8c, 0x47, 0xb6, 0x13, 0x12, 0xfe, 0x09, 0x24, 0x6e,
	0x48, 0x1c, 0xb8, 0x71, 0xe2, 0x7f, 0x44, 0x7e, 0xcf, 0x93, 0x99, 0xdd, 0xdc, 0xb7, 0x97, 0x68,
	0xfc, 0x7d, 0xcf, 0xcf, 0x9f, 0xfd, 0x3e, 0x3f, 0x87, 0x9c, 0x68, 0x5b, 0xe7, 0xcc, 0x2e, 0x99,
	0xd3, 0xb5, 0x2e, 0xf4, 0xdd, 0x36, 0xa9, 0x8d, 0x76, 0x9a, 0xe6, 0x42, 0x59, 0xa1, 0x99, 0xd2,
	0x96, 0x6d, 0x0c, 0x53, 0xf5, 0xfa, 0x82, 0x41, 0xa4, 0xae, 0xa5, 0x49, 0xfc, 0x97, 0x8f, 0x13,
	0xd2, 0x5a, 0x69, 0x9b, 0xaf, 0x64, 0x6d, 0x72, 0xf8, 0x49, 0x8c, 0x5e, 0x39, 0xc9, 0x54, 0x95,
	0x6b, 0x53, 0x72, 0xa7, 0x74, 0x95, 0x14, 0x5a, 0xf0, 0x82, 0x01, 0x6e, 0xbb, 0x83, 0xc9, 0x9f,
	0x11, 0x79, 0xf7, 0x55, 0x09, 0xec, 0xe7, 0xa7, 0xbf, 0xde, 0xd0, 0x8f, 0xc8, 0x71, 0x48, 0xcc,
	0x2a, 0x5e, 0xca, 0x38, 0x1a, 0x47, 0xd3, 0xa3, 0x74, 0x18, 0xb0, 0x6b, 0x5e, 0x4a, 0x7a, 0x4a,
	0x06, 0x6b, 0x93, 0x23, 0xfd, 0x00, 0xe8, 0xfe, 0xda, 0xe4, 0x40, 0x9d, 0x90, 0x5e, 0x6d, 0x64,
	0xae, 0x36, 0xf1, 0x01, 0x10, 0x61, 0x44, 0x3f, 0x26, 0x23, 0xfc, 0x62, 0x85, 0xac, 0xee, 0xdc,
	0x32, 0x3e, 0x1c, 0x47, 0xd3, 0x51, 0x7a, 0x8c, 0xe0, 0x15, 0x60, 0x93, 0x3f, 0x0e, 0xc9, 0xc3,
	0x57, 0x45, 0x79, 0x3d, 0xb8, 0xaf, 0x90, 0xf7, 0x0c, 0xf5, 0x00, 0xf6, 0x1c, 0x93, 0x27, 0xe4,
	0x9d, 0x6e, 0x48, 0xb3, 0xc4, 0x39, 0x2c, 0xf1, 0x76, 0x27, 0x12, 0xd7, 0x69, 0x53, 0x96, 0xd2,
	0x19, 0x25, 0xe2, 0x0b, 0x08, 0xc4, 0x94, 0xcf, 0x00, 0xa2, 0x1f, 0x10, 0x82, 0x21, 0x6e, 0x5b,
	0xcb, 0xf8, 0x4b, 0x58, 0xf3, 0x08, 0x90, 0xdb, 0x6d, 0x2d, 0xe9, 0xa7, 0xe4, 0x2d, 0xa4, 0x85,
	0xae, 0x2a, 0x29, 0x9c, 0xcc, 0xe2, 0xaf, 0xc6, 0xd1, 0x74, 0x90, 0xbe, 0x09, 0xf0, 0x0f, 0x0d,
	0x4a, 0xff, 0x8a, 0x9a, 0x44, 0xbe, 0x2c, 0xf1, 0xd7, 0xe3, 0x68, 0x3a, 0x3c, 0xfb, 0x3d, 0x79,
	0x3d, 0x55, 0x4e, 0x3a, 0x87, 0xc9, 0x84, 0x2e, 0x4b, 0x5d, 0x85, 0x4d, 0xcc, 0xab, 0x5c, 0xd3,
	0x7f, 0xa2, 0x66, 0x17, 0x35, 0x77, 0x4b, 0x56, 0x28, 0xeb, 0xe2, 0x6f, 0xc6, 0x07, 0xd3, 0xe1,
	0xd9, 0xe6, 0x3e, 0x04, 0x7a, 0x11, 0xe9, 0x08, 0xab, 0xc5, 0xdd, 0xf2, 0x4a, 0x59, 0x37, 0x99,
	0x91, 0xe3, 0x5d, 0x88, 0x42, 0x7b, 0x59, 0x29, 0x74, 0x95, 0x81, 0x2d, 0x47, 0x69, 0x18, 0xd1,
	0x0f, 0x09, 0xa9, 0x78, 0xa5, 0x03, 0xf7, 0x00, 0xb8, 0x0e, 0x32, 0xf9, 0xaf, 0x47, 0xe8, 0xfe,
	0x61, 0xd0, 0x09, 0xc1, 0xf5, 0x18, 0x37, 0x92, 0x33, 0xd5, 0x64, 0x45, 0x27, 0x5c, 0x1a, 0xc9,
	0xe7, 0x19, 0xfd, 0xa4, 0x39, 0xa4, 0xd6, 0x2f, 0x98, 0x1f, 0xa7, 0xde, 0x36, 0x8e, 0xf9, 0x8c,
	0xa0, 0xd3, 0x98, 0x51, 0x0b, 0xb6, 0x96, 0xc6, 0x2a, 0x5d, 0xc1, 0x25, 0x18, 0xa5, 0x98, 0x20,
	0x55, 0x8b, 0x5f, 0x10, 0x6e, 0x63, 0xbd, 0xa4, 0x26, 0xd6, 0xdf, 0x88, 0xc3, 0x10, 0x7b, 0x53,
	0xe7, 0x4d, 0xec, 0x05, 0x39, 0xc1, 0xd8, 0x5c, 0x9b, 0xdf, 0xb8, 0xc9, 0x58, 0xa6, 0xac, 0xe3,
	0x95, 0x90, 0xf1, 0x1b, 0x90, 0xfc, 0x11, 0xb0, 0x33, 0x24, 0x7f, 0x0c, 0x5c, 0x6b, 0x71, 0xab,
	0x57, 0x46, 0xc8, 0xb8, 0xd7, 0xd9, 0xd8, 0x0d, 0x40, 0xbe, 0xfc, 0x41, 0xc5, 0xaa, 0xce, 0xb8,
	0xdf, 0xa0, 0x2a, 0x65, 0xdc, 0x07, 0x87, 0xba, 0xd7, 0x6e, 0x00, 0x55, 0xca, 0xb0, 0xf7, 0x17,
	0xa0, 0xe6, 0xd6, 0x97, 0xfb, 0xef, 0x9d, 0x43, 0x73, 0xae, 0x0a, 0x14, 0x38, 0xb8, 0x47, 0x81,
	0x58, 0xf2, 0x19, 0x57, 0x05, 0xc8, 0xfb, 0x9c, 0xd0, 0xb6, 0x8c, 0xb5, 0x51, 0xda, 0x28, 0xb7,
	0x8d, 0x8f, 0xe0, 0xa8, 0x1f, 0x36, 0x75, 0x7c, 0x1e, 0xf0, 0xb6, 0x4b, 0xf1, 0x95, 0xd3, 0x4c,
	0x6e, 0x44, 0xb1, 0xca, 0x64, 0x16, 0x13, 0xe8, 0x1b, 0x58, 0x89, 0xcb, 0x95, 0xd3, 0x4f, 0x03,
	0x41, 0xbf, 0x27, 0xef, 0x87, 0xec, 0xa6, 0x6d, 0x6d, 0x46, 0xde, 0x29, 0xeb, 0xa4, 0x91, 0x59,
	0x3c, 0x84, 0x89, 0xa7, 0xb8, 0x8e, 0x69, 0x5a, 0x5c, 0xba, 0x0b, 0xa0, 0xdf, 0x92, 0xd3, 0x4e,
	0x82, 0x6a, 0x61, 0xba, 0xb3, 0x8f, 0x41, 0xe5, 0xc9, 0x6e, 0xf6, 0xf5, 0xc2, 0xb4, 0x53, 0x27,
	0xff, 0x1e, 0xbc, 0xd4, 0x89, 0xe1, 0x6e, 0xd2, 0x2f, 0xc8, 0xa3, 0xe6, 0xe0, 0x9c, 0x34, 0x39,
	0x17, 0xb2, 0xfb, 0x42, 0xd0, 0xd0, 0x58, 0x02, 0x05, 0xaf, 0xc1, 0x79, 0xe3, 0xdd, 0x4a, 0x6e,
	0x1c, 0x5b, 0xea, 0x9a, 0xf1, 0x2c, 0x33, 0xd2, 0xda, 0xf0, 0x6c, 0xe0, 0x81, 0x5c, 0xcb, 0x8d,
	0xfb, 0x49, 0xd7, 0x97, 0x48, 0xed, 0x59, 0xf7, 0xa0, 0xd3, 0xf0, 0x83, 0x75, 0x9f, 0x10, 0x1c,
	0xb2, 0xc2, 0x72, 0x95, 0xc1, 0xcd, 0x39, 0x4a, 0xb1, 0xcf, 0x5e, 0x79, 0x84, 0x7e, 0x47, 0xde,
	0xeb, 0x74, 0x36, 0x65, 0x59, 0x29, 0xb8, 0x75, 0x5e, 0x38, 0x17, 0x0e, 0x2e, 0xce, 0x20, 0xec,
	0xdd, 0xb7, 0x9a, 0xb9, 0x7d, 0xe6, 0xe9, 0x39, 0xb0, 0xf4, 0x82, 0x3c, 0x7e, 0x79, 0xee, 0x4a,
	0x94, 0x78, 0x02, 0x70, 0x8b, 0x06, 0x41, 0x35, 0x4e, 0x7c, 0x21, 0xca, 0xda, 0x7f, 0xed, 0xbd,
	0x29, 0xfd, 0xfd, 0x37, 0xe5, 0x94, 0x0c, 0x0a, 0xcb, 0xf1, 0x45, 0x19, 0x00, 0xdd, 0x2f, 0x2c,
	0x87, 0xf7, 0xe4, 0x31, 0xe9, 0x37, 0x2d, 0x08, 0xed, 0xd3, 0xe3, 0xd8, 0x7d, 0x9e, 0x90, 0x21,
	0x10, 0xe8, 0xd4, 0x60, 0x16, 0xe2, 0xa1, 0x19, 0x20, 0x8b, 0x1e, 0xfc, 0x6f, 0x38, 0xff, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x7c, 0x0f, 0x19, 0x80, 0x51, 0x08, 0x00, 0x00,
}
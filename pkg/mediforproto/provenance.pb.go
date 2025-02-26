// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: medifor/v1/provenance.proto

package mediforproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LinkType int32

const (
	LinkType_SPLICE     LinkType = 0
	LinkType_PROVENANCE LinkType = 1
	LinkType_SPEAKER    LinkType = 2
	LinkType_CAMERA     LinkType = 3
	LinkType_LOCATION   LinkType = 4
)

// Enum value maps for LinkType.
var (
	LinkType_name = map[int32]string{
		0: "SPLICE",
		1: "PROVENANCE",
		2: "SPEAKER",
		3: "CAMERA",
		4: "LOCATION",
	}
	LinkType_value = map[string]int32{
		"SPLICE":     0,
		"PROVENANCE": 1,
		"SPEAKER":    2,
		"CAMERA":     3,
		"LOCATION":   4,
	}
)

func (x LinkType) Enum() *LinkType {
	p := new(LinkType)
	*p = x
	return p
}

func (x LinkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkType) Descriptor() protoreflect.EnumDescriptor {
	return file_medifor_v1_provenance_proto_enumTypes[0].Descriptor()
}

func (LinkType) Type() protoreflect.EnumType {
	return &file_medifor_v1_provenance_proto_enumTypes[0]
}

func (x LinkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkType.Descriptor instead.
func (LinkType) EnumDescriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{0}
}

// FilterRequest is used to ask a provenance analytic to return the
// top N matches for a given image.
type ProvenanceFilteringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the request.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// The image to be sent to the provenance analytic.
	Image *Resource `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// The number of results to return
	ResultLimit int32 `protobuf:"varint,3,opt,name=result_limit,json=resultLimit,proto3" json:"result_limit,omitempty"`
}

func (x *ProvenanceFilteringRequest) Reset() {
	*x = ProvenanceFilteringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvenanceFilteringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvenanceFilteringRequest) ProtoMessage() {}

func (x *ProvenanceFilteringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvenanceFilteringRequest.ProtoReflect.Descriptor instead.
func (*ProvenanceFilteringRequest) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{0}
}

func (x *ProvenanceFilteringRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ProvenanceFilteringRequest) GetImage() *Resource {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ProvenanceFilteringRequest) GetResultLimit() int32 {
	if x != nil {
		return x.ResultLimit
	}
	return 0
}

// ImageMatch contains the image identifier and score denoting how
// closely it matches the probe image.
type ImageMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the image.
	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	// The image resource.
	MatchingImage *Resource `protobuf:"bytes,2,opt,name=matching_image,json=matchingImage,proto3" json:"matching_image,omitempty"`
	// The score for the matching image.
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ImageMatch) Reset() {
	*x = ImageMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMatch) ProtoMessage() {}

func (x *ImageMatch) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMatch.ProtoReflect.Descriptor instead.
func (*ImageMatch) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{1}
}

func (x *ImageMatch) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *ImageMatch) GetMatchingImage() *Resource {
	if x != nil {
		return x.MatchingImage
	}
	return nil
}

func (x *ImageMatch) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// ImageFilter contains the top N results that match the probe image
// along with the corresponding scores.
type FilteringResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of matching images with scores.  The length of the list
	// is determined by the result_limit provided in the request.
	Matches []*ImageMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	// Probe image sent for the filter request, reference (Optional)/
	Probe *Resource `protobuf:"bytes,2,opt,name=probe,proto3" json:"probe,omitempty"`
}

func (x *FilteringResult) Reset() {
	*x = FilteringResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilteringResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilteringResult) ProtoMessage() {}

func (x *FilteringResult) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilteringResult.ProtoReflect.Descriptor instead.
func (*FilteringResult) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{2}
}

func (x *FilteringResult) GetMatches() []*ImageMatch {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *FilteringResult) GetProbe() *Resource {
	if x != nil {
		return x.Probe
	}
	return nil
}

// ProvenanceGraphRequest provides a provenance analytic with a probe image
// and the results of the filtering task to the analytic.
type ProvenanceGraphRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the request.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// The image to be sent to the provenance analytic.
	Image *Resource `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Image matches from the provenance filtering task
	FilterResults []*ImageMatch `protobuf:"bytes,3,rep,name=filter_results,json=filterResults,proto3" json:"filter_results,omitempty"`
}

func (x *ProvenanceGraphRequest) Reset() {
	*x = ProvenanceGraphRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvenanceGraphRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvenanceGraphRequest) ProtoMessage() {}

func (x *ProvenanceGraphRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvenanceGraphRequest.ProtoReflect.Descriptor instead.
func (*ProvenanceGraphRequest) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{3}
}

func (x *ProvenanceGraphRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ProvenanceGraphRequest) GetImage() *Resource {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ProvenanceGraphRequest) GetFilterResults() []*ImageMatch {
	if x != nil {
		return x.FilterResults
	}
	return nil
}

// Association provides details describing the relationship between two probes.
type Association struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The specific tyle of link created
	Type LinkType `protobuf:"varint,1,opt,name=type,proto3,enum=mediforproto.LinkType" json:"type,omitempty"`
	// UUID of the probe image
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// UUID of the matching image
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// Real number that is higher if the association exists.
	Score float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
	// Explanation of the association and supporting evidence.
	Explanation string `protobuf:"bytes,5,opt,name=explanation,proto3" json:"explanation,omitempty"`
}

func (x *Association) Reset() {
	*x = Association{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Association) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Association) ProtoMessage() {}

func (x *Association) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Association.ProtoReflect.Descriptor instead.
func (*Association) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{4}
}

func (x *Association) GetType() LinkType {
	if x != nil {
		return x.Type
	}
	return LinkType_SPLICE
}

func (x *Association) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Association) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Association) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Association) GetExplanation() string {
	if x != nil {
		return x.Explanation
	}
	return ""
}

// Association provides details describing the relationship between two probes.
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of the image
	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	// Real number that is higher if an image is in the graph
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{5}
}

func (x *Node) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *Node) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// ProvenanceGraph
type ProvenanceGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of nodes in the graph.
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// List of associations that specify the relationship between nodes on the gra
	Associations []*Association `protobuf:"bytes,2,rep,name=associations,proto3" json:"associations,omitempty"`
	// Whether the algorithm opted out of the task
	OptOut bool `protobuf:"varint,3,opt,name=opt_out,json=optOut,proto3" json:"opt_out,omitempty"`
}

func (x *ProvenanceGraph) Reset() {
	*x = ProvenanceGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_medifor_v1_provenance_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvenanceGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvenanceGraph) ProtoMessage() {}

func (x *ProvenanceGraph) ProtoReflect() protoreflect.Message {
	mi := &file_medifor_v1_provenance_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvenanceGraph.ProtoReflect.Descriptor instead.
func (*ProvenanceGraph) Descriptor() ([]byte, []int) {
	return file_medifor_v1_provenance_proto_rawDescGZIP(), []int{6}
}

func (x *ProvenanceGraph) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *ProvenanceGraph) GetAssociations() []*Association {
	if x != nil {
		return x.Associations
	}
	return nil
}

func (x *ProvenanceGraph) GetOptOut() bool {
	if x != nil {
		return x.OptOut
	}
	return false
}

var File_medifor_v1_provenance_proto protoreflect.FileDescriptor

var file_medifor_v1_provenance_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d,
	0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x65, 0x64,
	0x69, 0x66, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x76, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7c, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x3d,
	0x0a, 0x0e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0d,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x22, 0x73, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f,
	0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x3f, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66,
	0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x22, 0xa1, 0x01, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x93,
	0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x12, 0x28, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c,
	0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x70,
	0x74, 0x4f, 0x75, 0x74, 0x2a, 0x4d, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x50, 0x4c, 0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x50, 0x52, 0x4f, 0x56, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x50, 0x45, 0x41, 0x4b, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4d,
	0x45, 0x52, 0x41, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x04, 0x32, 0xcc, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x5e, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x5e, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x42, 0x5f, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f,
	0x72, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0f,
	0x50, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x66, 0x6f, 0x72, 0x65, 0x6e, 0x73, 0x69, 0x63, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x66,
	0x6f, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x66, 0x6f, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_medifor_v1_provenance_proto_rawDescOnce sync.Once
	file_medifor_v1_provenance_proto_rawDescData = file_medifor_v1_provenance_proto_rawDesc
)

func file_medifor_v1_provenance_proto_rawDescGZIP() []byte {
	file_medifor_v1_provenance_proto_rawDescOnce.Do(func() {
		file_medifor_v1_provenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_medifor_v1_provenance_proto_rawDescData)
	})
	return file_medifor_v1_provenance_proto_rawDescData
}

var file_medifor_v1_provenance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_medifor_v1_provenance_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_medifor_v1_provenance_proto_goTypes = []interface{}{
	(LinkType)(0),                      // 0: mediforproto.LinkType
	(*ProvenanceFilteringRequest)(nil), // 1: mediforproto.ProvenanceFilteringRequest
	(*ImageMatch)(nil),                 // 2: mediforproto.ImageMatch
	(*FilteringResult)(nil),            // 3: mediforproto.FilteringResult
	(*ProvenanceGraphRequest)(nil),     // 4: mediforproto.ProvenanceGraphRequest
	(*Association)(nil),                // 5: mediforproto.Association
	(*Node)(nil),                       // 6: mediforproto.Node
	(*ProvenanceGraph)(nil),            // 7: mediforproto.ProvenanceGraph
	(*Resource)(nil),                   // 8: mediforproto.Resource
}
var file_medifor_v1_provenance_proto_depIdxs = []int32{
	8,  // 0: mediforproto.ProvenanceFilteringRequest.image:type_name -> mediforproto.Resource
	8,  // 1: mediforproto.ImageMatch.matching_image:type_name -> mediforproto.Resource
	2,  // 2: mediforproto.FilteringResult.matches:type_name -> mediforproto.ImageMatch
	8,  // 3: mediforproto.FilteringResult.probe:type_name -> mediforproto.Resource
	8,  // 4: mediforproto.ProvenanceGraphRequest.image:type_name -> mediforproto.Resource
	2,  // 5: mediforproto.ProvenanceGraphRequest.filter_results:type_name -> mediforproto.ImageMatch
	0,  // 6: mediforproto.Association.type:type_name -> mediforproto.LinkType
	6,  // 7: mediforproto.ProvenanceGraph.nodes:type_name -> mediforproto.Node
	5,  // 8: mediforproto.ProvenanceGraph.associations:type_name -> mediforproto.Association
	1,  // 9: mediforproto.Provenance.ProvenanceFiltering:input_type -> mediforproto.ProvenanceFilteringRequest
	4,  // 10: mediforproto.Provenance.ProvenanceGraphBuilding:input_type -> mediforproto.ProvenanceGraphRequest
	3,  // 11: mediforproto.Provenance.ProvenanceFiltering:output_type -> mediforproto.FilteringResult
	7,  // 12: mediforproto.Provenance.ProvenanceGraphBuilding:output_type -> mediforproto.ProvenanceGraph
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_medifor_v1_provenance_proto_init() }
func file_medifor_v1_provenance_proto_init() {
	if File_medifor_v1_provenance_proto != nil {
		return
	}
	file_medifor_v1_analytic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_medifor_v1_provenance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvenanceFilteringRequest); i {
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
		file_medifor_v1_provenance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageMatch); i {
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
		file_medifor_v1_provenance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilteringResult); i {
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
		file_medifor_v1_provenance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvenanceGraphRequest); i {
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
		file_medifor_v1_provenance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Association); i {
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
		file_medifor_v1_provenance_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_medifor_v1_provenance_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvenanceGraph); i {
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
			RawDescriptor: file_medifor_v1_provenance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_medifor_v1_provenance_proto_goTypes,
		DependencyIndexes: file_medifor_v1_provenance_proto_depIdxs,
		EnumInfos:         file_medifor_v1_provenance_proto_enumTypes,
		MessageInfos:      file_medifor_v1_provenance_proto_msgTypes,
	}.Build()
	File_medifor_v1_provenance_proto = out.File
	file_medifor_v1_provenance_proto_rawDesc = nil
	file_medifor_v1_provenance_proto_goTypes = nil
	file_medifor_v1_provenance_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProvenanceClient is the client API for Provenance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProvenanceClient interface {
	ProvenanceFiltering(ctx context.Context, in *ProvenanceFilteringRequest, opts ...grpc.CallOption) (*FilteringResult, error)
	ProvenanceGraphBuilding(ctx context.Context, in *ProvenanceGraphRequest, opts ...grpc.CallOption) (*ProvenanceGraph, error)
}

type provenanceClient struct {
	cc grpc.ClientConnInterface
}

func NewProvenanceClient(cc grpc.ClientConnInterface) ProvenanceClient {
	return &provenanceClient{cc}
}

func (c *provenanceClient) ProvenanceFiltering(ctx context.Context, in *ProvenanceFilteringRequest, opts ...grpc.CallOption) (*FilteringResult, error) {
	out := new(FilteringResult)
	err := c.cc.Invoke(ctx, "/mediforproto.Provenance/ProvenanceFiltering", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provenanceClient) ProvenanceGraphBuilding(ctx context.Context, in *ProvenanceGraphRequest, opts ...grpc.CallOption) (*ProvenanceGraph, error) {
	out := new(ProvenanceGraph)
	err := c.cc.Invoke(ctx, "/mediforproto.Provenance/ProvenanceGraphBuilding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvenanceServer is the server API for Provenance service.
type ProvenanceServer interface {
	ProvenanceFiltering(context.Context, *ProvenanceFilteringRequest) (*FilteringResult, error)
	ProvenanceGraphBuilding(context.Context, *ProvenanceGraphRequest) (*ProvenanceGraph, error)
}

// UnimplementedProvenanceServer can be embedded to have forward compatible implementations.
type UnimplementedProvenanceServer struct {
}

func (*UnimplementedProvenanceServer) ProvenanceFiltering(context.Context, *ProvenanceFilteringRequest) (*FilteringResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvenanceFiltering not implemented")
}
func (*UnimplementedProvenanceServer) ProvenanceGraphBuilding(context.Context, *ProvenanceGraphRequest) (*ProvenanceGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvenanceGraphBuilding not implemented")
}

func RegisterProvenanceServer(s *grpc.Server, srv ProvenanceServer) {
	s.RegisterService(&_Provenance_serviceDesc, srv)
}

func _Provenance_ProvenanceFiltering_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvenanceFilteringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvenanceServer).ProvenanceFiltering(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediforproto.Provenance/ProvenanceFiltering",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvenanceServer).ProvenanceFiltering(ctx, req.(*ProvenanceFilteringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provenance_ProvenanceGraphBuilding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvenanceGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvenanceServer).ProvenanceGraphBuilding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediforproto.Provenance/ProvenanceGraphBuilding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvenanceServer).ProvenanceGraphBuilding(ctx, req.(*ProvenanceGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Provenance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mediforproto.Provenance",
	HandlerType: (*ProvenanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProvenanceFiltering",
			Handler:    _Provenance_ProvenanceFiltering_Handler,
		},
		{
			MethodName: "ProvenanceGraphBuilding",
			Handler:    _Provenance_ProvenanceGraphBuilding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medifor/v1/provenance.proto",
}

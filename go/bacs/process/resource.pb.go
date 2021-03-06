// Code generated by protoc-gen-go.
// source: bacs/process/resource.proto
// DO NOT EDIT!

/*
Package process is a generated protocol buffer package.

It is generated from these files:
	bacs/process/resource.proto
	bacs/process/source.proto

It has these top-level messages:
	ResourceLimits
	ResourceUsage
	Source
	Buildable
	BuildSettings
	BuilderConfig
	BuildResult
	ExecutionResult
*/
package process

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ResourceLimits struct {
	TimeLimitMillis     uint64 `protobuf:"varint,1,opt,name=time_limit_millis" json:"time_limit_millis,omitempty"`
	MemoryLimitBytes    uint64 `protobuf:"varint,2,opt,name=memory_limit_bytes" json:"memory_limit_bytes,omitempty"`
	OutputLimitBytes    uint64 `protobuf:"varint,3,opt,name=output_limit_bytes" json:"output_limit_bytes,omitempty"`
	NumberOfProcesses   uint64 `protobuf:"varint,4,opt,name=number_of_processes" json:"number_of_processes,omitempty"`
	RealTimeLimitMillis uint64 `protobuf:"varint,5,opt,name=real_time_limit_millis" json:"real_time_limit_millis,omitempty"`
}

func (m *ResourceLimits) Reset()         { *m = ResourceLimits{} }
func (m *ResourceLimits) String() string { return proto.CompactTextString(m) }
func (*ResourceLimits) ProtoMessage()    {}

type ResourceUsage struct {
	// TODO: do we need to export Process::ResourceUsage::timeUsage?
	TimeUsageMillis  uint64 `protobuf:"varint,1,opt,name=time_usage_millis" json:"time_usage_millis,omitempty"`
	MemoryUsageBytes uint64 `protobuf:"varint,2,opt,name=memory_usage_bytes" json:"memory_usage_bytes,omitempty"`
}

func (m *ResourceUsage) Reset()         { *m = ResourceUsage{} }
func (m *ResourceUsage) String() string { return proto.CompactTextString(m) }
func (*ResourceUsage) ProtoMessage()    {}

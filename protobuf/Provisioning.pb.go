// Code generated by protoc-gen-go.
// source: Provisioning.proto
// DO NOT EDIT!

package textsecure

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProvisionEnvelope struct {
	PublicKey        []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`
	Body             []byte `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProvisionEnvelope) Reset()         { *m = ProvisionEnvelope{} }
func (m *ProvisionEnvelope) String() string { return proto.CompactTextString(m) }
func (*ProvisionEnvelope) ProtoMessage()    {}

func (m *ProvisionEnvelope) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ProvisionEnvelope) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type ProvisionMessage struct {
	IdentityKeyPublic  []byte  `protobuf:"bytes,1,opt,name=identityKeyPublic" json:"identityKeyPublic,omitempty"`
	IdentityKeyPrivate []byte  `protobuf:"bytes,2,opt,name=identityKeyPrivate" json:"identityKeyPrivate,omitempty"`
	Number             *string `protobuf:"bytes,3,opt,name=number" json:"number,omitempty"`
	ProvisioningCode   *string `protobuf:"bytes,4,opt,name=provisioningCode" json:"provisioningCode,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *ProvisionMessage) Reset()         { *m = ProvisionMessage{} }
func (m *ProvisionMessage) String() string { return proto.CompactTextString(m) }
func (*ProvisionMessage) ProtoMessage()    {}

func (m *ProvisionMessage) GetIdentityKeyPublic() []byte {
	if m != nil {
		return m.IdentityKeyPublic
	}
	return nil
}

func (m *ProvisionMessage) GetIdentityKeyPrivate() []byte {
	if m != nil {
		return m.IdentityKeyPrivate
	}
	return nil
}

func (m *ProvisionMessage) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

func (m *ProvisionMessage) GetProvisioningCode() string {
	if m != nil && m.ProvisioningCode != nil {
		return *m.ProvisioningCode
	}
	return ""
}

func init() {
}

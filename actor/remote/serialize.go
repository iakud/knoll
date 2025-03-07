package remote

import (
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Serializer interface {
	Serialize(msg any) (string, []byte, error)
	Deserialize(typeName string, data []byte) (any, error)
}

type ProtoSerializer struct{}

func (ProtoSerializer) Serialize(msg any) (string, []byte, error) {
	pm := msg.(proto.Message)
	data, err := proto.Marshal(pm)
	if err != nil {
		return "", nil, err
	}
	typeName := string(proto.MessageName(pm))
	return typeName, data, err
}

func (ProtoSerializer) Deserialize(typeName string, data []byte) (any, error) {
	n, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(typeName))
	if err != nil {
		return nil, err
	}
	pm := n.New().Interface()
	err = proto.Unmarshal(data, pm)
	return pm, err
}

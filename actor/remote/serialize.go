package remote

import (
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type Serializer interface {
	Serialize(msg any) ([]byte, error)
	TypeName(any) string
	Deserialize(string, []byte) (any, error)
}

type ProtoSerializer struct{}

func (ProtoSerializer) Serialize(msg any) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (ProtoSerializer) Deserialize(tname string, data []byte) (any, error) {
	pname := protoreflect.FullName(tname)
	n, err := protoregistry.GlobalTypes.FindMessageByName(pname)
	if err != nil {
		return nil, err
	}
	pm := n.New().Interface()
	err = proto.Unmarshal(data, pm)
	return pm, err
}

func (ProtoSerializer) TypeName(msg any) string {
	return string(proto.MessageName(msg.(proto.Message)))
}

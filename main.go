package main

import (
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/simplesteph/protobuf-example-go/src/complex"
	"github.com/simplesteph/protobuf-example-go/src/simple"
	"io/ioutil"
	"log"
	enumpb "protobuf/src/enum_example"
)

func main() {
	//sm := doSimple()
	//
	//readAndWriteDemo(sm)
	//jsonDemo(sm)

	doEnum()

	//doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	em := &enumpb.EnumMessage{}
	var in = "08021000184b2080c2d72f2a20f0321539a45078365c1a65944d010876c0efe45c0446101dacced7a2f29aa289321e088094ebdc03121473e0862d85828b12c1e8ae890844a39293ac16731a003a40a1b19c4195a98f900bac8e7060f5a09ae929d7f47a6ef5a4efe100f80e5d1272955aa88fc93587c59c912f2c18c389c6ddd1f76bfe2a2fa6b1a1c8d3b5341606"
	data, _ := hex.DecodeString(in)
	if err := proto.Unmarshal(data, em); err != nil {
		log.Fatalln("Failed to parse  transaction:", err)
	}

	em1 := &enumpb.AssetMessage{}
	if err := proto.Unmarshal(em.Asset, em1); err != nil {
		log.Fatalln("Failed to parse  transaction:", err)
	}

	fmt.Println(hex.EncodeToString(em1.RecipientAddress))
	//em := &enumpb.EnumMessage{}
	//var var1 uint32 = 10000
	//var var2 uint64 = 10000
	//
	//em1 := &enumpb.EnumMessage{
	//	ModuleID:        &var1,
	//	AssetID:         &var1,
	//	Fee:             &var2,
	//	Nonce:           &var2,
	//	SenderPublicKey: []byte("8f057d088a585d938c20d63e430a068d4cea384e588aa0b758c68fca21644dbc"),
	//	Asset:           []byte("f214d75bbc4b2ea89e433f3a45af803725416ec3"),
	//
	//}
	//em1.Signatures = append(em1.Signatures, []byte("204514eb1152355799ece36d17037e5feb4871472c60763bdafe67eb6a38bec632a8e2e62f84a32cf764342a4708a65fbad194e37feec03940f0ff84d3df2a05") )
	//em1.Signatures = append(em1.Signatures, []byte("0b6730e5898ca56fe0dc1c73de9363f6fc8b335592ef10725a8463bff101a4943e60311f0b1a439a2c9e02cca1379b80a822f4ec48cf212bff1f1c757e92ec02") )
	//
	//out, _ := proto.Marshal(em1)
	//if err := proto.Unmarshal(out, em); err != nil {
	//	log.Fatalln("Failed to parse  transaction:", err)
	//}
	//fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	fmt.Println(out)
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is:", sm.GetId())

	return &sm
}

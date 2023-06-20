package sample

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

)

const (
	ERR_GENERIC        = "Generic error"
	ERR_INVALID_PARAMS = "Invalid parameters"
	BYTESIZE           = 8
	MAX_BITWIDTH       = 32
	BUFF_SIZE		= 16
)

var maskMap = map[int]int{
	0: 0x0000,
	1: 0x0001,
	2: 0x0003,
	3: 0x0007,
	4: 0x000F,
	5: 0x001F,
	6: 0x003F,
	7: 0x007F,
	8: 0x00FF,
}

type s1 struct {
	A uint8
	B uint16
	C int32
}

func run() {

}

func A() {
	fmt.Printf("Running A .. \n")
	st := []byte(
		`{"A":45 , "B":55 , "C":10000}`)
	var s s1
	err := json.Unmarshal(st, &s)
	fmt.Printf("err : %v , s : %v \n", err, s)
}

func B(pos int, length int, value int) (value_out int , err error){
	bs := []byte{0xFF, 0xFF, 0xFF}
	bs2, pos2, err := SerializeField(value, pos, length, bs)
	fmt.Printf("err : %v , pos2 : %v , bs2 : %v \n", err, pos2, bs2)
	value , pos_out , err := DeSerializeBitField(pos , length , bs2)
	log("value : %v , pos_out : %v , err : %v \n", value , pos_out , err)
	value_out=value
	return value_out , err
}


func C ()(){
	// bitset.LittleEndian()
	// b := bitset.BitSet.From(0)
	// pos := 7
	// length := 6
	// value := int(0xA)
	// Serialize(value , pos , length , b)
}

func DeSerializeBitField(position int, length int, bs []byte) (value int, position_out int, err error) {
	position_out = position
	if !(((len(bs) * BYTESIZE) >= (position + length)) && (length > 0) && (length < MAX_BITWIDTH)) {
		return -1, position_out, errors.New(ERR_INVALID_PARAMS)
	}
	local_buffer := make([]byte , BUFF_SIZE)
	copy(local_buffer , bs)
	log("pos : %v , length : %v " , position , length)
	target_fb := 0
	source_fb := position/BYTESIZE
	source_lb := (position + length - 1) / BYTESIZE
	target_lb := length / BYTESIZE
	// local_buffer[target_fb] , err = UnMaskFirstByte(uint8(position) , uint8(length) , bs[source_fb])
	// log("firstbyte err : %v , local_buffer : %v \n" , err , local_buffer)
	// local_buffer[target_lb] , err = UnMaskLastByte(uint8(position) , uint8(length) , bs[source_lb])
	// log("lastbyte err : %v , local_buffer : %v \n" , err , local_buffer)
	log("source_lb : %v " , source_lb)
	for i:= 0 ; i< (target_lb - target_fb ) ; i++ {
		local_buffer[i] = bs[source_fb + i]
	}
	position_out = position + length
	mask := maskMap[length]
	var value_local uint64
	err = binary.Read(bytes.NewReader(local_buffer[:]), binary.LittleEndian, &value_local)
	log("\nvalue : %v \n" , value_local)
	value= int(value_local)
	value = value >> (position % BYTESIZE)
	value = value & mask
	// value = value & mask
	
	return value, position_out, err
}

func SerializeField(value int, position int, length int, bs []byte) (bs_out []byte, position_out int, err error) {
	// Convert value to the little endian/ big endian encoding. ?
	// How is the bit position defined for different encodings?
	if !(((len(bs) * BYTESIZE) >= (position + length)) && (length > 0) && (length < MAX_BITWIDTH)) {
		return bs, position, errors.New(ERR_INVALID_PARAMS)
	}
	mask := maskMap[length]

	// truncate value
	value = value & mask
	bitshift := position % BYTESIZE

	// TODO: Left shift regardless of the little / big endian ?
	value = value << bitshift
	bs_out = make([]byte, 8)
	copy(bs_out , bs)
	log("Copiedbs : %v to bs_out : %v \n" , bs , bs_out)
	value_bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(value_bs, uint64(value))
	log("value : %v , value_bs : %v \n" , value, value_bs)

	source_lb := (bitshift+ length - 1) / BYTESIZE
	// Identify start and end byte.
	fb := position / BYTESIZE
	target_lb := (position + length - 1) / BYTESIZE
	log("source_lb : %v , fb : %v , target_lb : %v \n" , source_lb, fb , target_lb)
	bs_out[fb] , err = MaskFirstByte(uint8(position) , uint8(length) , bs_out[fb] , value_bs[0])
	log("err : %v , bs_out : %v \n" , err, bs_out[fb])
	bs_out[target_lb] , err = MaskLastByte(uint8(position), uint8(length) , bs_out[source_lb] ,  value_bs[source_lb]) 
	log("err : %v , bs_out : %v \n" , err, bs_out[fb])
	for i := 1 ; i < ((length / BYTESIZE) -1 ); i++ {
		bs_out[fb+i] = value_bs[i]
	}
	position_out = (position + length - 1) % BYTESIZE
	log("bs_out : %v \n" , bs_out)
	return bs_out, position_out, nil
}

func MaskFirstByte(position uint8 , length uint8 , in byte , value byte)(out byte , err error){
	position = position % BYTESIZE
	mask := byte(maskMap[int(length)])
	mask = mask << position
	out = in & (^mask)
	out = out |  (value)
	log("Mask : %v , in : %v ,  out : %v , position : %v , length : %v, value: %v   \n" , 
			mask , in,  out , position , length ,value)
	return out, nil
}

func MaskLastByte(position uint8 , length uint8 , in byte , value byte)(out byte , err error){
	out = in
	mask := byte(0)
	position = position % BYTESIZE
	if ( (position / BYTESIZE) != ( (position+length-1)/BYTESIZE)){
		length = (position+length)%BYTESIZE
		mask = byte(maskMap[int(length)])
		mask = mask << 0
		out = in & (^mask) 
		out = out | value
		log("Mask : %v , in : %v ,  out : %v , position : %v , length : %v, value: %v   \n" , 
				mask , in,  out , position , length ,value)
	} else {
		return in , errors.New(ERR_INVALID_PARAMS)
	}
	log("Mask : %v , out : %v  \n" , mask , out)
	return out, nil
}

func UnMaskFirstByte(position uint8, length uint8 , in byte)(out byte , err error){
	position = position % BYTESIZE
	mask := byte(maskMap[int(length)])
	mask = mask << position
	out = in & mask
	return out , nil
}

func UnMaskLastByte (position uint8, length uint8 , in byte)(out byte , err error){
	if ( (position / BYTESIZE) != ( (position+length-1)/BYTESIZE)){
		length_updated := (position + length ) % BYTESIZE
		mask := byte(maskMap[int(length_updated)])
		mask = mask << 0 
		out = in & mask
		return out , nil
	} 
	return in , errors.New(ERR_INVALID_PARAMS)
}

// func Serialize(value int , position int , length int, bs bitset.BitSet) (bs_out bitset.BitSet, position_out int, err error) {
// 	v1 := bitset.BitSet.From([]uint64{uint64(value)})
// }


func log (s string, args ... interface{}) {
	fmt.Printf(s, args...)
}



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
	0: 	0x00000000,
	1: 	0x00000001,
	2: 	0x00000003,
	3: 	0x00000007,
	4: 	0x0000000F,
	5: 	0x0000001F,
	6: 	0x0000003F,
	7: 	0x0000007F,
	8: 	0x000000FF,
	9: 	0x000001FF,
	10: 0x000003FF,
	11: 0x000007FF,
	12: 0x00000FFF,
	13: 0x00001FFF,
	14: 0x00003FFF,
	15: 0x00007FFF,
	16: 0x0000FFFF,
	17: 0x0001FFFF,
	18: 0x0003FFFF,
	19: 0x0007FFFF,
	20: 0x000FFFFF,
	21: 0x001FFFFF,
	22: 0x003FFFFF,
	23: 0x007FFFFF,
	24: 0x00FFFFFF,
	25: 0x01FFFFFF,
	26: 0x03FFFFFF,
	27: 0x07FFFFFF,
	28: 0x0FFFFFFF,
	29: 0x1FFFFFFF,
	30: 0x3FFFFFFF,
	31: 0x7FFFFFFF,
	32: 0xFFFFFFFF,
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
	bs := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, }
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
	// copy(local_buffer , bs)
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
	for i:= 0 ; i< (target_lb - target_fb +1) ; i++ {
		local_buffer[i] = bs[source_fb + i]
	}
	position_out = position + length
	mask := maskMap[length]
	var value_local uint64
	err = binary.Read(bytes.NewReader(local_buffer[:]), binary.LittleEndian, &value_local)
	log("\n local_buffer:%v, value_local:%v  , position:%v , mask:%v \n" ,local_buffer, value_local , position , mask)
	value= int(value_local)
	value = value >> (position % BYTESIZE)
	value = value & mask
	// value = value & mask
	
	return value, position_out, err
}

func SerializeField(value int, position int, length int, bs []byte) (bs_out []byte, position_out int, err error) {
	// TODO : Convert value to the little endian/ big endian encoding. ?
	// How is the bit position defined for different encodings?
	if !(((len(bs) * BYTESIZE) >= (position + length)) && (length > 0) && (length < MAX_BITWIDTH)) {
		return bs, position, errors.New(ERR_INVALID_PARAMS)
	}
	mask := maskMap[length]
	local_buffer := make([]byte , BUFF_SIZE)
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
	target_lb := (position + length - 1) / BYTESIZE
	target_fb := position/BYTESIZE
	copy(local_buffer[0:] , bs[target_fb:])
	log("source_lb:%v , target_fb:%v , target_lb:%v bitshift:%v len:%v\n" , source_lb, target_fb , target_lb , bitshift , length)
	local_buffer[0] , err = MaskFirstByte(uint8(position) , uint8(length) , local_buffer[0] , value_bs[0])
	log("err : %v , local_buffer : %v \n" , err, local_buffer[0])
	local_buffer[source_lb] , err = MaskLastByte(uint8(position), uint8(length) , local_buffer[source_lb] ,  value_bs[source_lb]) 
	log("err : %v , local_buffer : %v \n" , err, local_buffer[source_lb])
	lastIndex := 0
	if ((bitshift+length)%BYTESIZE == 0) {
		lastIndex = ((bitshift+length)/BYTESIZE ) -1
	} else {
		lastIndex = (bitshift+length)/BYTESIZE
	}
	for i := 1 ; i < lastIndex  ; i++ {
		local_buffer[i] = value_bs[i]
		log("i:%v , byte:%v\n" , i , local_buffer[i])
	}
	position_out = (position + length - 1) % BYTESIZE
	copy(bs_out[target_fb:target_lb+1] , local_buffer[0:(source_lb+1)])
	log("lastIndex:%v , bs_out:%v , local_buffer:%v  \n" ,lastIndex, bs_out , local_buffer)
	return bs_out, position_out, nil
}

func MaskFirstByte(position uint8 , length uint8 , in byte , value byte)(out byte , err error){
	position = position % BYTESIZE
	if ( (position / BYTESIZE) != ( (position+length-1)/BYTESIZE) ) {
		length = BYTESIZE - position
	}  
	mask := byte(maskMap[int(length)])
	mask = mask << position
	out = in & (^mask)
	out = out |  value
	log("Maskfirst : Mask : %v , in : %v ,  out : %v , position : %v , length : %v, value: %v   \n" , 
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
		log("MaskLast : Mask : %v , in : %v ,  out : %v , position : %v , length : %v, value: %v   \n" , 
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



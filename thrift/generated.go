
package thrift
import "github.com/v2pro/wombat/generic"
import "github.com/thrift-iterator/go-benchmark/thrift/media"
import "github.com/thrift-iterator/go/protocol/binary"
func init() {
generic.RegisterExpandedFunc("Decode_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator",Decode_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator)}
func DecodeSimpleValue_DT_ptr_string_ST_ptr_binary__Iterator(dst *string,src *binary.Iterator){
*dst = string(src.ReadString())
	
}
func DecodeAnything_DT_ptr_string_ST_ptr_binary__Iterator(dst *string,src *binary.Iterator){


DecodeSimpleValue_DT_ptr_string_ST_ptr_binary__Iterator(dst, src)

}
func DecodePointer_DT_ptr_ptr_string_ST_ptr_binary__Iterator(dst **string,src *binary.Iterator){

defDst := new(string)
DecodeAnything_DT_ptr_string_ST_ptr_binary__Iterator(defDst, src)
*dst = defDst
return
}
func DecodeAnything_DT_ptr_ptr_string_ST_ptr_binary__Iterator(dst **string,src *binary.Iterator){


DecodePointer_DT_ptr_ptr_string_ST_ptr_binary__Iterator(dst, src)

}
func DecodeSimpleValue_DT_ptr_media__Int_ST_ptr_binary__Iterator(dst *media.Int,src *binary.Iterator){
*dst = media.Int(src.ReadInt32())
	
}
func DecodeAnything_DT_ptr_media__Int_ST_ptr_binary__Iterator(dst *media.Int,src *binary.Iterator){


DecodeSimpleValue_DT_ptr_media__Int_ST_ptr_binary__Iterator(dst, src)

}
func DecodeEnum_DT_ptr_media__Size_ST_ptr_binary__Iterator(dst *media.Size,src *binary.Iterator){
*dst = media.Size(src.ReadInt32())
	
}
func DecodeAnything_DT_ptr_media__Size_ST_ptr_binary__Iterator(dst *media.Size,src *binary.Iterator){


DecodeEnum_DT_ptr_media__Size_ST_ptr_binary__Iterator(dst, src)

}
func DecodeStruct_DT_ptr_media__Image_ST_ptr_binary__Iterator(dst *media.Image,src *binary.Iterator){


	
	

	
	

	
	

	
	

	
	

src.ReadStructHeader()
for {
	fieldType, fieldId := src.ReadStructField()
	if fieldType == 0 {
		return
	}
	switch fieldId {
		
			case 1:
				DecodeAnything_DT_ptr_string_ST_ptr_binary__Iterator(&dst.URI, src)
		
			case 2:
				DecodeAnything_DT_ptr_ptr_string_ST_ptr_binary__Iterator(&dst.Title, src)
		
			case 3:
				DecodeAnything_DT_ptr_media__Int_ST_ptr_binary__Iterator(&dst.Width, src)
		
			case 4:
				DecodeAnything_DT_ptr_media__Int_ST_ptr_binary__Iterator(&dst.Height, src)
		
			case 5:
				DecodeAnything_DT_ptr_media__Size_ST_ptr_binary__Iterator(&dst.Size, src)
		
		default:
			src.Discard(fieldType)
	}
}
}
func DecodeAnything_DT_ptr_media__Image_ST_ptr_binary__Iterator(dst *media.Image,src *binary.Iterator){


DecodeStruct_DT_ptr_media__Image_ST_ptr_binary__Iterator(dst, src)

}
func DecodePointer_DT_ptr_ptr_media__Image_ST_ptr_binary__Iterator(dst **media.Image,src *binary.Iterator){

defDst := new(media.Image)
DecodeAnything_DT_ptr_media__Image_ST_ptr_binary__Iterator(defDst, src)
*dst = defDst
return
}
func DecodeAnything_DT_ptr_ptr_media__Image_ST_ptr_binary__Iterator(dst **media.Image,src *binary.Iterator){


DecodePointer_DT_ptr_ptr_media__Image_ST_ptr_binary__Iterator(dst, src)

}
func DecodeSlice_DT_ptr_slice_ptr_media__Image_ST_ptr_binary__Iterator(dst *[]*media.Image,src *binary.Iterator){

_, length := src.ReadListHeader()
for i := 0; i < length; i++ {
	elem := new(*media.Image)
	DecodeAnything_DT_ptr_ptr_media__Image_ST_ptr_binary__Iterator(elem, src)
	*dst = append(*dst, *elem)
}
}
func DecodeAnything_DT_ptr_slice_ptr_media__Image_ST_ptr_binary__Iterator(dst *[]*media.Image,src *binary.Iterator){


DecodeSlice_DT_ptr_slice_ptr_media__Image_ST_ptr_binary__Iterator(dst, src)

}
func DecodeSimpleValue_DT_ptr_media__Long_ST_ptr_binary__Iterator(dst *media.Long,src *binary.Iterator){
*dst = media.Long(src.ReadInt64())
	
}
func DecodeAnything_DT_ptr_media__Long_ST_ptr_binary__Iterator(dst *media.Long,src *binary.Iterator){


DecodeSimpleValue_DT_ptr_media__Long_ST_ptr_binary__Iterator(dst, src)

}
func DecodePointer_DT_ptr_ptr_media__Int_ST_ptr_binary__Iterator(dst **media.Int,src *binary.Iterator){

defDst := new(media.Int)
DecodeAnything_DT_ptr_media__Int_ST_ptr_binary__Iterator(defDst, src)
*dst = defDst
return
}
func DecodeAnything_DT_ptr_ptr_media__Int_ST_ptr_binary__Iterator(dst **media.Int,src *binary.Iterator){


DecodePointer_DT_ptr_ptr_media__Int_ST_ptr_binary__Iterator(dst, src)

}
func DecodeSlice_DT_ptr_slice_string_ST_ptr_binary__Iterator(dst *[]string,src *binary.Iterator){

_, length := src.ReadListHeader()
for i := 0; i < length; i++ {
	elem := new(string)
	DecodeAnything_DT_ptr_string_ST_ptr_binary__Iterator(elem, src)
	*dst = append(*dst, *elem)
}
}
func DecodeAnything_DT_ptr_slice_string_ST_ptr_binary__Iterator(dst *[]string,src *binary.Iterator){


DecodeSlice_DT_ptr_slice_string_ST_ptr_binary__Iterator(dst, src)

}
func DecodeEnum_DT_ptr_media__Player_ST_ptr_binary__Iterator(dst *media.Player,src *binary.Iterator){
*dst = media.Player(src.ReadInt32())
	
}
func DecodeAnything_DT_ptr_media__Player_ST_ptr_binary__Iterator(dst *media.Player,src *binary.Iterator){


DecodeEnum_DT_ptr_media__Player_ST_ptr_binary__Iterator(dst, src)

}
func DecodeStruct_DT_ptr_media__Media_ST_ptr_binary__Iterator(dst *media.Media,src *binary.Iterator){


	
	

	
	

	
	

	
	

	
	

	
	

	
	

	
	

	
	

	
	

	
	

src.ReadStructHeader()
for {
	fieldType, fieldId := src.ReadStructField()
	if fieldType == 0 {
		return
	}
	switch fieldId {
		
			case 1:
				DecodeAnything_DT_ptr_string_ST_ptr_binary__Iterator(&dst.URI, src)
		
			case 2:
				DecodeAnything_DT_ptr_ptr_string_ST_ptr_binary__Iterator(&dst.Title, src)
		
			case 3:
				DecodeAnything_DT_ptr_media__Int_ST_ptr_binary__Iterator(&dst.Width, src)
		
			case 4:
				DecodeAnything_DT_ptr_media__Int_ST_ptr_binary__Iterator(&dst.Height, src)
		
			case 5:
				DecodeAnything_DT_ptr_string_ST_ptr_binary__Iterator(&dst.Format, src)
		
			case 6:
				DecodeAnything_DT_ptr_media__Long_ST_ptr_binary__Iterator(&dst.Duration, src)
		
			case 7:
				DecodeAnything_DT_ptr_media__Long_ST_ptr_binary__Iterator(&dst.Size, src)
		
			case 8:
				DecodeAnything_DT_ptr_ptr_media__Int_ST_ptr_binary__Iterator(&dst.Bitrate, src)
		
			case 9:
				DecodeAnything_DT_ptr_slice_string_ST_ptr_binary__Iterator(&dst.Person, src)
		
			case 10:
				DecodeAnything_DT_ptr_media__Player_ST_ptr_binary__Iterator(&dst.Player, src)
		
			case 11:
				DecodeAnything_DT_ptr_ptr_string_ST_ptr_binary__Iterator(&dst.Copyright, src)
		
		default:
			src.Discard(fieldType)
	}
}
}
func DecodeAnything_DT_ptr_media__Media_ST_ptr_binary__Iterator(dst *media.Media,src *binary.Iterator){


DecodeStruct_DT_ptr_media__Media_ST_ptr_binary__Iterator(dst, src)

}
func DecodePointer_DT_ptr_ptr_media__Media_ST_ptr_binary__Iterator(dst **media.Media,src *binary.Iterator){

defDst := new(media.Media)
DecodeAnything_DT_ptr_media__Media_ST_ptr_binary__Iterator(defDst, src)
*dst = defDst
return
}
func DecodeAnything_DT_ptr_ptr_media__Media_ST_ptr_binary__Iterator(dst **media.Media,src *binary.Iterator){


DecodePointer_DT_ptr_ptr_media__Media_ST_ptr_binary__Iterator(dst, src)

}
func DecodeStruct_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator(dst *media.MediaContent,src *binary.Iterator){


	
	

	
	

src.ReadStructHeader()
for {
	fieldType, fieldId := src.ReadStructField()
	if fieldType == 0 {
		return
	}
	switch fieldId {
		
			case 1:
				DecodeAnything_DT_ptr_slice_ptr_media__Image_ST_ptr_binary__Iterator(&dst.Image, src)
		
			case 2:
				DecodeAnything_DT_ptr_ptr_media__Media_ST_ptr_binary__Iterator(&dst.Media, src)
		
		default:
			src.Discard(fieldType)
	}
}
}
func DecodeAnything_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator(dst *media.MediaContent,src *binary.Iterator){


DecodeStruct_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator(dst, src)

}
func Decode_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator(dst interface{},src interface{}){

DecodeAnything_DT_ptr_media__MediaContent_ST_ptr_binary__Iterator(dst.(*media.MediaContent), src.(*binary.Iterator))

}
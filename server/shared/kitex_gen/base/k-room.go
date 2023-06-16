// Code generated by Kitex v0.5.2. DO NOT EDIT.

package base

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
)

func (p *Room) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.I32 {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 7:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField7(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Room[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Room) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Name = v

	}
	return offset, nil
}

func (p *Room) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.MaxPlayer = v

	}
	return offset, nil
}

func (p *Room) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CurrentPlayer = v

	}
	return offset, nil
}

func (p *Room) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Host = v

	}
	return offset, nil
}

func (p *Room) FastReadField6(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI32(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Port = v

	}
	return offset, nil
}

func (p *Room) FastReadField7(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.RoomId = v

	}
	return offset, nil
}

// for compatibility
func (p *Room) FastWrite(buf []byte) int {
	return 0
}

func (p *Room) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Room")
	if p != nil {
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
		offset += p.fastWriteField7(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Room) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Room")
	if p != nil {
		l += p.field1Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
		l += p.field7Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Room) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "name", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Name)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Room) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "max_player", thrift.I32, 3)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.MaxPlayer)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Room) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "current_player", thrift.I32, 4)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.CurrentPlayer)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Room) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "host", thrift.STRING, 5)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Host)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Room) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "port", thrift.I32, 6)
	offset += bthrift.Binary.WriteI32(buf[offset:], p.Port)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Room) fastWriteField7(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "room_id", thrift.I64, 7)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.RoomId)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Room) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("name", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.Name)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Room) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("max_player", thrift.I32, 3)
	l += bthrift.Binary.I32Length(p.MaxPlayer)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Room) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("current_player", thrift.I32, 4)
	l += bthrift.Binary.I32Length(p.CurrentPlayer)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Room) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("host", thrift.STRING, 5)
	l += bthrift.Binary.StringLengthNocopy(p.Host)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Room) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("port", thrift.I32, 6)
	l += bthrift.Binary.I32Length(p.Port)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Room) field7Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("room_id", thrift.I64, 7)
	l += bthrift.Binary.I64Length(p.RoomId)

	l += bthrift.Binary.FieldEndLength()
	return l
}
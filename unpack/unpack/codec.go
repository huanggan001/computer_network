package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

// 定义固定的消息头，长度为8字节
const Msg_Header = "12345678"

// Encode 将给定的内容按照特定的消息协议编码到 io.Writer 中
// 消息协议结构: 消息头 (Msg_Header) + 内容长度 (content_length) + 内容 (content)
func Encode(bytesBuffer io.Writer, content string) error {
	// 写入消息头（Msg_Header）到输出流
	// 消息头的长度是固定的，可以用于校验接收的数据是否符合协议
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err // 如果写入失败，返回错误
	}

	// 计算内容长度，最长 4 字节（int32），即最大可表示2^31-1长度的字节
	content_len := int32(len([]byte(content)))

	// 写入内容长度到输出流，采用大端序 (Big Endian) 进行编码
	// 这样接收方可以知道后续数据的字节长度
	//content_len 写入时占用 固定的 4 字节，即使内容本身的实际长度不满 4 字节
	if err := binary.Write(bytesBuffer, binary.BigEndian, content_len); err != nil {
		return err // 如果写入失败，返回错误
	}

	// 写入实际的消息内容到输出流
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err // 如果写入失败，返回错误
	}

	// 如果一切顺利，返回 nil，表示没有错误
	return nil
}

// Decode 从 io.Reader 中读取数据，并根据消息协议解码内容
// 返回解码后的消息内容 (bodyBuf) 和可能出现的错误
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	// 读取消息头，首先申请一个和消息头长度相等的缓冲区
	MagicBuf := make([]byte, len(Msg_Header))

	// 从输入流中读取与消息头相同长度的数据到 MagicBuf
	if _, err = io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err // 如果读取失败，返回错误
	}

	// 检查读取到的消息头是否与预定义的消息头 (Msg_Header) 一致
	if string(MagicBuf) != Msg_Header {
		return nil, errors.New("msg_header error") // 如果不匹配，返回错误
	}

	// 申请4字节的缓冲区，用来读取内容长度（int32）
	lengthBuf := make([]byte, 4)

	// 从输入流中读取内容长度 (4字节)，存入 lengthBuf
	if _, err = io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err // 如果读取失败，返回错误
	}

	// 将读取到的4字节数据转换为大端序的长度值
	length := binary.BigEndian.Uint32(lengthBuf)

	// 根据读取到的长度值，申请一个新的缓冲区用于存放消息内容
	bodyBuf = make([]byte, length)

	// 从输入流中读取与内容长度相等的数据到 bodyBuf
	if _, err = io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err // 如果读取失败，返回错误
	}

	// 返回读取到的消息内容和可能的错误
	return bodyBuf, err
}

package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"chatPro/common/message"
	"net"
)
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
	// Pkglen uint32
}

func (this *Transfer) ReadPkg() (msg message.Message, err error) {
	// buf := make([]byte, 8096)
	fmt.Println("read the message send by client")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read server side error", err)
		return
	}
	// fmt.Println("conn.Read len buf", this.Buf[:4])

	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkglen])

	if n != int(pkglen) || err != nil {
		return
	}
	err = json.Unmarshal(this.Buf[:pkglen], &msg)
	if err != nil {
		fmt.Println("error of json unmarshal,", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// send data length to client.
	var pkglen uint32
	pkglen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkglen)
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("error of conn.Write(buf[0:4]),", err)
		return
	}
	// send data to client
	n, err = this.Conn.Write(data)
	if n != int(pkglen) && err != nil {
		fmt.Println("error of conn.Write(data),", err)
		return
	}
	return
}

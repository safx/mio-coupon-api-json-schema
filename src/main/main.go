package main

import (
	"iijmio"
	"fmt"
)

func main() {
  mio := iijmio.NewService(nil)

  cs, err := mio.CouponList(nil)
  fmt.Println(cs, err)

  ps, err := mio.PacketList(nil)
  fmt.Println(ps, err)
}


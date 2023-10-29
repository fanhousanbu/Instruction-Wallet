package infra

import (
	"github.com/tarm/serial"
	"time"
)

// SmsGatewayConnector 短信网关连接器
type SmsGatewayConnector struct {
	port *serial.Port
}

// Listen 打开并侦听串口
func (c *SmsGatewayConnector) Listen(name string, baud int, reader func([]byte)) error {
	s := &serial.Config{Name: name, Baud: baud, ReadTimeout: time.Second * 5}
	if p, err := serial.OpenPort(s); err != nil {
		return err
	} else {
		c.port = p
	}

	if reader != nil {
		go readFromSmsGatewayConnector(c.port, reader)
	}

	return nil
}

// Write 往串口发送指令
func (c *SmsGatewayConnector) Write(msg []byte) (int, error) {
	return c.port.Write(msg)
}

func readFromSmsGatewayConnector(port *serial.Port, reader func([]byte)) {
	for {
		buf := make([]byte, 10240)
		if n, e := port.Read(buf); e == nil && n > 0 {
			reader(buf[:n])
		}
	}
}

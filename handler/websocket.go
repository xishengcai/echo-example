package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
)

var upgrader = websocket.Upgrader{} // use default options

func Echo(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Print("upgrade:", err)
		return err
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	return nil
}

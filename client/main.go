package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	// Conectar al servidor RPC
	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// Crear los items
	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	// Agregar los items a la base de datos
	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)

	// Ver la base de datos antes de la eliminación
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database :", db)

	// Eliminar los items
	client.Call("API.DeleteItem", b, &reply)
	client.Call("API.DeleteItem", a, &reply)

	// Ver la base de datos después de la eliminación
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database after deletion:", db)

	// Ver la base de datos después de eliminar el último item
	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Final Database:", db)
}

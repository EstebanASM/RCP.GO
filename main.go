package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

type API int

var database []Item

// Obtener todos los elementos de la base de datos
func (a *API) GetDB(empty string, reply *[]Item) error {
	*reply = database
	return nil
}

// Obtener un elemento por su título
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}

	*reply = getItem
	return nil
}

// Agregar un nuevo elemento a la base de datos
func (a *API) AddItem(item Item, reply *Item) error {
	// Asegurarse de que no haya duplicados antes de agregar
	for _, val := range database {
		if val.Title == item.Title {
			return nil // No agregar si ya existe un elemento con el mismo título
		}
	}
	database = append(database, item)
	*reply = item
	return nil
}

// Editar un elemento existente en la base de datos
func (a *API) EditItem(item Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == item.Title {
			database[idx] = Item{item.Title, item.Body} // Actualizar el elemento
			changed = database[idx]
			break
		}
	}

	*reply = changed
	return nil
}

// Eliminar un elemento de la base de datos
func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	var updatedDatabase []Item

	// Filtrar los elementos que no coinciden con el item a eliminar
	for _, val := range database {
		if val.Title != item.Title || val.Body != item.Body {
			updatedDatabase = append(updatedDatabase, val)
		} else {
			del = val
		}
	}

	// Reemplazar la base de datos con el nuevo slice filtrado
	database = updatedDatabase
	*reply = del
	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}
}

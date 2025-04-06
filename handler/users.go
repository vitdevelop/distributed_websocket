package handler

import (
	"net"
	"os"
	"sync"
)

type User struct {
	Id       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Image    string `json:"image,omitempty"`
	Instance string `json:"instance"`
}

type UserSession struct {
	user User
	conn net.Conn
}

var names []User
var users = make(map[uint]UserSession)
var mutex sync.Mutex
var instanceName string

func init() {
	name := os.Getenv("INSTANCE_NAME")
	if name == "" {
		name = "Helicarrier"
	}
	instanceName = name

	names = append(names, User{
		Id:       1,
		Name:     "Black Widow",
		Image:    "/heroes/black_widow.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       2,
		Name:     "Captain America",
		Image:    "/heroes/captain_america.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       3,
		Name:     "Cyclop",
		Image:    "/heroes/cyclop.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       4,
		Name:     "Daredevil",
		Image:    "/heroes/daredevil.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       5,
		Name:     "Deadpool",
		Image:    "/heroes/deadpool.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       6,
		Name:     "Doctor Strange",
		Image:    "/heroes/doctor_strange.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       7,
		Name:     "Drax",
		Image:    "/heroes/drax.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       8,
		Name:     "Gamora",
		Image:    "/heroes/gamora.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       9,
		Name:     "Ghost Rider",
		Image:    "/heroes/ghost_rider.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       10,
		Name:     "Groot",
		Image:    "/heroes/groot.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       11,
		Name:     "Hawkeye",
		Image:    "/heroes/hawkeye.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       12,
		Name:     "Heimdall",
		Image:    "/heroes/heimdall.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       13,
		Name:     "Hulk",
		Image:    "/heroes/hulk.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       14,
		Name:     "Iceman",
		Image:    "/heroes/iceman.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       15,
		Name:     "Iron Man",
		Image:    "/heroes/iron_man.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       16,
		Name:     "Jean Grey",
		Image:    "/heroes/jean_grey.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       17,
		Name:     "Nick Fury",
		Image:    "/heroes/nick_fury.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       18,
		Name:     "Professor X",
		Image:    "/heroes/professor_x.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       19,
		Name:     "Rocket",
		Image:    "/heroes/rocket.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       20,
		Name:     "Rogue",
		Image:    "/heroes/rogue.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       21,
		Name:     "Spiderman",
		Image:    "/heroes/spiderman.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       22,
		Name:     "Star Lord",
		Image:    "/heroes/star_lord.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       23,
		Name:     "Storm",
		Image:    "/heroes/storm.png",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       24,
		Name:     "Thor",
		Image:    "/heroes/thor.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       25,
		Name:     "Vision",
		Image:    "/heroes/vision.jpg",
		Instance: instanceName,
	})
	names = append(names, User{
		Id:       26,
		Name:     "Wolverine",
		Image:    "/heroes/wolverine.jpg",
		Instance: instanceName,
	})
}

func ConnectAvailableUser(conn net.Conn) User {
	mutex.Lock()
	defer mutex.Unlock()

	var user User
	if len(names) == 0 {
		user = User{
			Id:    uint(len(users) + 1),
			Name:  "Mistique",
			Image: "/heroes/mystique.jpg",
		}
	} else {
		user = names[len(names)-1]
		names = names[:len(names)-1]
	}

	users[user.Id] = UserSession{
		user: user,
		conn: conn,
	}

	return user
}

func ReturnAvailableUser(user User) {
	mutex.Lock()
	defer mutex.Unlock()

	names = append(names, user)
	delete(users, user.Id)
}

func GetConnectedUsers() []User {
	mutex.Lock()
	defer mutex.Unlock()

	connectedUsers := make([]User, 0, len(users))
	for id := range users {
		connectedUsers = append(connectedUsers, users[id].user)
	}

	return connectedUsers
}

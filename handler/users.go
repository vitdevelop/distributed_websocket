package handler

import (
	"net"
	"sync"
)

type User struct {
	Id    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

type UserSession struct {
	user User
	conn net.Conn
}

var names []User
var users = make(map[uint]UserSession)
var mutex sync.Mutex

func init() {
	names = append(names, User{
		Id:    1,
		Name:  "Black Widow",
		Image: "/heroes/black_widow.jpg",
	})
	names = append(names, User{
		Id:    2,
		Name:  "Captain America",
		Image: "/heroes/captain_america.jpg",
	})
	names = append(names, User{
		Id:    3,
		Name:  "Cyclop",
		Image: "/heroes/cyclop.jpg",
	})
	names = append(names, User{
		Id:    4,
		Name:  "Daredevil",
		Image: "/heroes/daredevil.png",
	})
	names = append(names, User{
		Id:    5,
		Name:  "Deadpool",
		Image: "/heroes/deadpool.png",
	})
	names = append(names, User{
		Id:    6,
		Name:  "Doctor Strange",
		Image: "/heroes/doctor_strange.jpg",
	})
	names = append(names, User{
		Id:    7,
		Name:  "Drax",
		Image: "/heroes/drax.png",
	})
	names = append(names, User{
		Id:    8,
		Name:  "Gamora",
		Image: "/heroes/gamora.jpg",
	})
	names = append(names, User{
		Id:    9,
		Name:  "Ghost Rider",
		Image: "/heroes/ghost_rider.jpg",
	})
	names = append(names, User{
		Id:    10,
		Name:  "Groot",
		Image: "/heroes/groot.png",
	})
	names = append(names, User{
		Id:    11,
		Name:  "Hawkeye",
		Image: "/heroes/hawkeye.png",
	})
	names = append(names, User{
		Id:    12,
		Name:  "Heimdall",
		Image: "/heroes/heimdall.jpg",
	})
	names = append(names, User{
		Id:    13,
		Name:  "Hulk",
		Image: "/heroes/hulk.png",
	})
	names = append(names, User{
		Id:    14,
		Name:  "Iceman",
		Image: "/heroes/iceman.jpg",
	})
	names = append(names, User{
		Id:    15,
		Name:  "Iron Man",
		Image: "/heroes/iron_man.jpg",
	})
	names = append(names, User{
		Id:    16,
		Name:  "Jean Grey",
		Image: "/heroes/jean_grey.jpg",
	})
	names = append(names, User{
		Id:    17,
		Name:  "Nick Fury",
		Image: "/heroes/nick_fury.png",
	})
	names = append(names, User{
		Id:    18,
		Name:  "Professor X",
		Image: "/heroes/professor_x.jpg",
	})
	names = append(names, User{
		Id:    19,
		Name:  "Rocket",
		Image: "/heroes/rocket.png",
	})
	names = append(names, User{
		Id:    20,
		Name:  "Rogue",
		Image: "/heroes/rogue.png",
	})
	names = append(names, User{
		Id:    21,
		Name:  "Spiderman",
		Image: "/heroes/spiderman.png",
	})
	names = append(names, User{
		Id:    22,
		Name:  "Star Lord",
		Image: "/heroes/star_lord.jpg",
	})
	names = append(names, User{
		Id:    23,
		Name:  "Storm",
		Image: "/heroes/storm.png",
	})
	names = append(names, User{
		Id:    24,
		Name:  "Thor",
		Image: "/heroes/thor.jpg",
	})
	names = append(names, User{
		Id:    25,
		Name:  "Vision",
		Image: "/heroes/vision.jpg",
	})
	names = append(names, User{
		Id:    26,
		Name:  "Wolverine",
		Image: "/heroes/wolverine.jpg",
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

func broadcastUserMessage(user User, message WsMessage) {
	for _, userSession := range users {
		if userSession.user.Id == user.Id {
			continue
		}

		conn := userSession.conn

		sendWsMessage(conn, message)
	}
}

package kenopsialobby

type Lobby struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Players   []Player `json:"players"`
	Bots      []Bot    `json:"bots"`
	CreatorId string   `json:"creatorId"`
	CreatedAt int64    `json:"createdAt"`
}

type Bot struct {
	Id       uint8  `json:"id"`
	Username string `json:"username"`
	AvatarId uint8  `json:"avatarId"`
}

type Player struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	AvatarId uint8  `json:"avatarId"`
}

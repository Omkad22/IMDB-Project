// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type DirectorInput struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
	BornAt       string `json:"bornAt"`
	Description  string `json:"description"`
}

type FavoriteInput struct {
	MovieID    int    `json:"movieID"`
	UserID     int    `json:"userID"`
	MovieTitle string `json:"movieTitle"`
}

type MovieInput struct {
	Description string `json:"description"`
	Title       string `json:"title"`
	Rank        int    `json:"rank"`
	Genre       string `json:"genre"`
	DirectorID  int    `json:"director_id"`
	Image       string `json:"image"`
	Topic       string `json:"topic"`
	Text        string `json:"text"`
	Year        int    `json:"year"`
}

type ReviewInput struct {
	Topic   string `json:"topic"`
	Text    string `json:"text"`
	Rank    int    `json:"rank"`
	MovieID int    `json:"movieID"`
}

type UserInput struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	Password    string `json:"password"`
	Profile     string `json:"profile"`
	Email       string `json:"email"`
	Birthday    string `json:"birthday"`
}

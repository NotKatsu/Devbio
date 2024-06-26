package structs

import "time"

type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type HashedAndSaltedPassword struct {
	HashedPassword string
	RandomSalt     string
}

type SessionCreated struct {
	SessionAuthentication string `json:"session_authentication"`
}

type UserResponse struct {
	Username        string   `json:"username"`
	ProfilePicture  []byte   `json:"profile_picture"`
	Description     string   `json:"description"`
	Skills          []string `json:"skills"`
	Interests       []string `json:"interests"`
	Location        string   `json:"location"`
	SpokenLanguages []string `json:"spoken_languages"`
	Badges          []string `json:"badges"`
	IsSetup         bool     `json:"is_setup"`
	IsHirable       bool     `json:"is_hirable"`
	IsDisabled      bool     `json:"is_disabled"`
	Colour          int64    `json:"selected_colour"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type ExploreData struct {
	Rank            float64 `json:"rank"`
	Username        string  `json:"username"`
	AvgRating       float32 `json:"avg_rating"`
	YearsExperience int     `json:"years_experience"`
	Commits         int     `json:"commits"`
	OpenProjects    int     `json:"open_projects"`
	Boosts          int     `json:"boosts"`
}

type ExploreDatsResponse struct {
	ExploreData []ExploreData `json:"explore_data"`
}

type ExploreResponse struct {
	ExploreData []ExploreData
}

type NotificationsResponse struct {
	Notifications []string `json:"notifications"`
}

type Connection struct {
	IsShown         bool      `json:"is_shown"`
	Username        string    `json:"username"`
	AccountUsername string    `json:"account_username"`
	ConnectionType  string    `json:"connection_type"`
	ConnectionDate  time.Time `json:"connection_date"`
}

type ConnectionsResponse struct {
	Connections []Connection
}

type Statistics struct {
	Username           string `json:"username"`
	ProfileViews       string `json:"profile_views"`
	ConnectionsClicked string `json:"connections_clicked"`
}

type StatisticsResponse struct {
	Statistics []Statistics `json:"statistics"`
}

type RepositoryResponse struct {
	RepositoryName        string `json:"repository_name"`
	RepositoryDescription string `json:"repository_description"`
	RepositoryURL         string `json:"repository_url"`
	StarCount             int    `json:"star_count"`
	Language              string `json:"language"`
	IsShown               bool   `json:"is_shown"`
}

type RepositoriesResponse struct {
	Repositories []RepositoryResponse `json:"github_repositories"`
}

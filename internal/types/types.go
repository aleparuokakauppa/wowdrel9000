package types

type Guess struct {
    Version int `json:"version"`
    Guess string `json:"guess"`
}

type Letter struct {
    Char rune `json:"char"`
    Status string `json:"status"`
}

type GuessResponse struct {
    Version int `json:"version"`
    Win bool `json:"win"`
    Letters [5]Letter `json:"letters"`
}

type RealWordResponse struct {
    Version int `json:"version"`
    IsReal bool `json:"isreal"`
}

type ServerConfigResponse struct {
    Version int `json:"version"`
    ServerEndPoint string `json:"serverendpoint"`
}

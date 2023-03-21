package activitypub

type Activity struct {
    Type      string `json:"type"`
    Actor     string `json:"actor"`
    Object    string `json:"object"`
    Published string `json:"published"`
}

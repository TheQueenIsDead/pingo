package models

type Target struct {
	Id        int    `db:"primarykey, autoincrement"`
	Type      string `json:"type" db:"type"`
	Source    string `json:"source" db:"source"`
	Frequency int    `json:"frequency" db:"frequency"`
	Unit      string `json:"unit" db:"unit"`
}

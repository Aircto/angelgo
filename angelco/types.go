package angelco

import (
	"time"
)

type User struct {
	Name          string
	Bio           string
	BlogUrl       string `json:"blog_url"`
	WhatIDo       string `json:"what_i_do"`
	WhatIveBuilt  string `json:"what_ive_built"`
	OnlineBioUrl  string `json:"online_bio_url"`
	GithubUrl     string `json:"github_url"`
	ResumeUrl     string `json:"resume_url"`
	TwitterUrl    string `json:"twitter_url"`
	FacebookUrl   string `json:"facebook_url"`
	LinkedinUrl   string `json:"linkedin_url"`
	AngelListUrl  string `json:"angellist_url"`
	AboutMeUrl    string `json:"aboutme_url"`
	BehanceUrl    string `json:"behance_url"`
	DribbbleUrl   string `json:"dribbble_url"`
	Criteria      string
	Image         string
	Email         string
	Id            int64
	FollowerCount int64 `json:"follower_count"`
	Investor      bool
	Locations     []Location
	Roles         []Role
	Skills        []Skill
}

type CommonType struct {
	Id           int64
	TagType      string `json:"tag_type"`
	Name         string
	DisplayName  string `json:"display_name"`
	AngelListUrl string `json:"angelist_url"`
}

type Location CommonType

type Role CommonType

type Skill struct {
	CommonType
	Level string
}

type StartupRole struct {
	Confirmed bool
	CreatedAt time.Time `json:"created_at"`
	StartedAt string    `json:"started_at"`
	EndedAt   string    `json:"ended_at"`
	Role      string
	Id        int64
	Startup   Startup `json:"startup"`
	Tagged    Tag     `json:"tagged"`
}

type Startup struct {
	Hidden           bool
	CommunityProfile bool      `json:"community_profile"`
	CompanySize      string    `json:"company_size"`
	AngelListUrl     string    `json:"angellist_url"`
	CompanyUrl       string    `json:"company_url"`
	CreatedAt        time.Time `json:"created_at"`
	HighConcept      string    `json:"high_concept"`
	LogoUrl          string    `json:"logo_url"`
	Name             string
	ProductDesc      string    `json:"product_desc"`
	ThumbUrl         string    `json:"thumb_url"`
	CrunchbaseUrl    string    `json:"crunchbase_url"`
	TwitterUrl       string    `json:"twitter_url"`
	BlogUrl          string    `json:"blog_url"`
	FacebookUrl      string    `json:"facebook_url"`
	LinkedinUrl      string    `json:"linkedin_url"`
	VideoUrl         string    `json:"video_url"`
	UpdatedAt        time.Time `json:"updated_at"`
	Quality          int64
	Id               int64
	FollowerCount    int64 `json:"follower_count"`
	Markets          []Tag
	Locations        []Tag
	CompanyType      []Tag `json:"company_type"`
	Status           StatusUpdate
	Screenshots      []string
}

type Tag struct {
	AngelListUrl  string `json:"angellist_url"`
	Bio           string
	Image         string
	Name          string
	Type          string
	Title         string
	FollowerCount int64 `json:"follower_count"`
	Id            int64
}

type Job struct {
	Id           int64
	SalaryMin    int64     `json:"salary_min"`
	SalaryMax    int64     `json:"salary_max"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Title        string
	Description  string
	EquityCliff  string `json:"equity_cliff"`
	EquityMin    string `json:"equity_min"`
	EquityMax    string `json:"equity_max"`
	EquityVest   string `json:"equity_vest"`
	JobType      string `json:"job_type"`
	AngelListUrl string `json:"angellist_url"`
	CurrencyCode string `json:"currency_code"`
	RemoteOk     bool   `json:"remote_ok"`
	Startup      *Startup
	Tags         []Tag
}

type AccessToken struct {
	Key  string `json:"access_token"`
	Type string `json:"token_type"`
}

type StatusUpdate struct {
	Id        int64
	Message   string
	CreatedAt time.Time `json:"created_at"`
}

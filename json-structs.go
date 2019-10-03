package main

type BitbucketLinkHref struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

type BitbucketFromRefProjectLink struct {
	Self []BitbucketLinkHref `json:"self"`
}

type BitbucketRefProject struct {
	Key         string `json:"key"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Type        string `json:"type"`
}

type BitbucketObjectLinks struct {
	Clone []BitbucketLinkHref `json:"clone"`
	Self  []BitbucketLinkHref `json:"self"`
}

type BitbucketRefRepository struct {
	Slug          string               `json:"slug"`
	Id            int                  `json:"id"`
	Name          string               `json:"name"`
	ScmId         string               `json:"scmId"`
	State         string               `json:"state"`
	StatusMessage string               `json:"statusMessage"`
	Forkable      bool                 `json:"forkable"`
	Project       BitbucketRefProject  `json:"project"`
	Links         BitbucketObjectLinks `json:"links"`
}

type BitbucketRef struct {
	Id           string                 `json:"id"`
	DisplayId    string                 `json:"displayId"`
	LatestCommit string                 `json:"latestCommit"`
	Repository   BitbucketRefRepository `json:"repository"`
}

type BitbucketUser struct {
	Name         string               `json:"name"`
	EmailAddress string               `json:"emailAddress"`
	Id           int                  `json:"id"`
	DisplayName  string               `json:"displayName"`
	Active       bool                 `json:"active"`
	Slug         string               `json:"slug"`
	Type         string               `json:"type"`
	Links        BitbucketObjectLinks `json:"links"`
}

type BitbucketUserInRole struct {
	User     BitbucketUser `json:"user"`
	Role     string        `json:"role"`
	Approved bool          `json:"approved"`
	Status   string        `json:"status"`
}

type BitbucketPropertiesMergeResult struct {
	Outcome string `json:"outcome"`
	Current bool   `json:"current"`
}

type BitbucketProperties struct {
	MergeResult       BitbucketPropertiesMergeResult `json:"mergeResult"`
	ResolvedTaskCount int                            `json:"resolvedTaskCount"`
	OpenTaskCount     int                            `json:"openTaskCount"`
}

type BitbucketIndexValue struct {
	Id          int                   `json:"id"`
	Version     int                   `json:"version"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	State       string                `json:"state"`
	Open        bool                  `json:"open"`
	Closed      bool                  `json:"closed"`
	CreatedDate int64                 `json:"createdDate"`
	UpdatedDate int64                 `json:"updatedDate"`
	FromRef     BitbucketRef          `json:"fromRef"`
	ToRef       BitbucketRef          `json:"toRef"`
	Locked      bool                  `json:"locked"`
	Author      BitbucketUserInRole   `json:"author"`
	Reviewers   []BitbucketUserInRole `json:"reviewers"`
	// Participants []?
	Properties BitbucketProperties  `json:"properties"`
	Links      BitbucketObjectLinks `json:"links"`
}

type BitbucketPullRequestIndex struct {
	Size          int                   `json:"size"`
	Limit         int                   `json:"limit"`
	IsLastPage    bool                  `json:"isLastPage"`
	Values        []BitbucketIndexValue `json:"values"`
	Start         int                   `json:"start"`
	NextPageStart int                   `json:"nextPageStart"`
}

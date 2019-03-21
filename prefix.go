package seo

import "github.com/ecletus/helpers"

var PREFIX string

func init() {
	PREFIX = helpers.GetCalledDir()
}

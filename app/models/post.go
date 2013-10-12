package models

import "github.com/robfig/revel"

type Post struct {
  Id int64
  Title string
  Body string
  Active bool
}

func (post *Post) Validate(v *revel.Validation) {
  v.Check(
    post.Title,
    revel.Required{},
    revel.MinSize{3},
    revel.MaxSize{256},
  )

  v.Check(
    post.Body,
    revel.Required{},
    revel.MinSize{10},
    revel.MaxSize{2048},
  )
}

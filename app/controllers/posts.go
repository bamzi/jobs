package controllers

import(
  "github.com/robfig/revel"
  "jobs/app/models"
)

type Posts struct {
	*revel.Controller
}

func (c Posts) Index() revel.Result {
  results, err := Db.Select(
    models.Post{},
    `select * from posts where active = ?`, true)
  if err != nil {
    panic(err)
  }

  var posts []*models.Post
  for _, r := range results {
    post := r.(*models.Post)
    posts = append(posts, post)
  }

  return c.Render(posts)
}

func (c Posts) New() revel.Result {
	return c.Render()
}

func (c Posts) Create() revel.Result {
	return c.Render()
}

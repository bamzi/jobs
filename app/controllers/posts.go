package controllers

import(
  "github.com/robfig/revel"
  "github.com/gylaz/jobs/app/models"
  "github.com/gylaz/jobs/app/routes"
)

type Posts struct {
	*revel.Controller
}

func (c Posts) Index() revel.Result {
  results, err := Db.Select(
    models.Post{},
    `select * from posts where active = $1`, true)
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

func (c Posts) Create(post models.Post) revel.Result {
  post.Active = true
  err := Db.Insert(&post)
  if err != nil {
    panic(err)
  }

  return c.Redirect(routes.Posts.Index())
}

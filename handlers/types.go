package handlers


type RegisterInput struct{
  Username   string  `json:"username" binding:"required,min=3"`
  Email      string  `json:"email" binding:"required,email"`
  Password   string  `json:"password" binding:"required,min=6"`
}

type LoginInput struct{
  Email      string  `json:"email" binding:"required,email"`
  Password   string  `json:"password" binding:"required,min=6"`
}
package message

const (
  Type_Login = "LoginMessage"
  Type_LoginResult = "LoginResultMessage"
)

type Message struct {
  Type string `json:"type"`
  Data string `json:"data"`
}

type LoginMessage struct {
  UserId int `json:"user_id"`
  UserPwd string `json:"user_pwd"`
  UserName string `json:"user_name"`
}

type LoginResultMessage struct {
  Code int `json:"code"`
  Error string `json:"error"`
}

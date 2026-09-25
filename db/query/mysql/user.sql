-- name: CreateUser :execresult
-- description: 创建用户
-- args: {id: string, avatar: string, nickname: string, password: string, phone: string, sex: int}
-- returns: {id: string}
INSERT INTO users(
   id,    avatar,nickname,password,phone,sex
) VALUES( 
    ?, ?,  ?, ?, ?, ?
);
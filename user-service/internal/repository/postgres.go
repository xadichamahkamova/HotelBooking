package repository

import (
	"database/sql"
	pb "user-service/genproto/userpb"
	"user-service/logger"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepo{
		DB: db,
	}
}

func (db *UserRepo) RegisterUser(req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

	logger.Info("RegisterUser called with request:", req)

	resp := pb.RegisterUserResponse{}
	query := `
	INSERT INTO users(user_name, password, email) 
	VALUES($1, $2, $3)
	RETURNING id, user_name, email`
	err := db.DB.QueryRow(query,
		req.UserName,
		req.Password,
		req.Email,
	).Scan(
		&resp.UserId,
		&resp.UserName,
		&resp.Email,
	)
	if err != nil {
		logger.Error("Error inserting user:", err)
		return nil, err
	}
	logger.Info("User registered successfully with ID:", resp.UserId)
	return &resp, nil
}

func (db *UserRepo) LoginUser(req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	logger.Info("LoginUser called with request:", req)

	resp := pb.LoginUserResponse{}
	query := `
	SELECT id, email
	FROM users 
	WHERE user_name=$1 AND password=$2`
	err := db.DB.QueryRow(query,
		req.UserName,
		req.Password,
	).Scan(
		&resp.UserId,
		&resp.UserEmail,
	)
	if err != nil {
		logger.Error("Error logging in user:", err)
		return nil, err
	}
	logger.Info("User logged in successfully with ID:", resp.UserId)
	return &resp, nil
}

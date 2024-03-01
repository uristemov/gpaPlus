package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/uristemov/repeatPro/api"
	"github.com/uristemov/repeatPro/internal/entity"
	"github.com/uristemov/repeatPro/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) (string, error) {

	hashPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return "", fmt.Errorf("hash password error %w", err)
	}

	u.Password = hashPassword

	return m.Repository.CreateUser(ctx, u)
}

func (m *Manager) Login(ctx context.Context, email, password string) (string, string, error) {

	user, err := m.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", fmt.Errorf("user not found")
		}

		return "", "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreateAccessToken(user.Id.String(), user.Email, m.Config.Auth.Access.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create access token err: %w", err)
	}

	refreshToken, err := m.Token.CreateRefreshToken(user.Id.String(), m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create refresh token err: %w", err)
	}

	err = m.Cache.TokenCache.SetToken(ctx, user.Id.String(), refreshToken, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (m *Manager) VerifyToken(token string) (string, error) {

	claim, err := m.Token.ValidateAccessToken(token)
	if err != nil {
		return "", fmt.Errorf("validate token err: %w", err)
	}

	return claim.UserID, nil
}

func (m *Manager) Refresh(ctx context.Context, refreshToken string) (string, string, error) {

	claim, err := m.Token.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("validate token err: %w", err)
	}

	token, err := m.Cache.TokenCache.GetToken(ctx, claim.UserID)
	if err != nil {
		return "", "", err
	}

	if token == "" {
		return "", "", fmt.Errorf("old refresh token from cache is empty error")
	}

	// TODO check user with user_id for existence
	user, err := m.GetUserById(ctx, claim.UserID)
	if err != nil {
		m.logger.Errorf("get user from refresh token err: %w", err)
		return "", "", err
	}

	accessToken, err := m.Token.CreateAccessToken(claim.UserID, user.Email, m.Config.Auth.Access.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create access token err: %w", err)
	}

	newRefreshToken, err := m.Token.CreateRefreshToken(claim.UserID, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", fmt.Errorf("create refresh token err: %w", err)
	}

	err = m.Cache.TokenCache.SetToken(ctx, claim.UserID, newRefreshToken, m.Config.Auth.Refresh.TimeToLive)
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func (m *Manager) UpdateUser(ctx context.Context, id string, req *api.UpdateUserRequest) error {

	user, err := m.Repository.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Password != "" {
		req.Password, err = util.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user.Password = req.Password
	}
	if req.ImagePath != "" {
		user.ImagePath = req.ImagePath
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.UniversityId != "" {
		user.UniversityId = req.UniversityId
	}
	if req.RoleId != 0 {
		user.RoleId = req.RoleId
	}

	err = m.Cache.UserCache.DeleteUser(ctx, user.Id.String())
	if err != nil {
		return err
	}

	err = m.Repository.UpdateUser(ctx, id, req)
	if err != nil {
		return err
	}

	_ = m.Cache.UserCache.SetUser(ctx, user)

	return nil
}

func (m *Manager) UpgradeUser(ctx context.Context, id string, req *api.UpdateUserRequest) error {

	user, err := m.Repository.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	if req.UniversityId != "" {
		user.UniversityId = req.UniversityId
	}

	user.Verified = req.Verified

	if req.RoleId != 0 {
		user.RoleId = req.RoleId
	}

	err = m.Cache.UserCache.DeleteUser(ctx, user.Id.String())
	if err != nil {
		return err
	}

	err = m.Repository.UpgradeUser(ctx, id, req)
	if err != nil {
		return err
	}

	_ = m.Cache.UserCache.SetUser(ctx, user)

	return nil
}

func (m *Manager) GetUserById(ctx context.Context, id string) (*entity.User, error) {

	user, err := m.Cache.UserCache.GetUser(ctx, id)
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = m.Repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	user, err := m.Repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = m.Cache.UserCache.SetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) DeleteUser(ctx context.Context, id string) error {

	err := m.Cache.UserCache.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	err = m.Cache.TokenCache.DeleteToken(ctx, id)
	if err != nil {
		m.logger.Infof("delete refresh token when delete user err: %w", err)
	}

	return m.Repository.DeleteUser(ctx, id)
}

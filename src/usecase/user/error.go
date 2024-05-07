package userusecase

import "errors"

var ErrUserNotFound = errors.New("User tidak ditemukan")
var ErrEmailAlreadyUsed = errors.New("Email sudah digunakan")
var ErrInvalidPassword = errors.New("Password salah")
var ErrInvalidToken = errors.New("Token tidak valid")
var ErrExpiredToken = errors.New("Token kadaluwarsa")
var ErrTokenNotFound = errors.New("Token tidak ditemukan")
var ErrInvalidUser = errors.New("Email atau password salah")

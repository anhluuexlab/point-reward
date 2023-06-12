package utils

import "errors"

type Error struct {
	Message string
}

var (
	UserConflict          = errors.New("Người dùng đã tồn tại")
	UserNotFound          = errors.New("Người dùng không tồn tại")
	UserNotUpdated        = errors.New("Cập nhật thông tin người dùng thất bại")
	SignUpFail            = errors.New("Đăng ký thất bại")
	BalanceNotEnough      = errors.New("Tài khoản của bạn không đủ")
	AmountGreatThan10     = errors.New("Số tiền phải lớn hơn 10")
	TransactionNotFound   = errors.New("Giao dịch không tồn tại")
	TransactionIsRefunded = errors.New("Giao dịch hoàn tiền đã thực hiện trước đó")
)

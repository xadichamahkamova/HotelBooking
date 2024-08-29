package service

import (
	"context"
	pb "notification-service/genproto/notificationpb"
	"notification-service/internal/email"
)

type NotificationService struct {
	pb.UnimplementedNotificationServiceServer
	Repo email.NotificationRepo
}

func NewNotificationService(repo email.NotificationRepo) *NotificationService {
	return &NotificationService{
		Repo: repo,
	}
}

func (s *NotificationService) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	return s.Repo.SendEmail(req)
}
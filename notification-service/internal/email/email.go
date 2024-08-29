package email

import (
	"fmt"

	pb "notification-service/genproto/notificationpb"
	cfg "notification-service/internal/pkg/load"
	"notification-service/logger"

	"gopkg.in/gomail.v2"
)

type NotificationRepo struct {
	Config cfg.Config 
}

func NewNotificationRepo(cfg cfg.Config) *NotificationRepo {
	return &NotificationRepo{
		Config: cfg,
	}
}

func (n *NotificationRepo) SendEmail(req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	
	logger.Info("SendEmail function started")

	subject := "----Welcome buddy----"

	body := fmt.Sprintf(`
    <!DOCTYPE html>
      <html lang="uz">
      <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Xona holati haqida xabar</title>
        <style>
          body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
          }
          .container {
            width: 80%%;
            margin: 20px auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
          }
          .header {
            background-color: #4CAF50;
            color: white;
            text-align: center;
            padding: 10px;
            border-radius: 5px 5px 0 0;
          }
          .content {
            padding: 20px;
            text-align: center;
          }
          .message {
            font-size: 18px;
            font-weight: bold;
            margin-bottom: 20px;
          }
          .details {
            background-color: #f9f9f9;
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 10px;
            margin-bottom: 20px;
          }
          .success {
            color: #4CAF50;
          }
          .error {
            color: #f44336;
          }
          .info {
            color: #2196F3;
          }
          .footer {
            text-align: center;
            padding: 10px;
            font-size: 12px;
            color: #777;
          }
        </style>
      </head>
      <body>
        <div class="container">
          <div class="header">
            <h1>Xona holati haqida xabar</h1>
          </div>
          <div class="content">
            <p class="message">%s</p>
          </div>
          <div class="footer">
            <p>Bu avtomatik xabar. Iltimos, javob bermang.</p>
          </div>
        </div>
      </body>
      </html>
  `, req.HotelName)

	m := gomail.NewMessage()
	m.SetHeader("From", n.Config.Email.From)
	m.SetHeader("To", req.To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	logger.Info("Sending email to:", req.To)

	d := gomail.NewDialer("smtp.gmail.com", 587, n.Config.Email.From, n.Config.Email.Password)

	if err := d.DialAndSend(m); err != nil {
		logger.Error("Failed to send email:", err)
		return nil, err
	}

	logger.Info("Email sent successfully to:", req.To)

	return &pb.SendEmailResponse{
		Message: "Sended email succesfully",
	}, nil
}

package template

import (
	"html/template"
	"bytes"
)

func RenderOTPTemplate(otp string) (string, error) {
	const otpTemplate = `
	<!DOCTYPE html>
	<html>
	<head>
		<title>OTP Confirmation</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f7f7f7;
				margin: 0;
				padding: 0;
			}

			.container {
				background-color: #ffffff;
				max-width: 600px;
				margin: 0 auto;
				padding: 20px;
				border-radius: 5px;
				box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
			}

			h1 {
				color: #333;
			}

			p {
				font-size: 16px;
				line-height: 1.6;
				color: #555;
			}

			.notification {
				background-color: #f1f1f1;
				padding: 10px;
				margin-top: 20px;
			}

			.footer {
				background-color: #f1f1f1;
				padding: 10px;
				text-align: center;
			}

			.footer p {
				font-size: 14px;
				color: #777;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>OTP Confirmation</h1>
			<p>Hello,</p>
			<p>Your OTP for verification is: <strong>{{.OTP}}</strong></p>
			<p>If this was you, please use the OTP to complete your verification.</p>
			<p>If this wasn't you, please take immediate action to secure your account.</p>
			<div class="notification">
				<p>If you have any questions, please contact our support team.</p>
			</div>
		</div>
		<div class="footer">
			<p>Best regards, Reproduction Health Team</p>
		</div>
	</body>
	</html>
	`

	tmpl, err := template.New("otpTemplate").Parse(otpTemplate)
	if err != nil {
		return "", err
	}

	var emailBodyContent bytes.Buffer
	data := struct {
		OTP string
	}{
		OTP: otp,
	}

	err = tmpl.Execute(&emailBodyContent, data)
	if err != nil {
		return "", err
	}

	return emailBodyContent.String(), nil
}

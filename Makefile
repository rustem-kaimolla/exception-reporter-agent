build:
	docker build -f deployment/Dockerfile -t exception-reporter-agent .

up:
	docker run --env-file .env -p 9000:9000 exception-reporter-agent

test:
	echo '{"message":"Call to undefined method App\\\\Services\\\\PaymentService::process()", "file":"/var/www/app/Services/PaymentService.php", "line":87, "code":0, "trace":[{"file":"/var/www/app/Http/Controllers/PaymentController.php","line":45,"function":"process","class":"App\\\\Services\\\\PaymentService"}], "app":"billing-service", "env":"production", "timestamp":"2025-08-05T18:50:00Z"}' | nc 127.0.0.1 9000

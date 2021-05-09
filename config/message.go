package config

import "os"

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

var MailConfig EmailConfig

var AlertEmailList = []string{
	"pythonup@hotmail.com",
}

type Kafka struct {
	Broker []string
	Topic  map[string]string
}

var KafkaConfig Kafka

type RabbitMQ struct {
	Host  string
	Port  string
	User  string
	Pass  string
	Queue map[string]string
}

var RabbitMQConfig RabbitMQ

func init() {
	if os.Getenv("GIN_MODE") == "release" {
		KafkaConfig = Kafka{
			[]string{
				"127.0.0.1:9092",
			},
			map[string]string{
				"topic": "topic",
			},
		}
		RabbitMQConfig = RabbitMQ{
			"localhost",
			"5672",
			"guest",
			"guest",
			map[string]string{
				"queue": "queue",
			},
		}
	} else {
		KafkaConfig = Kafka{
			[]string{
				"127.0.0.1:9092",
			},
			map[string]string{
				"topic": "topic",
			},
		}
		RabbitMQConfig = RabbitMQ{
			"localhost",
			"5672",
			"guest",
			"guest",
			map[string]string{
				"queue": "queue",
			},
		}
	}

	MailConfig = EmailConfig{
		"smtp.163.com",
		465,
		"username",
		"password",
	}
}

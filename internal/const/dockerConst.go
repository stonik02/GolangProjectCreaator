package constants

const (
	Docker_header = `version: '3.8'
services:

`
	Docker_pg = `  postgres:
    image: postgres:latest
    restart: always
    environment:
      POSTGRES_DB: EQueue
      POSTGRES_PASSWORD: qwerty123
      POSTGRES_USER: denis_postgresql
    ports:
      - 5432:5432

`
	Docker_redis = `  redis:
    image: redis:latest
    command: redis-server
    volumes:
      - redis:/var/lib/redis
      - redis-config:/usr/local/etc/redis/redis.conf
    ports:
      - 6379:6379
    networks:
      - redis-network
volumes:
  redis:
  redis-config:

networks:
  redis-network:
    driver: bridge
`
)

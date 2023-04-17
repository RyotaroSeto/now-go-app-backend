DNS = host=nowdb user=postgres password=postgres dbname=nowdb port=5432 sslmode=disable TimeZone=Asia/Tokyo

gooseup:
	goose postgres "$(DNS)" up

.PHONY: gooseup

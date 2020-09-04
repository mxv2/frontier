CUR_DIR := $(shell pwd)
CONTAINER_NAME := "prometheus"

prometheus-up:
	@docker stop $(CONTAINER_NAME) || true
	@docker rm $(CONTAINER_NAME) || true
	@docker run \
        -p 9090:9090 \
        -v $(CUR_DIR)/prometheus.yml:/etc/prometheus/prometheus.yml \
        --name $(CONTAINER_NAME) \
        prom/prometheus
include ./config.mk

install:
	@if [ ! -d pact/bin ]; then\
		curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash;\
    fi

consumer-unit:
	go test ./consumer -run unit -count=1

consumer-pact:
	go test ./consumer -run pact -count=1

.PHONY: install consumer-unit consumer-pact
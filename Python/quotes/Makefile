.PHONY: help
help:
	@echo "Usage:"
	@echo "		virtualenv       Set up virtual environment"
	@echo "		activate         Activate virtual environment"
	@echo "		source           Source environment variable"
	@echo "		install          Install requirements.txt"
	@echo "		operate          Operate on local"

# Define the image name and tag
IMAGE_NAME = quotes
IMAGE_TAG = latest

.PHONY: virtualenv
virtualenv:
	virtualenv venv -p python3

.PHONY: activate
activate:
	. venv/bin/activate

.PHONY: source
source:
	source .env

.PHONY: install
install:
	pip3 install --trusted-host pypi.python.org -r requirements.txt

.PHONY: operate
operate:
	python3 main.py
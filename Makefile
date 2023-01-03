.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

# Adapted from https://www.thapaliya.com/en/writings/well-documented-makefiles/
.PHONY: help
help: ## Display this help.
help:
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-45s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

GO_FILES := $(wildcard *.go)

.PHONY: fmt
fmt: ## Format all code.
fmt:
	gofumpt -w $(GO_FILES)
	gci write --skip-generated -s standard -s default -s 'prefix(jdb.sh/garmin)' $(GO_FILES)

compose.yaml: ## Generate a Docker Compose file for Grafana and Postgresql.
compose.yaml: compose.jsonnet
	jsonnet $< -o $@

.PHONY: grafana-database
grafana-database: ## Create Grafana database.
grafana-database:
	echo "SELECT 'CREATE DATABASE grafana' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'grafana')\gexec" | PGPASSWORD=garmin psql -h localhost -U garmin

FIT_DIR   := fit

$(FIT_DIR):
	mkdir -p $@

$(FIT_DIR)/Primary/GARMIN/Activity: $(FIT_DIR)
	mkdir -p $@

MOUNT_DIR := /tmp/garmin

$(MOUNT_DIR):
	mkdir -p $@

$(MOUNT_DIR)/Primary: | $(MOUNT_DIR)
	sudo jmtpfs -o umask=0022,gid=100,uid=1000,allow_other $(MOUNT_DIR)
	touch $@

.PHONY: mount
mount: ## Mount the Garmin device to $(MOUNT_DIR).
mount: $(MOUNT_DIR)/Primary

.PHONY: umount
umount: ## Unmount the Garmin device from $(MOUNT_DIR).
umount:
	sudo umount $(MOUNT_DIR)

.PHONY: rsync
rsync: ## Rsync Garmin FIT files to $(FIT_DIR).
rsync: $(FIT_DIR)/Primary/GARMIN/Activity
	rsync -aRv \
		$(MOUNT_DIR)/./Primary/GARMIN/Activity \
		$(MOUNT_DIR)/./Primary/GARMIN/Monitor \
		$(MOUNT_DIR)/./Primary/GARMIN/Sleep \
		$(CURDIR)/$(FIT_DIR)

postgresql/query.sql.go: schema.sql query.sql
	sqlc generate

.PHONY: generate
generate: ## Generate Go code.
generate: postgresql/query.sql.go

.PHONY: compose
compose: ## Run container environment.
compose: compose.yaml
	podman-compose down
	podman-compose up -d

RELEASE := 21.94.00

FitSDKRelease_%.zip: ## Download a Fit SDK release.
FitSDKRelease_%.zip:
	curl -LO https://developer.garmin.com/downloads/fit/sdk/$@

Profile.xlsx: ## Extract the Fit SDK profile.
Profile.xlsx: FitSDKRelease_$(RELEASE).zip
	unzip -o $< $@
	touch $@

FitCSVTool.jar: ## Extract the FitCSVTool.
FitCSVTool.jar: FitSDKRelease_$(RELEASE).zip
	unzip -o -j $< java/FitCSVTool.jar
	touch $@

%.csv: ## Convert a Fit file into a CSV file.
%.csv: %.fit
	java -jar FitCSVTool.jar -b $< $@

%.fit: %.FIT
	mv $< $@

.PHONY: ingest
ingest: ## Ingest rsync'd data.
ingest: generate
	go run ./

vendor: ## Update vendored Go source code.
vendor: PersonalProfile.xlsx
	fitgen -hrst -verbose -sdk $(RELEASE)-Personal $< $@/github.com/tormoder/fit

Profile.csv.0 Profile.csv.1: ## Convert profile into CSV for editing.
Profile.csv.0 Profile.csv.1: Profile.xlsx
	ssconvert --export-file-per-sheet $< Profile.csv

PersonalProfile.xlsx: ## Convert modified personal profile into XLSX for code generation.
PersonalProfile.xlsx: Profile.csv.0 Profile.csv.1
	ssconvert --verbose --merge-to=PersonalProfile.xlsx $^

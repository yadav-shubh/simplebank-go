# Project Setup Guide

This guide provides steps to set up database migrations and SQL code generation for the project.

## Prerequisites

1. **Install Golang Migrate**  
   Golang Migrate is required to manage database migrations. Install it at the OS level.

   - **Installation Link**: [Golang Migrate - Installation Guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

2. **Install sqlc**  
   sqlc is a tool that generates type-safe database code from SQL files, allowing for seamless integration with Go code.

   - **macOS**: Install via Homebrew
     ```bash
     brew install sqlc
     ```
   - **Ubuntu**: Install via Snap
     ```bash
     sudo snap install sqlc
     ```
   - **init the sqlc**: initialize the sqlc in project
      ```bash
     sqlc init
     ```
   - **Run sqlc to generate go code**: Generate Go code from table using sqlc
      ```bash
     sqlc generate
     ```
      

## Steps

### 1. Generate Initial Schema Migration
To create an initial schema migration file, use the following command:

   ```bash
   migrate create -ext sql -dir db/migration -seq init_schema
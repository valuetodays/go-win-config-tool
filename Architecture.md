# Project Architecture

This document describes the directory structure, module responsibilities, and design principles for the Go + Wails project.

## 1. Overview

The project is a Wails-based desktop app.
We maintain clear separation between:

* Configuration (config)
* Core business logic (internal/domain, internal/service)
* Data transfer to UI (internal/dto)
* Utilities (internal/util)
* App orchestration (app.go and app_*.go)



## 2. Directory Structure

go-win-config-tool/
├── config/               # External configuration (YAML) handling
│   ├── root.go           # RootYAML definition
│   ├── software_config.go
│   ├── shortcut_config.go
│   └── normalize.go      # Normalize functions
├── internal/
│   ├── domain            # Runtime business models
│   │   ├── software_model.go
│   │   └── shortcut_model.go
│   │
│   ├── service           # Business logic services
│   │   ├── path_service.go
│   │   ├── env_service.go
│   │   ├── software_service.go
│   │   └── shortcut_service.go
│   │
│   ├── dto               # Data Transfer Objects for Vue / UI
│   │   ├── software_dto.go
│   │   ├── path_dto.go
│   │   ├── env_dto.go
│   │   └── shortcut_dto.go
│   │
│   └── util              # Internal utility functions
│       └── util.go       # e.g., ExpandWinEnv, path helpers
│
├── frontend/             # Wails frontend (Vue + Vite)
├── app.go                # App root, initializes services
├── app_paths.go          # Paths-related app functions
├── app_envs.go           # Envs-related app functions
├── app_softwares.go      # Softwares-related app functions
├── app_shortcuts.go      # Shortcuts-related app functions
├── app_reload.go         # Reload logic when switching config
└── main.go

## 3. Module Responsibilities

### 3.1 config

* Parse YAML configuration files.
* Provide normalized structures to services.
* Accessible globally (top-level, not internal), because it represents external input.

### 3.2 internal/domain

* Core runtime models.
* Should not be exposed to frontend directly.
* Immutable from UI perspective.

### 3.3 internal/dto

* Data structures returned to Vue frontend.
* Can include UI-specific fields (status, color, formatted strings).
* Mapped from domain models or config.
* Only for internal module consumption, hence inside internal.

### 3.4 internal/service

* Encapsulates business logic:

  * Path existence and creation
  * Environment variable checks
  * Software and shortcut management
* Returns DTOs to frontend.

### 3.5 internal/util

* Helper functions, reusable inside the project:

  * Path operations
  * Environment variable expansion (ExpandWinEnv)
  * File utilities
* Should not contain business logic.

### 3.6 app_*.go

* App orchestration for each module/tab.
* Calls service methods and maps results to DTOs for frontend.
* Keeps app.go clean by splitting per functional area.



## 4. Design Principles

1. Separation of Concerns

   * Config / Domain / DTO / Service / App / Frontend clearly separated.
2. Internal boundaries

   * Use internal to prevent external Go modules from importing project internals.
3. DTO vs Domain

   * Domain: core business logic
   * DTO: frontend representation
4. Config outside internal

   * config represents external input; not part of core logic.
5. Utilities

   * Reusable, stateless, pure functions in internal/util.
6. Environment variable handling

   * Always normalize paths via internal/util.ExpandWinEnv.
7. Extensibility

   * Tabs (Paths / Envs / Softwares / Shortcuts) can be added by creating new service + DTO + app_*.go without touching unrelated code.



## 5. Recommendations

* When adding new features, create:

  1. service for business logic
  2. domain models for runtime data
  3. dto for frontend output
  4. app_*.go to connect frontend and service
* Keep DTO strictly for frontend; domain should not include UI-specific fields.
* All utilities should be in internal/util to ensure reusability and clarity.
* Use internal for anything internal to the module, top-level for external-facing (like config).



## 6. Example: Expanding Software Base Dir

softwareBase := util.ExpandWinEnv(root.DefDirs.SoftwareDir)

* root.DefDirs.SoftwareDir comes from config YAML
* ExpandWinEnv is a pure helper in internal/util
* Ensures paths are normalized and environment variables expanded


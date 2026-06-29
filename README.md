# TaskPilot

TaskPilot bridges project management systems and AI coding assistants.

It synchronizes tasks from external providers (such as Jira) into Taskwarrior and launches AI assistants with task-aware context to help developers plan and execute work.

## Why TaskPilot?

Modern development workflows are fragmented:

* Tasks live in Jira, Linear, GitHub Issues, or other systems.
* Developers manage personal workflows in tools like Taskwarrior.
* AI assistants operate with little or no awareness of task context.

TaskPilot connects these systems together.

```text
Jira
  ↓
TaskPilot Sync
  ↓
Taskwarrior
  ↓
TaskPilot Work
  ↓
Claude / OpenAI / Ollama
```

The goal is to make task context portable, assistant-agnostic, and developer-friendly.

---

## Features

### Task Synchronization

Import tasks from external providers into Taskwarrior.

```bash
taskpilot sync
```

Current providers:

* Jira
* Fake provider (testing)

---

### AI Assistant Launching

Launch an assistant with task context.

```bash
taskpilot work claude JAT-1
```

TaskPilot:

1. Retrieves the task from Taskwarrior.
2. Builds assistant context.
3. Launches the selected assistant.
4. Starts the assistant in planning mode.

---

## Configuration

TaskPilot is configured using environment variables.

### Jira

```bash
export TASKPILOT_PROVIDER=jira

export JIRA_BASE_URL="https://company.atlassian.net"
export JIRA_EMAIL="developer@company.com"
export JIRA_API_TOKEN="your-api-token"

export JIRA_JQL='assignee = currentUser() AND status != Done'
```

### Example JQL Queries

My active work:

```text
assignee = currentUser() AND status != Done
```

High-priority issues:

```text
assignee = currentUser()
AND priority in (Highest, High)
AND status != Done
```

Current sprint:

```text
assignee = currentUser()
AND sprint in openSprints()
```

Backend tasks:

```text
project = API
AND labels = backend
AND status != Done
```

Bug fixes:

```text
assignee = currentUser()
AND issuetype = Bug
AND status != Done
```

TaskPilot does not restrict JQL usage. Any valid Jira query can be used.

### Fake Provider

For development and testing:

```bash
export TASKPILOT_PROVIDER=fake
```

---

### Assistant Configuration

TaskPilot does not install or configure AI assistants.

Assistants must already be available on your system.

Examples:

```bash
claude --version
```

```bash
ollama --version
```

TaskPilot simply discovers and launches configured assistants.

---

### Verify Configuration

Run:

```bash
taskpilot sync
```

If configuration is valid, tasks should be imported into Taskwarrior.

View imported tasks:

```bash
task +jira
```

### Provider Architecture

Providers are pluggable.

Current:

* Jira
* Fake Provider

Planned:

* GitHub Issues
* Linear
* Trello
* Custom Providers

---

### Assistant Architecture

Assistants are pluggable.

Current:

* Claude

Planned:

* OpenAI
* Ollama
* Additional local and remote assistants

---

## Installation

### Requirements

* Go
* Taskwarrior

### Build

```bash
go build -o taskpilot ./cmd/task
```

### Install

```bash
./scripts/install.sh
```

---

## Usage

### Synchronize Tasks

```bash
taskpilot sync
```

Imports tasks from the configured provider into Taskwarrior.

---


### Work On A Task

```bash
taskpilot work claude JAT-1
```

Launches Claude with context derived from task `JAT-1`.

---

## Project Structure

```text
internal/
├── app
├── core
├── dispatcher
├── logger
├── plugins
│   ├── assistants
│   ├── jira
│   └── taskwarrior
├── rules
└── shared
```

### Core

Domain models and interfaces.

### Providers

External task sources.

### Assistants

AI integrations.

### Rules

Filtering, transformation, and synchronization policies.

---

## Roadmap

### Current

* Jira → Taskwarrior synchronization
* Claude integration
* Planning mode support
* Modular provider architecture
* Modular assistant architecture

### Upcoming

* Rich Jira descriptions
* Attachment support
* Image handling
* Workspace-aware assistant launching
* OpenAI integration
* Ollama integration
* Additional providers

---

## Philosophy

TaskPilot is not an AI assistant.

TaskPilot is an orchestration layer.

Its responsibility is to:

* Retrieve task context.
* Normalize task data.
* Route context to assistants.
* Keep providers and assistants decoupled.

This allows developers to use their preferred task system and their preferred AI assistant without changing workflows.

---

## License

MIT


# CLAUDE Development Context for omnistrate-sdk-go

*Last Updated: 2025-09-30T02:00:19.269218*

## 🎯 Service Purpose & Context
omnistrate-sdk-go is a key microservice in the Omnistrate platform ecosystem.

## 📊 Repository Metrics
- **Total Commits**: 82
- **Active Branch**: main
- **Technologies**: Go, Make, GitHub Actions
- **Health Status**: healthy

## 🏗️ Architecture & Patterns

### Technology Stack
- **Go**: Primary development language with modern Go patterns
- **Make**: Build automation and development workflow
- **GitHub Actions**: Continuous integration and deployment

### File Organization
```
omnistrate-sdk-go/
├── cmd/                 # Entry points
├── pkg/                 # Main packages
├── config/              # Configuration
├── test/                # Tests
├── scripts/             # Build scripts
└── deploy/              # Deployment configs
```

## 🔄 Recent Development Activity

### Latest Commits (Last 30 Days)
No recent commits in the last 30 days.

### Active Development Areas
- None

## 👥 Development Team Context

### Contributors & Expertise
- **pberton**: 46 commits (Medium activity)
- **Yuhui**: 16 commits (Medium activity)
- **Alok**: 7 commits (Light activity)
- **dependabot[bot]**: 6 commits (Light activity)
- **Tomislav Simeunovic**: 4 commits (Light activity)
- **maziar kaveh**: 1 commits (Light activity)
- **Xinyi**: 1 commits (Light activity)
- **Alok Nikhil**: 1 commits (Light activity)

## 🛠️ Development Workflow

### Build & Test Pipeline
```bash
# Standard workflow for omnistrate-sdk-go
make tidy              # Clean dependencies
make build             # Build service
make test              # Run tests
make lint              # Code quality
make integration-test  # Integration tests (if available)
```

### Code Quality Metrics
- **Test Coverage**: 66 test files
- **Code Size**: 25.2 MB
- **Go Files**: 1421 files

## 🔍 Code Analysis Insights

### Health Assessment
**Overall Health**: Healthy

✅ Service follows best practices with no detected issues.

### Recommended Focus Areas
- Increase test coverage (current ratio suggests room for improvement)
- Consider regular maintenance and updates

## 🚀 AI Assistant Guidelines

When developing for omnistrate-sdk-go:

1. **Context Awareness**: Review recent commits to understand current development direction
2. **Pattern Consistency**: Follow established patterns visible in the codebase
3. **Integration Mindset**: Consider impact on other Omnistrate services
4. **Quality Standards**: Maintain high code quality and test coverage
5. **Documentation**: Keep documentation current with code changes

### Common Development Scenarios
- **Feature Addition**: Follow microservice patterns, add comprehensive tests
- **Bug Fixes**: Reproduce with tests, fix root cause, prevent regression
- **Refactoring**: Maintain API compatibility, update documentation
- **Performance**: Profile before optimizing, measure improvements

## 📈 Performance & Monitoring

### Key Metrics to Monitor
- Service response times
- Error rates and types
- Resource utilization
- Dependency health

### Observability Stack
- Logs: Structured logging with consistent fields
- Metrics: Prometheus-compatible metrics
- Traces: OpenTelemetry distributed tracing
- Health: HTTP health check endpoints

## 🔗 Integration Context

omnistrate-sdk-go integrates with the broader Omnistrate platform. Consider these integration points:
- Shared libraries from `commons/`
- API contracts defined in `api-design/`
- Orchestration patterns from service orchestration components
- Infrastructure dependencies and configurations

---
*This context is automatically maintained through MCP integration. Git data is refreshed daily.*

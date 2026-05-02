# Gridcore — Seeded Findings (24 total)

## SAST (10)
| ID | CWE | File | Severity |
|---|---|---|---|
| GRID-SAST-001 | CWE-78  | `internal/scada/runner.go` | High |
| GRID-SAST-002 | CWE-798 | `internal/config/config.go` | Critical |
| GRID-SAST-003 | CWE-326 | `internal/transport/tls.go` | High |
| GRID-SAST-004 | CWE-338 | `internal/auth/token.go` | Medium |
| GRID-SAST-005 | CWE-117 | `internal/middleware/log.go` | Medium |
| GRID-SAST-006 | CWE-918 | `internal/dr/dispatch.go` | High |
| GRID-SAST-007 | CWE-362 | `internal/cache/store.go` | Medium |
| GRID-SAST-008 | CWE-248 | `internal/api/meters.go` | Medium |
| GRID-SAST-009 | CWE-327 | `internal/auth/token.go` | Medium |
| GRID-SAST-010 | CWE-119 | `internal/protocol/modbus.go` | High |

## IaC (8)
| ID | Class | File |
|---|---|---|
| GRID-IAC-001 | storage-public        | `infra/terraform/storage.tf` |
| GRID-IAC-002 | nsg-open-mgmt         | `infra/terraform/network.tf` |
| GRID-IAC-003 | pg-flex-public        | `infra/terraform/postgres.tf` |
| GRID-IAC-004 | pg-no-tls             | `infra/terraform/postgres.tf` |
| GRID-IAC-005 | container-root        | `Dockerfile` |
| GRID-IAC-006 | container-latest      | `Dockerfile` |
| GRID-IAC-007 | k8s-no-resource-limit | `infra/k8s/deployment.yaml` |
| GRID-IAC-008 | k8s-host-pid          | `infra/k8s/deployment.yaml` |

## SCA (3)
| ID | Module | Version | CVE |
|---|---|---|---|
| GRID-SCA-001 | github.com/dgrijalva/jwt-go | 3.2.0   | CVE-2020-26160 |
| GRID-SCA-002 | gopkg.in/yaml.v2            | 2.2.2   | CVE-2019-11254 |
| GRID-SCA-003 | github.com/gin-gonic/gin    | 1.6.0   | CVE-2020-28483 |

## Pipeline (3)
| ID | File | Description |
|---|---|---|
| GRID-CI-001 | `.github/workflows/ci.yml` | Hardcoded Azure SP credentials |
| GRID-CI-002 | `.github/workflows/ci.yml` | `actions/checkout@v1` |
| GRID-CI-003 | `.github/workflows/ci.yml` | Missing permissions block |

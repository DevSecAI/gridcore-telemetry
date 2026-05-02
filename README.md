# Gridcore Telemetry

> ⚠️ **ARKO Benchmark Application — contains intentional vulnerabilities. Do not deploy.**

Smart-meter telemetry ingestion and demand-response orchestration for the (fictional) **Gridcore** distribution system operator. Part of the [ARKO Coverage Benchmarks](https://github.com/DevSecAI/arko-benchmarks) suite.

Stack: Go 1.22 · Chi · TimescaleDB · Azure (Terraform) · Kubernetes.

## Coverage

- **10 SAST** — command injection, hardcoded creds, weak TLS config, insecure random, log injection, SSRF, race condition, panic on user input, weak hash, unsafe pointer.
- **8 IaC** — Azure Storage public access, NSG rules, Postgres flexible server config, Dockerfile, K8s.
- **3 SCA** — Go modules with known CVEs.
- **3 pipeline misconfigs**.

Frameworks exercised: **NIS2**, **IEC 62443**, **NIST CSF**.

See [`BENCHMARK.md`](./BENCHMARK.md).

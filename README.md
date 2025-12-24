This repository is used to solve observability-related assignments.
Each assignment or problem statement is implemented in its own dedicated branch, named after the corresponding assignment.

Reviewers can check out the relevant branch to review a specific solution. Once completed, all assignment branches are merged into the main branch for reference and consolidation.

# Story 1: Basic Observability Setup
## Problem Statement
You have a simple Go HTTP service that handles user requests, but you have no visibility into how long requests take or if any are failing. Users are complaining about slow responses, but you can't identify which endpoints are problematic.
### Tasks
1. Create a simple Go[ or any other language] HTTP application / use any open-source HTTP application:
2. Add OpenTelemetry instrumentation:Use auto-Instrumentation
3. Set up OTel Collector locally:<br>
    i. Configure OTLP receiver<br>
    ii. Export to backend (clickhouse)<br>
    iii. Run collector with Docker<br>
4. Visualize in Grafana:<br>
    i. Connect Grafana to trace backend<br>
    ii. View traces for your requests<br>
    iii. Identify which endpoint is slowest<br>
### Success Criteria<br>
i. Application successfully sends traces to OTel Collector<br>
ii. Can view traces in Grafana for all endpoints<br>
iii. Can identify the slowest endpoint from trace data<br>
iv. Traces show duration and status for each request<br>

#!/bin/bash -e

# generate mocks
mockery -name ExecutionSDK -dir ./orchestrator -output ./orchestrator/mocks
mockery -name EventSDK -dir ./orchestrator -output ./orchestrator/mocks
mockery -name ProcessSDK -dir ./orchestrator -output ./orchestrator/mocks

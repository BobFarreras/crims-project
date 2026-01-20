run-all:
docker-compose up -d

test-unit:
cd backend && go test ./...
cd frontend && npm test

test-e2e:
npx playwright test

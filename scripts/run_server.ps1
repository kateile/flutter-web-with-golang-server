Write-Host "Starting..."
Set-Location $PSScriptRoot
cd..
cd backend

go run server.go

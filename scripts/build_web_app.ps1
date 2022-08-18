Write-Host "Building web app........"

Set-Location $PSScriptRoot
cd..
cd frontend
flutter build web

Set-Location $PSScriptRoot
Copy-Item -Path "..\frontend\build\web\" -Destination "..\backend\web\" -Recurse -Force

Write-Host "Done building web app..."
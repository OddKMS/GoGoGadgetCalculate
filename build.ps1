#Build-script for GoCalculate
$env:GOOS="linux"
$goPath = Split-Path (Split-Path $PSScriptRoot -Parent) -Parent;
$goZipCmd = $goPath + "\bin\build-lambda-zip.exe -o GoCalculate.zip GoCalculate";

go build -o GoCalculate;

Invoke-Expression $goZipCmd;
#reset the GOOS env-var to let me build to .exe in VSCode directly from terminal
$env:GOOS="windows"
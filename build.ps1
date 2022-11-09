$INNOSETUP="$pwd\nvm.iss"
$ORIG="$pwd"
$env:GOARCH="386"
$SRC="$pwd\src"

if (Test-Path "nvm.exe") {
  Remove-Item -Recurse -Force nvm.exe
}

echo "Building nvm.exe"
cd $SRC
go build -o ..\nvm.exe .
cd $ORIG

$DIST="$pwd\dist"
$BIN="$pwd\bin"

if (Test-Path $DIST) {
  Remove-Item -Recurse $DIST
}

New-item -Path $DIST -ItemType Directory

Move-Item nvm.exe $DIST
$cmd = "..\buildtools\zip.exe -j -9 -r nvm-noinstall.zip nvm.exe"

$listArr = ,"noinstall.cmd"
for($i=0; $i -lt $listArr.Length; $i++) {
  $name = $listArr[$i]
  Copy-Item "$BIN\$name" "$DIST\"
  $cmd += " $name"
}

cd $DIST
echo $cmd
Invoke-Expression $cmd
cd $ORIG

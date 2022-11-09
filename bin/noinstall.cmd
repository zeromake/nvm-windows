set NVM_HOME="%~dp0"
set NVM_SYMLINK="%~dp0nodejs"

if "%PROCESSOR_ARCHITECTURE%" == "X86" (
    set SYS_ARCH="32"
) else (
    set SYS_ARCH="64"
)

(echo root: %NVM_HOME% && echo path: %NVM_SYMLINK% && echo arch: %SYS_ARCH% && echo proxy: none) > %NVM_HOME%settings.txt

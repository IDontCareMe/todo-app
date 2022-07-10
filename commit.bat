@echo off
git add .
set /P comment="Enter comment: "
git commit -m %comment%
git push -u origin main
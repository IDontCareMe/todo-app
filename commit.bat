@echo off
git add .
set /P comment="Enter comment:/n"
git commit -m %comment%
git push -u origin main
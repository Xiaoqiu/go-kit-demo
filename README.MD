git remote set-url origin https://<githubtoken>@github.com/<username>/<repositoryname>.git
git push -u origin main

## github 443问题
git config --global https.proxy http://127.0.0.1:7890
git config --global http.proxy http://127.0.0.1:7890